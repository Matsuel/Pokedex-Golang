package pokedex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	pos "pokedex/struct"
)

var Pokes pos.Infos

var Pokemons []pos.Pokemon

func GetDatas() []pos.Pokemon {
	url := "https://pokeapi.co/api/v2/pokemon?limit=300"

	data, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(data.Body)

	json.Unmarshal(responseData, &Pokes)

	// fmt.Println(Pokes.Results)

	for _, url := range Pokes.Results {

		data, _ := http.Get(string(url.Url))
		responseData, _ := ioutil.ReadAll(data.Body)
		var Pokemon pos.Pokemon

		json.Unmarshal(responseData, &Pokemon)
		Pokemons = append(Pokemons, Pokemon)
	}
	return Pokemons
}
