package job

import (
	. "../structure"
	"encoding/json"
	"net"
)

type Job struct {
	Name    	string `json:"name"`
	Os   		string `json:"os"`
	Distro  	string `json:"distro"`
	Command 	string `json:"command"`
}

type RequestJob struct {
	Request
	Job			Job		`json:"job,omitempty"`
}

func (request RequestJob) Send(connection net.Conn){
	rJson, _ := json.Marshal(request)
	_, _ = connection.Write([]byte(string(rJson) + "\n"))
}