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

type PokemonResponse struct {
	Abilities              []Abilities     `json:"abilities"`
	BaseExperience         int             `json:"base_experience"`
	Cries                  Cries           `json:"cries"`
	Forms                  []Forms         `json:"forms"`
	GameIndices            []GameIndices   `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              []HeldItems     `json:"held_items"`
	Id                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []Moves         `json:"moves"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	PastAbilities          []PastAbilities `json:"past_abilities"`
	PastStats              []PastStats     `json:"past_stats"`
	PastTypes              []any           `json:"past_types"`
	Species                Species         `json:"species"`
	Sprites                Sprites         `json:"sprites"`
	Stats                  []Stats         `json:"stats"`
	Types                  []Types         `json:"types"`
	Weight                 int             `json:"weight"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Forms struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GameIndices struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type HeldItems struct {
	Item           Item                     `json:"item"`
	VersionDetails []HeldItemVersionDetails `json:"version_details"`
}

type Item struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type HeldItemVersionDetails struct {
	Rarity  int     `json:"rarity"`
	Version Version `json:"version"`
}

type Moves struct {
	Move                Move                  `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}

type Move struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionGroupDetails struct {
	LevelLearnedAt  int             `json:"level_learned_at"`
	MoveLearnMethod MoveLearnMethod `json:"move_learn_method"`
	Order           any             `json:"order"`
	VersionGroup    VersionGroup    `json:"version_group"`
}

type MoveLearnMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionGroup struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PastAbilities struct {
	Abilities  []Abilities `json:"abilities"`
	Generation Generation  `json:"generation"`
}

type Generation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PastStats struct {
	Generation Generation `json:"generation"`
	Stats      []Stats    `json:"stats"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Species struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Sprites struct {
	BackDefault      string   `json:"back_default"`
	BackFemale       string   `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  string   `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      string   `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale string   `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}

type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}

type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type Home struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationIx   GenerationIx   `json:"generation-ix"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}

type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

type Yellow struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}

type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}

type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}

type Gold struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}

type Silver struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}

type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}

type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}

type DiamondPearl struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type HeartgoldSoulsilver struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type Platinum struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type GenerationIx struct {
	ScarletViolet ScarletViolet `json:"scarlet-violet"`
}

type ScarletViolet struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}

type BlackWhite struct {
	Animated         Animated `json:"animated"`
	BackDefault      string   `json:"back_default"`
	BackFemale       string   `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  string   `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      string   `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale string   `json:"front_shiny_female"`
}

type Animated struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}

type OmegarubyAlphasapphire struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type XY struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}

type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type UltraSunUltraMoon struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

type GenerationViii struct {
	BrilliantDiamondShiningPearl BrilliantDiamondShiningPearl `json:"brilliant-diamond-shining-pearl"`
	Icons                        Icons                        `json:"icons"`
}

type BrilliantDiamondShiningPearl struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
