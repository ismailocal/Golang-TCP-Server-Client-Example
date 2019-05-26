package job

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

import "../peer"

var jobs [] Job

func Get() []Job {
	return jobs
}

func New(job Job){
	jobs = append(jobs, job)
}

func (job Job) Run() string{
	if len(job.Command) > 1 {
		arg := strings.Split(job.Command, " ")
		cmd := exec.Command(arg[0], arg[1])
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		return out.String()
	}else{
		return ""
	}
}

func (job Job) Request(client peer.Peer) RequestJob{
	request := RequestJob{}
	request.Msg = "register"
	request.Ticket = "126743"
	request.PeerId = client.PeerId
	request.Token = client.Token
	request.Job = job

	return request
}