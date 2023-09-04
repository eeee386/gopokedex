package service

import (
	"GoPokedex/pokecache"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var cache = pokecache.NewCache(time.Second * 300)

type LocationResponseObject struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) LocationResponseObject {
	val, ok := cache.Get(url)
	var resultBody LocationResponseObject
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &resultBody)
		return resultBody
	} else {
		json.Unmarshal(val, &resultBody)
		return resultBody
	}
}
