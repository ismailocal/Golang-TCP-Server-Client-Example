package structure

import (
	"encoding/json"
	"net"
)

type Request struct {
	Msg    		string 	`json:"msg"`
	PeerId   	string 	`json:"peer_id"`
	Token      	string 	`json:"token"`
	ServerIp	string	`json:"server_ip,omitempty"`
	Ticket		string	`json:"ticket,omitempty"`
	Status		string	`json:"status,omitempty"`
}

func (request Request) Send(connection net.Conn){
	rJson, _ := json.Marshal(request)
	_, _ = connection.Write([]byte(string(rJson) + "\n"))
}