package node

import "net/http"

type Node struct {
	IpAddress string
}

func NewNode() *Node {
	//logic for creating new node on aws
	createdNode := Node{
		IpAddress: "",
	}
	return &createdNode
}

func (n *Node) CheckStatus() bool {
	_, err := http.Get(n.IpAddress + "/healthCheck")

	if err != nil {
		return false
	} else {
		return true
	}
}
