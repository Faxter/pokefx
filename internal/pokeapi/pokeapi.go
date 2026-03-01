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

func ConvertResponseToJson[T any](response []byte) (T, error) {
	var resultJson T
	err := json.Unmarshal(response, &resultJson)
	if err != nil {
		fmt.Println("JSON conversion error: %w", err)
		return resultJson, fmt.Errorf("JSON conversion error: %w", err)
	}
	return resultJson, nil
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
