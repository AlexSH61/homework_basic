package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func SendGetRequest(serverUrl, path string) (string, error) {
	response, err := http.Get(serverUrl + "/" + path)
	if err != nil {
		return "", fmt.Errorf("get failed: %v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading :%v", err)
	}
	return string(body), nil
}
func SendPostRequest(serverUrl, path string, payload []byte) (string, error) {
	response, err := http.Post(serverUrl+"/"+path, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("post failed:%v", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading: %v", err)
	}
	return string(body), nil
}
