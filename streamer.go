package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//favorites := []string{"zai", "fogged", "draskyl", "sing_sing", "d"}

	//https://api.twitch.tv/channels/sing_sing
	fmt.Println("Hello, 世界" + clientID())
}

func clientID() string {
	file, e := ioutil.ReadFile("./client.id")
	if e != nil {
		panic(e)
	}
	return string(file)
}
