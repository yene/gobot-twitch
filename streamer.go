package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//favorites := []string{"zai", "fogged", "draskyl", "sing_sing", "d"}
	// parse with https://api.twitch.tv/kraken/streams?game=Dota+2&limit=10&channel=zai,sing_sing,draskyl

	// https://api.twitch.tv/kraken/streams/

	requestURL := "https://api.twitch.tv/kraken/streams?game=Dota+2&limit=10"
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err)
	}
	streams, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var dat JSONResult
	if err := json.Unmarshal(streams, &dat); err != nil {
		panic(err)
	}

	for _, g := range dat.Streams {
		fmt.Println("Stream: " + g.Channel.Name + " - " + g.Channel.Status + " - " + g.Channel.URL)
	}

	//https://api.twitch.tv/channels/sing_sing stream == null?

}

func clientID() string {
	file, e := ioutil.ReadFile("./client.id")
	if e != nil {
		panic(e)
	}
	return string(file)
}

// JSON structs
type JSONResult struct {
	Streams []JSONStreams `json:"streams"`
}

type JSONStreams struct {
	Channel JSONChannel `json:"channel"`
}

type JSONChannel struct {
	Name   string `json:"display_name"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Views  int    `json:"views"`
}
