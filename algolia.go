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

func addIndex(url string, name string) {
	jsonFile, err := os.Open("config.json")
	defer jsonFile.Close()
	byte, err := ioutil.ReadAll(jsonFile)

	type Config struct {
		AppId  string `json:"appId"`
		Apikey string `json:"apikey"`
	}

	var config Config
	json.Unmarshal(byte, &config)

	client := search.NewClient("JFUHLV2WO0", "b0a6da35268c835cf1e2853e3588e10a")
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

func configIndex() {
	client := search.NewClient("JFUHLV2WO0", "b0a6da35268c835cf1e2853e3588e10a")
	index := client.InitIndex("demo_media")

	res, err := index.SetSettings(search.Settings{
		// Select the attributes you want to search in
		SearchableAttributes: opt.SearchableAttributes(
			"post_title", "author_name", "categories", "content",
		),
		// Define business metrics for ranking and sorting
		CustomRanking: opt.CustomRanking(
			"desc(post_date)", "desc(record_index)",
		),
		// Set up some attributes to filter results on
		AttributesForFaceting: opt.AttributesForFaceting(
			"categories",
		),
		// Define the attribute we want to distinct on
		AttributeForDistinct: opt.AttributeForDistinct(
			"post_id",
		),
	})
	// error handling
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
