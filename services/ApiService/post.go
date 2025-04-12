package apiservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Post(endpoint string, payload *map[string]interface{}) {
	fmt.Println("************************")
	fmt.Println(*payload)
	fmt.Println("************************")
	jsonData, _ := json.Marshal(*payload)

	_, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err.Error())
	}

}
