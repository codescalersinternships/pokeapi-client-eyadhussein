// Package pokeapiclient provides an http client for the PokeAPI.
package pokeapiclient

import (
	"net/http"
	"time"
)

const pokeAPIUrl = "https://pokeapi.co/api/v2/pokemon"

// Client is an interface for the PokeAPI client
type Client interface {
	GetPokemonByName(string) (Pokemon, error)
	GetPokemonByID(int) (Pokemon, error)
	GetPokemons(int) (PokemonList, error)
}

// PokeClient is a client for the PokeAPI
type PokeClient struct {
	apiUrl string
	client *http.Client
}

// NewPokeClient creates a new PokeClient
func NewPokeClient(url string, timeout time.Duration) *PokeClient {
	return &PokeClient{
		client: &http.Client{
			Timeout: timeout,
		},
		apiUrl: url,
	}
}
