package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiCall struct {
	url          string
	nextPage     string
	previousPage string
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

func CreateApiCall(baseurl string) ApiCall {
	return ApiCall{baseurl, "", ""}
}

func (a *ApiCall) SendRequest() ([]byte, error) {
	res, err := http.Get(a.url)
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
	return body, nil
}

func ConvertResponseToJson(response []byte) ApiResponse {
	resultJson := ApiResponse{}
	err := json.Unmarshal(response, &resultJson)
	if err != nil {
		fmt.Println("JSON conversion error: %w", err)
		return ApiResponse{}
	}
	return resultJson
}

func (a *ApiResponse) ExtractNames() []string {
	var names []string
	for _, result := range a.Results {
		names = append(names, result.Name)
	}
	return names
}
