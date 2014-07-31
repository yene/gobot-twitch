package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//topDota2Streams()
	favoriteDota2Streams()
}

func favoriteDota2Streams() {
	favorites := []string{"zai", "fogged", "draskyl", "sing_sing", "d"}
	concatenated := strings.Join(favorites, ",")
	requestURL := "https://api.twitch.tv/kraken/streams?game=Dota+2&limit=10&channel=" + concatenated
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

}

func topDota2Streams() {
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
