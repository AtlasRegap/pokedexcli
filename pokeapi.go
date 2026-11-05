package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AtlasRegap/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type areaDetail struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getJSON(cache *pokecache.Cache, url string, target any) error {
	data, found := cache.Get(url)
	if found {
		return json.Unmarshal(data, target)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	cache.Add(url, data)
	return nil
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func fetchLocationAreas(cache *pokecache.Cache, pageURL *string) (locationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var resp locationAreasResponse
	err := getJSON(cache, url, &resp)
	if err != nil {
		return locationAreasResponse{}, err
	}
	return resp, nil
}

func fetchAreaDetail(cache *pokecache.Cache, areaName string) (areaDetail, error) {
	url := baseURL + "/location-area/" + areaName

	var area areaDetail
	err := getJSON(cache, url, &area)
	if err != nil {
		return areaDetail{}, err
	}
	return area, nil
}

func fetchPokemon(cache *pokecache.Cache, name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var p Pokemon
	err := getJSON(cache, url, &p)
	if err != nil {
		return Pokemon{}, err
	}
	return p, nil
}
