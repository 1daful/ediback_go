package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

type Config struct {
	AppId  string `json:"appId"`
	Apikey string `json:"apikey"`
}

func getConfig() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()
	byte, err := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byte, &config)
	return config
}
func addIndex(url string, name string) {
	config := getConfig()
	client := search.NewClient(config.AppId, config.Apikey)
	index := client.InitIndex(name)

	getRes, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer getRes.Body.Close()

	var posts []map[string]interface{}
	err = json.NewDecoder(getRes.Body).Decode(&posts)
	// error handling
	if err != nil {
		log.Println(err)
	}

	res, err := index.SaveObjects(posts, opt.AutoGenerateObjectIDIfNotExist(true))
	// error handling
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}

func configIndex(name string) {
	config := getConfig()
	client := search.NewClient(config.AppId, config.Apikey)
	index := client.InitIndex(name)

	res, err := index.SetSettings(search.Settings{
		// Select the attributes you want to search in
		SearchableAttributes: opt.SearchableAttributes(
			"title", "authors", "genre", "description", "tags", "isbn",
		),
		// Define business metrics for ranking and sorting
		CustomRanking: opt.CustomRanking(
			/*post_date*/ "desc(created)", "desc(record_index)",
		),
		// Set up some attributes to filter results on
		AttributesForFaceting: opt.AttributesForFaceting(
			"genre", "tags", "created",
		),
		// Define the attribute we want to distinct on
		AttributeForDistinct: opt.AttributeForDistinct(
			"id",
		),
	})
	// error handling
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
