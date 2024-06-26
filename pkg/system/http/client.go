package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	GetRequest(endpoint string, params map[string]string) ([]byte, error)
}

type client struct {
	config APIConfig
}

func NewClient(config APIConfig) Client {
	return client{config: config}
}

func (c client) GetRequest(endpoint string, params map[string]string) ([]byte, error) {
	apiUrl := fmt.Sprintf("%s%s?apikey=%s", c.config.Host, endpoint, c.config.APIKey)

	request, errRequest := http.NewRequest("GET", apiUrl, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	if params != nil {
		q := request.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		request.URL.RawQuery = q.Encode()
	}

	request.Header.Set("accept", "application/json; charset=utf-8")

	httpClient := &http.Client{}
	response, errDo := httpClient.Do(request)
	if errDo != nil {
		return nil, errDo
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(respBytes))
	}

	return respBytes, nil
}
