package pokeapiclient

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strconv"

	"github.com/codescalersinternships/pokeapi-client-eyadhussein/pkg/backoff"
)

// Pokemon represents a Pokemon object
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	IsDefault      bool   `json:"is_default"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`

	URL string `json:"url"`
}

// PokemonList represents a list of Pokemon objects
type PokemonList struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

var (
	ErrTimeout        = errors.New("request timed out")     // request times out
	ErrNon200Response = errors.New("non-200 HTTP response") // response status code is not 200
	ErrInvalidJSON    = errors.New("invalid JSON response") // JSON response is invalid
)

// GetPokemonByID returns a Pokemon object by ID
func (p *PokeClient) GetPokemonByID(id int) (*Pokemon, error) {
	backoff := backoff.NewRealBackOff(1, 3)
	url := p.apiUrl + "/" + strconv.Itoa(id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := backoff.Retry(func() (*http.Response, error) {
		resp, err := p.client.Do(req)
		return resp, err
	})

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, ErrTimeout
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrNon200Response
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Pokemon
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	return &result, nil
}

// GetPokemonByName returns a Pokemon object by name
func (p *PokeClient) GetPokemonByName(name string) (*Pokemon, error) {
	backoff := backoff.NewRealBackOff(1, 3)
	url := p.apiUrl + "/" + name

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := backoff.Retry(func() (*http.Response, error) {
		resp, err := p.client.Do(req)
		return resp, err
	})

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, ErrTimeout
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrNon200Response
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Pokemon
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	return &result, nil
}

// GetPokemons returns a list of Pokemon objects with the specified limit
func (p *PokeClient) GetPokemons(limit int) (*PokemonList, error) {
	backoff := backoff.NewRealBackOff(1, 3)
	url := p.apiUrl + "?limit=" + strconv.Itoa(limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := backoff.Retry(func() (*http.Response, error) {
		resp, err := p.client.Do(req)
		return resp, err
	})
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return nil, ErrTimeout
		}
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result PokemonList
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	return &result, nil
}
