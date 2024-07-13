package pokeapiclient

import (
	"net/http"
	"time"
)

const pokeAPIUrl = "https://pokeapi.co/api/v2/pokemon"

type Client interface {
	GetPokemonByName(string) (Pokemon, error)
	GetPokemonByID(int) (Pokemon, error)
	GetPokemons(int) (PokemonList, error)
}

type PokeClient struct {
	apiUrl string
	client *http.Client
}

func NewPokeClient(url string, timeout time.Duration) *PokeClient {
	return &PokeClient{
		client: &http.Client{
			Timeout: timeout,
		},
		apiUrl: url,
	}
}
