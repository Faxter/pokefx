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

func ConvertMapListResponseToJson(response []byte) MapListResponse {
	resultJson := MapListResponse{}
	err := json.Unmarshal(response, &resultJson)
	if err != nil {
		fmt.Println("JSON conversion error: %w", err)
		return MapListResponse{}
	}
	return resultJson
}

func ConvertSpecificMapResponseToJson(response []byte) SpecificMapResponse {
	resultJson := SpecificMapResponse{}
	err := json.Unmarshal(response, &resultJson)
	if err != nil {
		fmt.Println("JSON conversion error: %w", err)
		return SpecificMapResponse{}
	}
	return resultJson
}

func (m *MapListResponse) ExtractMapNames() []string {
	var names []string
	for _, result := range m.Results {
		names = append(names, result.Name)
	}
	return names
}

func (s *SpecificMapResponse) ExtractPokemonEncounters() []string {
	var pokemonEncounters []string
	for _, encounter := range s.PokemonEncounters {
		pokemonEncounters = append(pokemonEncounters, encounter.Pokemon.Name)
	}
	return pokemonEncounters
}
