package peer

import (
	"bufio"
	"net"
)

type Peer struct {
	PeerId   	string
	Token    	string
	ServerIp 	string
	Auth 		bool
	Connection 	net.Conn
	Reader 		bufio.Reader
}