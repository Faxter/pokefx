package pokeapi

type MapListResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type SpecificMapResponse struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	Id                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod                 `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetails `json:"version_details"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EncounterMethodVersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type Version struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

type Language struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon                   `json:"pokemon"`
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EncounterVersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}

type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
