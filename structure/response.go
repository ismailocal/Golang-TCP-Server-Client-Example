package structure

import (
	"encoding/json"
	"net"
)

type Response struct {
	Msg    		string 	`json:"msg"`
	PeerId   	string 	`json:"peer_id"`
	Token      	string 	`json:"token"`
	Status      string 	`json:"status,omitempty"`
	Reason      string 	`json:"reason,omitempty"`
}

func (response Response) Send(connection net.Conn){
	rJson, _ := json.Marshal(response)
	_, _ = connection.Write([]byte(string(rJson) + "\n"))
}