package api

type LocationAreaResult struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates,omitempty"`
	GameIndex            int                    `json:"game_index,omitempty"`
	ID                   int                    `json:"id,omitempty"`
	Location             Location               `json:"location,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Names                []Names                `json:"names,omitempty"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters,omitempty"`
}
type EncounterMethod struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Version struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type EncounterVersionDetails struct {
	Rate    int     `json:"rate,omitempty"`
	Version Version `json:"version,omitempty"`
}
type EncounterMethodRates struct {
	EncounterMethod EncounterMethod           `json:"encounter_method,omitempty"`
	VersionDetails  []EncounterVersionDetails `json:"version_details,omitempty"`
}
type Location struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Language struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Names struct {
	Language Language `json:"language,omitempty"`
	Name     string   `json:"name,omitempty"`
}
type Pokemon struct {
	Name           string `json:"name,omitempty"`
	BaseExperience int    `json:"base_experience"`
	URL            string `json:"url,omitempty"`
}
type ConditionValues struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Method struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type EncounterDetails struct {
	Chance          int               `json:"chance,omitempty"`
	ConditionValues []ConditionValues `json:"condition_values,omitempty"`
	MaxLevel        int               `json:"max_level,omitempty"`
	Method          Method            `json:"method,omitempty"`
	MinLevel        int               `json:"min_level,omitempty"`
}
type VersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details,omitempty"`
	MaxChance        int                `json:"max_chance,omitempty"`
	Version          Version            `json:"version,omitempty"`
}
type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon,omitempty"`
	VersionDetails []VersionDetails `json:"version_details,omitempty"`
}

type RespShallowLocations struct {
	Count    int     `json:"count,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}
