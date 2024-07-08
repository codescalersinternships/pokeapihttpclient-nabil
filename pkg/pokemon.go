package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetPokemonByname is a method of Client struct, used to fetch pokemon by its name field
func (c *Client) GetPokemonByname(ctx context.Context, pokemonName string) (Pokemon, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://pokeapi.co/api/v2/pokemon/"+pokemonName, nil)
	if err != nil {
		return Pokemon{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code returned! %v", response.StatusCode)
	}
	var poke Pokemon
	err = json.NewDecoder(response.Body).Decode(&poke)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error decoding into json, %v", err)
	}
	return poke, nil
}

// GetPokemonList is a method of Client struct, used to fetch number of pokemons
func (c *Client) GetPokemonList(ctx context.Context, limit, offset int) ([]string, error) {
	var pokeNamesList []string
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://pokeapi.co/api/v2/pokemon?limit="+strconv.Itoa(limit)+"&offset="+strconv.Itoa(offset), nil)
	if err != nil {
		return pokeNamesList, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return pokeNamesList, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return pokeNamesList, fmt.Errorf("unexpected status code returned! %v", response.StatusCode)
	}
	var pokeList pokeList
	err = json.NewDecoder(response.Body).Decode(&pokeList)
	if err != nil {
		return pokeNamesList, err
	}
	for _, poke := range pokeList.Results {
		pokeNamesList = append(pokeNamesList, poke.Name)
	}
	return pokeNamesList, nil
}
