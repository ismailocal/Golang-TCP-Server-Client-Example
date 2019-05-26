package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

import "../peer"
import "../job"
import "../database"

import . "../structure"

var listener net.Listener

func Listen(port string) net.Listener{
	listener, _ = net.Listen("tcp", ":" + port)
	fmt.Println("Running server")
	return listener
}

func Connection() net.Conn{
	connection, _ := listener.Accept()
	fmt.Println("Connected peer")
	return connection
}

func Dial(port string) net.Conn{
	connection, _ := net.Dial("tcp", ":" + port)
	return connection
}

func ConnectionHandler(connection net.Conn){

	database.Connection()
	defer database.DB.Close()

	client := peer.Peer{Connection: connection}

	for{
		switch client.Auth {
		case false:
			request := Request{}

			message, _ := bufio.NewReader(connection).ReadString('\n')
			err := json.Unmarshal([]byte(message), &request)

			if err != nil {
				client.RejectedResponse().Send(connection)
				continue
			}

			client.PeerId = request.PeerId
			client.Token = request.Token

			database.AddLog(connection, client.Token)

			hasPeer := peer.Check(client)
			if !hasPeer {
				client.RejectedResponse().Send(connection)
				continue
			}
			client.Auth = true
			peer.Register(client, request, connection)
			client.AcceptedResponse().Send(connection)

		case true:
			fmt.Println("Sending jobs")
			for i, _job := range job.Get() {
				_job.Request(client).Send(connection)

				message, _ := bufio.NewReader(connection).ReadString('\n')
				fmt.Println(fmt.Sprintf("Job %d completed", i + 1))
				database.AddResult(connection, client.PeerId, message)
			}
			fmt.Println("All jobs completed")
			return
		}
	}
}

func ShutdownHandler(){
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		for _, client := range peer.Get(){
			if client.Connection != nil {
				client.ShutdownResponse().Send(client.Connection)
				client.Connection.Close()
			}
		}
		os.Exit(0)
	}()
}