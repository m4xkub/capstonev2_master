package apiservice

import "net/http"

func Get(endpoint string) {
	_, err := http.Get(endpoint)

	if err != nil {
		return
	}
}
