package main

import (
	"Spotify-FLAC-dl/Handler"
	"fmt"
)

func main() {
	h := Handler.UserContent{}
	p := []Handler.Playlist{}
	var input string
	fmt.Println("Input spotify playlist link: ")
	fmt.Scanln(&input)
	h.SetUrl(input)

	for i := 0; i < 7; i++ {
		var currentUrl string
		h.GetSpotifyPlaylist(&p[i], currentUrl)
		currentUrl = string(p[i].Next)
		p[i].GetPlaylist()
	}
}
