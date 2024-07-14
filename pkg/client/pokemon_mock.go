package pokeapiclient

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var expectedPokemon = []byte(`{"id":35,"name":"clefairy","base_experience":113,"height":6,"is_default":true,"order":64,"weight":75}`)

const pokemonsJSONFileName = "test-data/pokemons.json"

type MockHTTPClient struct {
	Client struct {
		Timeout time.Duration
	}
	SleepDuration time.Duration
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.SleepDuration > m.Client.Timeout {
		return &http.Response{StatusCode: http.StatusRequestTimeout}, ErrTimeout
	}
	return &http.Response{StatusCode: http.StatusOK}, nil
}

type MockPokeClient struct {
	apiUrl string
	client *MockHTTPClient
}

func NewMockPokeClient(url string, timeout time.Duration, sleepDuration time.Duration) *MockPokeClient {
	mockClient := &MockHTTPClient{
		Client: struct {
			Timeout time.Duration
		}{Timeout: timeout},
		SleepDuration: sleepDuration,
	}
	return &MockPokeClient{
		client: mockClient,
		apiUrl: url,
	}
}

func (m *MockPokeClient) GetPokemonByID(id int) (*Pokemon, error) {
	url := m.apiUrl + "/" + strconv.Itoa(id)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	_, err = m.client.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := expectedPokemon, nil

	result := Pokemon{}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *MockPokeClient) GetPokemonByName(name string) (*Pokemon, error) {
	url := m.apiUrl + "/" + name

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	_, err = m.client.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := expectedPokemon, nil

	result := Pokemon{}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *MockPokeClient) GetPokemons(limit int) (*PokemonList, error) {
	url := m.apiUrl + "?limit=" + strconv.Itoa(limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	_, err = m.client.Do(req)

	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(pokemonsJSONFileName)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	result := PokemonList{}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
