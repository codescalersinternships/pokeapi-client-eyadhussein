package pokeapiclient

import (
	"net/http"
)

type Client interface {
	GetPokemonByName(string) (Pokemon, error)
	GetPokemonById(int) (Pokemon, error)
}

type PokeClient struct {
	apiUrl string
	client *http.Client
}

func NewPokeClient() *PokeClient {
	return &PokeClient{
		client: &http.Client{},
		apiUrl: "https://pokeapi.co/api/v2/pokemon",
	}
}
