package peer

import "net"

import . "../structure"

var peers []Peer

func Get() [] Peer{
	return peers
}

func New(peer Peer) Peer{
	peers = append(peers, peer)
	return peer
}

func Check(client Peer) bool{
	for _, peer := range peers{
		if client.PeerId == peer.PeerId && client.Token == peer.Token{
			return true
		}
	}

	return false
}

func Register(client Peer, request Request, connection net.Conn){
	for i, peer := range peers{
		if client.PeerId == peer.PeerId && client.Token == peer.Token{
			peers[i].Connection = connection
			peers[i].Auth = true
			peers[i].ServerIp = request.ServerIp
		}
	}
}

func (peer Peer) RejectedResponse() Response{
	response := Response{}
	response.Msg = "register"
	response.PeerId = peer.PeerId
	response.Token = peer.Token
	response.Status = "rejected"
	return response
}

func (peer Peer) AcceptedResponse() Response{
	response := Response{}
	response.Msg = "register"
	response.PeerId = peer.PeerId
	response.Token = peer.Token
	response.Status = "accepted"

	return response
}

func (peer Peer) ShutdownResponse() Response{
	response := Response{}
	response.Msg = "shutdown"
	response.Token = peer.Token
	response.Reason = "Upgrade server"

	return response
}