package http

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	GetRequest(endpoint string, params url.Values) (string, error)
}

type client struct {
	config APIConfig
}

func NewClient(config APIConfig) Client {
	return client{config: config}
}

func (c client) GetRequest(endpoint string, params url.Values) (string, error) {
	apiUrl := fmt.Sprintf("%s%s?apikey=%s", c.config.Host, endpoint, c.config.APIKey)
	request, errRequest := http.NewRequest("GET", apiUrl, nil)
	if errRequest != nil {
		return "", errRequest
	}

	request.Header.Set("accept", "application/json; charset=utf-8")

	httpClient := &http.Client{}
	response, errDo := httpClient.Do(request)
	if errDo != nil {
		return "", errDo
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", errors.New(string(respBytes))
	}

	return string(respBytes), nil
}
