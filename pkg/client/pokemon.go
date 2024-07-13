package pokeapiclient

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"strconv"
)

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

type PokemonList struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemon `json:"results"`
}

var (
	ErrTimeout        = errors.New("request timed out")
	ErrNon200Response = errors.New("non-200 HTTP response")
	ErrInvalidJSON    = errors.New("invalid JSON response")
)

func (p *PokeClient) GetPokemonByID(id int) (*Pokemon, error) {
	url := p.apiUrl + "/" + strconv.Itoa(id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := p.client.Do(req)
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

func (p *PokeClient) GetPokemonByName(name string) (*Pokemon, error) {
	url := p.apiUrl + "/" + name

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := p.client.Do(req)
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

func (p *PokeClient) GetPokemons(limit int) (*PokemonList, error) {
	url := p.apiUrl + "?limit=" + strconv.Itoa(limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := p.client.Do(req)
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
