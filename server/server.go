package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	pos "pokedex/struct"
)

var Pokes pos.Infos

func main() {

	url := "https://pokeapi.co/api/v2/pokemon?limit=300"

	data, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(data.Body)

	json.Unmarshal(responseData, &Pokes)

	fmt.Println(Pokes.Results)

	// for name, url := range Pokes.Results {
	// 	fmt.Println(name)
	// 	fmt.Println(url)
	// }
}
