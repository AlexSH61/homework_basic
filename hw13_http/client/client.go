package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGetRequest(serverURL, path string) (string, error) {
	url := fmt.Sprintf("http://%s/%s", serverURL, path)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SendPostRequest(serverURL, path string, payload []byte) (string, error) {
	url := fmt.Sprintf("http://%s/%s", serverURL, path)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
