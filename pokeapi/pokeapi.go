package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiCall struct {
	baseUrl     string
	limitQuery  string
	offsetQuery string
	resultJson  ApiResponse
}

type ApiResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func CreateApiCall(baseurl, limit, offset string) ApiCall {
	return ApiCall{baseurl, limit, offset, ApiResponse{}}
}

func (a *ApiCall) RequestNames() ([]string, error) {
	res, err := http.Get(a.createFullUrl())
	if err != nil {
		return nil, fmt.Errorf("Network error: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("IO error: %w", err)
	}
	err = a.convertResponseToJson(body)
	if err != nil {
		return nil, err
	}
	return a.extractNames(), nil
}

func (a *ApiCall) createFullUrl() string {
	return a.baseUrl + "?" + a.limitQuery + "&" + a.offsetQuery
}

func (a *ApiCall) convertResponseToJson(response []byte) error {
	err := json.Unmarshal(response, &a.resultJson)
	if err != nil {
		return fmt.Errorf("JSON conversion error: %w", err)
	}
	return nil
}

func (a *ApiCall) extractNames() []string {
	var names []string
	for _, result := range a.resultJson.Results {
		names = append(names, result.Name)
	}
	return names
}
