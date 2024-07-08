package pokeapiclient

import (
	"reflect"
	"testing"
)

func TestPokeClient_GetPokemon(t *testing.T) {
	t.Run("get pokemon using id successfully", func(t *testing.T) {
		expected := &Pokemon{
			Id:             35,
			Name:           "clefairy",
			BaseExperience: 113,
			Height:         6,
			IsDefault:      true,
			Order:          64,
			Weight:         75,
		}

		pokeClient := NewPokeClient()

		pokemon, err := pokeClient.GetPokemonById(35)

		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(pokemon, expected) {
			t.Errorf("expected %+v got %+v", expected, pokemon)
		}

	})
}
