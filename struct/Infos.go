package pokedex

type Infos struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Poke `json:"results"`
}

type Poke struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokedex struct {
	PokeTab []Pokemon
}
