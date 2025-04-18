package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	apiservice "github.com/m4xkub/capstonev2_master/services/ApiService"
)

type Node struct {
	Id        string
	PublicIp  string
	PrivateIp string
}

type HealthCheckResponse struct {
	Role       string `json:"role" example:"Primary"`
	DiskStatus string `json:"disk-status" example:"UpToDate"`
}

func NewNode(publicIp string, privateIp string) *Node {
	//logic for creating new node on aws
	createdNode := Node{
		PublicIp:  publicIp,
		PrivateIp: privateIp,
	}
	return &createdNode
}

func (n *Node) CheckStatus() (*HealthCheckResponse, error) {
	response, err := http.Get("http://" + n.PublicIp + ":8080" + "/healthCheck")
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
	// fmt.Println("What the fuck")
	fmt.Println(healthCheckResponse)
	// fmt.Println("What the fuck")

	return &healthCheckResponse, nil

}

func (n *Node) PromoteToPrimary() error {
	status, err := n.CheckStatus()
	if err != nil {
		return err
	}

	if status.Role == "Primary" {
		return errors.New("node is already roled as primary")
	}

	_, err = http.Get("http://" + n.PublicIp + ":8080" + "/promote")

	if err != nil {
		return err
	}

	_, err = apiservice.Get(fmt.Sprintf("http://%s:8080/mountVolume", n.PublicIp))

	if err != nil {
		return err
	}
	return nil
}

func (n *Node) DemoteToSecondary() error {
	status, err := n.CheckStatus()
	if err != nil {
		return err
	}

	if status.Role == "Secondary" {
		return errors.New("node is already roled as secondary")
	}

	_, err = http.Get("http://" + n.PublicIp + ":8080" + "/demote")

	if err != nil {
		return err
	}

	return nil

}
