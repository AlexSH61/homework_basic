package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

func SendGetRequest(serverURL, path string) (string, error) {
	url := fmt.Sprintf("http://%s/%s", serverURL, path)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return "", err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SendPostRequest(serverURL, path string, payload []byte) (string, error) {
	url := fmt.Sprintf("http://%s/%s", serverURL, path)
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
