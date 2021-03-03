package main

import (
	"Spotify-FLAC-dl/Handler"
	"fmt"
  "log"
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
}

func getInput() string {
	var input string

	fmt.Println("Input spotify playlist link: ")
	fmt.Scanln(&input)
  return input 
}
