package main

import (
	"Spotify-FLAC-dl/Handler"
	"fmt"
	"log"
	"strings"
)

func main() {
	p := make([]Handler.Playlist, 100, 100)
	h := Handler.New(getInput())

	err := h.GetSpotifyPlaylist(p)
	if err != nil {
		log.Fatalln(err)
	}
	err = h.GetDeezerLinks(p)
	if err != nil {
		log.Fatalln(err)
	}
	h.DownloadAll()
}

func getInput() string {
	var input string

	fmt.Println("Input spotify playlist link: ")
	fmt.Scanln(&input)
	input = strings.Split(input, "?si=")[0]
	input = strings.Split(input, "playlist/")[1]
	fmt.Println(input)
	return input
}
