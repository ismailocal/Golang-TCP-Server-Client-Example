package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

import "./server"
import "./peer"
import "./job"

import . "./structure"

func main() {
	connection := server.Dial("8081")

	client := peer.Peer{}

	for {
		responseRegister := Response{}
		requestJob := job.RequestJob{}

		switch client.Auth {
		case false:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Request > ")
			message, _ := reader.ReadString('\n')

			_, err := connection.Write([]byte(message))
			if err != nil {
				fmt.Println("Connection lose!")
				break
			}

			message, _ = bufio.NewReader(connection).ReadString('\n')
			_ = json.Unmarshal([]byte(message), &responseRegister)

			if responseRegister.Status == "accepted" {
				fmt.Print(message + " \n")
				client.PeerId = responseRegister.PeerId
				client.Token = responseRegister.Token
				client.Auth = true
			}else{
				fmt.Print(message + " \n")
			}
		case true:
			message, _ := bufio.NewReader(connection).ReadString('\n')
			_ = json.Unmarshal([]byte(message), &requestJob)

			switch requestJob.Msg {
			case "register":
				_job := job.Job{Name: requestJob.Job.Name, Os: requestJob.Job.Os, Distro: requestJob.Job.Distro, Command: requestJob.Job.Command}
				jobResponse := _job.Run()
				jobResponseJson, _ := json.Marshal(jobResponse)
				_, _ = connection.Write([]byte(string(jobResponseJson) + "\n"))
			case "shutdown":
				fmt.Println(message)
			}
		}
	}
}
