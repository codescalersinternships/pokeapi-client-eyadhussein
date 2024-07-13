package pokeapiclient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	loadTimeoutTests = []struct {
		name     string
		timeout  time.Duration
		expected error
	}{
		{"valid timeout", validTimeout, nil},
		{"invalid timeout", inValidTimeout, ErrTimeout},
	}

	loadMethodTests = []struct {
		name   string
		method string
	}{
		{"valid method", http.MethodGet},
		{"invalid method", http.MethodPost},
	}
)

const (
	testTimeout    = 3 * time.Second
	validTimeout   = testTimeout - 1*time.Second
	inValidTimeout = testTimeout + 1*time.Second
)

func TestClient_GetPokemonByID(t *testing.T) {
	for _, tt := range loadTimeoutTests {
		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, tt.timeout)

			_, err := pokeClient.GetPokemonByID(35)

			assertEqual(t, err, tt.expected)
		})
	}
	for _, tt := range loadMethodTests {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != tt.method {
				t.Error("expected method", tt.method, "got", r.Method)
			}
		}))

		defer mockServer.Close()

		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(mockServer.URL, testTimeout, validTimeout)
			_, _ = pokeClient.GetPokemonByID(35)
		})
	}
	t.Run("get pokemon using id successfully", func(t *testing.T) {
		pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, validTimeout)

		expected := &Pokemon{}

		err := json.Unmarshal(expectedPokemon, expected)

		assertNoError(t, err)

		pokemon, err := pokeClient.GetPokemonByID(35)

		assertNoError(t, err)

		assertEqual(t, pokemon, expected)
	})

	t.Run("handle invalid JSON response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id":35,"name":"clefairy","base_experience":113,"height":6,"is_default":true,"order":64,"weight":75`))
			assertNoError(t, err)
		}))
		defer mockServer.Close()

		pokeClient := NewPokeClient(mockServer.URL, testTimeout)

		_, err := pokeClient.GetPokemonByID(35)

		assertEqual(t, err, ErrInvalidJSON)
	})
}

func TestClient_GetPokemonByName(t *testing.T) {
	for _, tt := range loadTimeoutTests {
		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, tt.timeout)

			_, err := pokeClient.GetPokemonByName("clefairy")

			assertEqual(t, err, tt.expected)
		})
	}
	for _, tt := range loadMethodTests {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != tt.method {
				t.Error("expected method", tt.method, "got", r.Method)
			}
		}))

		defer mockServer.Close()

		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(mockServer.URL, testTimeout, validTimeout)
			_, _ = pokeClient.GetPokemonByName("clefairy")
		})
	}
	t.Run("get pokemon using id successfully", func(t *testing.T) {
		pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, validTimeout)

		expected := &Pokemon{}

		err := json.Unmarshal(expectedPokemon, expected)

		assertNoError(t, err)

		pokemon, err := pokeClient.GetPokemonByName("clefairy")

		assertNoError(t, err)

		assertEqual(t, pokemon, expected)
	})

	t.Run("handle invalid JSON response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id":35,"name":"clefairy","base_experience":113,"height":6,"is_default":true,"order":64,"weight":75`))
			assertNoError(t, err)
		}))
		defer mockServer.Close()

		pokeClient := NewPokeClient(mockServer.URL, testTimeout)

		_, err := pokeClient.GetPokemonByName("clefairy")

		assertEqual(t, err, ErrInvalidJSON)
	})
}

func TestClient_GetPokemons(t *testing.T) {
	for _, tt := range loadTimeoutTests {
		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, tt.timeout)

			_, err := pokeClient.GetPokemons(20)

			assertEqual(t, err, tt.expected)
		})
	}
	for _, tt := range loadMethodTests {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != tt.method {
				t.Error("expected method", tt.method, "got", r.Method)
			}
		}))

		defer mockServer.Close()

		t.Run(tt.name, func(t *testing.T) {
			pokeClient := NewMockPokeClient(mockServer.URL, testTimeout, validTimeout)
			_, _ = pokeClient.GetPokemons(20)
		})
	}

	t.Run("handle invalid JSON response", func(t *testing.T) {
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id":35,"name":"clefairy","base_experience":113,"height":6,"is_default":true,"order":64,"weight":75`))
			assertNoError(t, err)
		}))
		defer mockServer.Close()

		pokeClient := NewPokeClient(mockServer.URL, testTimeout)

		_, err := pokeClient.GetPokemons(20)

		assertEqual(t, err, ErrInvalidJSON)
	})
	t.Run("get list of pokemons successfully", func(t *testing.T) {

		expected := decodePokeList(t, pokemonsJSONFileName)

		pokeClient := NewMockPokeClient(pokeAPIUrl, testTimeout, validTimeout)

		pokemons, err := pokeClient.GetPokemons(20)

		assertNoError(t, err)

		assertEqual(t, pokemons, expected)
	})
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %+v got %+v", want, got)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected nil but got %v", err)
	}
}

func decodePokeList(t *testing.T, filename string) *PokemonList {
	t.Helper()
	jsonFile, err := os.Open(filename)

	if err != nil {
		t.Error(err)
	}

	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)

	if err != nil {
		t.Error(err)
	}

	pokemonList := &PokemonList{}

	err = json.Unmarshal(data, pokemonList)

	if err != nil {
		t.Error(err)
	}

	return pokemonList
}
