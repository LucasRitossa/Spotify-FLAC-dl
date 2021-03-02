package main

import (
	"Spotify-FLAC-dl/Handler"
	"fmt"
)

func main() {
	p := make([]Handler.Playlist, 100, 100)
  h := Handler.New(getInput())

	fmt.Println(h.GetSpotifyPlaylist(p))
  h.PrintPlaylist(p)
}

func getInput() string {
	var input string

	fmt.Println("Input spotify playlist link: ")
	fmt.Scanln(&input)
  return input 
}
