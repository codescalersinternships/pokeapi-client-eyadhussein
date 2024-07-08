package pokeapiclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	IsDefault      bool   `json:"is_default"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func (p *PokeClient) GetPokemonById(id int) (*Pokemon, error) {
	req, err := http.NewRequest(http.MethodGet, p.apiUrl+"/"+strconv.Itoa(id), nil)

	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	result := Pokemon{}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
