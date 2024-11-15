package server

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func MakeRequest(url, method string, headers map[string]string, body []byte) ([]byte, error) {
	client := http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
