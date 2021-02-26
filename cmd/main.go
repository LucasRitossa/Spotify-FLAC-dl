package main

import (
	"Spotify-FLAC-dl/Handler"
	"fmt"
)

func main() {
	h := Handler.UserContent{}
	p := make([]Handler.Playlist, 1, 100)

	var input string
	fmt.Println("Input spotify playlist link: ")
	fmt.Scanln(&input)
	h.SetUrl(input)
	h.GetSpotifyPlaylist(p)
}
