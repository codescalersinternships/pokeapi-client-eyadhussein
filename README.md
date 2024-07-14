# PokeAPI Client

This package provides a simple HTTP client for interacting with the PokeAPI API. It's designed to retrieve a pokemon using its name or id as well as retrieve a list of pokemons up to a certain limit.

## Features

- Easy-to-use client for fetching pokemon
- Built-in retry mechanism with backoff strategy
- Timeout handling

## Installation

To use this package in your Go project, you can install it using:

```
go get github.com/codescalersinternships/pokeapi-client-eyadhussein
```

## Usage

Here's a basic example of how to use the PokeAPI Client:

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/codescalersinternships/datetime-client-eyadhussein/pkg/pokeapiclient"
)

func main() {
    // Create a new client
    client := pokeapiclient.NewPokeClient("https://pokeapi.co/api/v2/pokemon", 10*time.Second)

    // Get a pokemon by Id
    pokemonWithID, err := client.GetPokemonByID(35)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pokemon with id 35: %+v\n", pokemonWithID)

    // Get a pokemon by name
    pokemonWithName, err := client.GetPokemonByName("clefairy")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pokemon with name clefairy: %+v\n", pokemonWithName)

    // Get a list of pokemons
    pokemons, err := client.GetPokemons(20)
    if err != nil {
        log.Fatal(err)
    }
}
```

# How to Test

Run

```bash
go test -v ./...
```
