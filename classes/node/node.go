package node

import (
	"encoding/json"
	"io"
	"net/http"
)

type Node struct {
	IpAddress string
}

type HealthCheckResponse struct {
	Role       string `json:"role" example:"Primary"`
	DiskStatus string `json:"disk-status" example:"UpToDate"`
}

func NewNode() *Node {
	//logic for creating new node on aws
	createdNode := Node{
		IpAddress: "",
	}
	return &createdNode
}

func (n *Node) CheckStatus() (*HealthCheckResponse, error) {
	response, err := http.Get(n.IpAddress + "/healthCheck")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var healthCheckResponse HealthCheckResponse
	if err := json.Unmarshal(body, &healthCheckResponse); err != nil {
		return nil, err
	}

	return &healthCheckResponse, nil

}
