package main

import (
	_ "github.com/lib/pq"
)

import "./peer"
import "./job"
import "./server"

func bootstrap(){
	job.New(job.Job{"List Files", "linux", "debian", "ls -al"})
	job.New(job.Job{"Free Memory Space", "linux", "centos", "free -h"})

	peer.New(peer.Peer{PeerId: "1", Token: "1234"})
	peer.New(peer.Peer{PeerId: "2", Token: "5678"})
}

func main() {

	bootstrap()

	listener := server.Listen("8081")
	defer listener.Close()

	server.ShutdownHandler()

	for{
		go server.ConnectionHandler(server.Connection())
	}
}