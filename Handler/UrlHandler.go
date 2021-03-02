package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

type UserContent struct {
	spotifyUrl string
	Token      string `json:"Token"`
}

type Playlist struct {
	Items []struct {
		Track struct {
			Album struct {
				Name string `json:"name"`
			} `json:"album"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
	Next  string `json:"next"`
	Total int    `json:"total"`
}

func New(userPlaylist string) UserContent {
	u := UserContent{}
	u.spotifyUrl = userPlaylist
	getARL(&u)
	return u
}

func getARL(u *UserContent) {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(u)
}

func (u *UserContent) PrintPlaylist(p []Playlist) {
	fmt.Println(p)
}

func (u *UserContent) GetSpotifyPlaylist(p []Playlist) (string, error) {
	var req *http.Request
	var playlistLength float64 = 1

	for i := 0; i < int(playlistLength); i++ {

		if i == 0 {
			req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+u.spotifyUrl+"/tracks?market=US&fields=items(track(name%2Calbum(name)))%2Cnext%2Ctotal", nil)
		} else {
			req, _ = http.NewRequest("GET", p[i-1].Next, nil)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+u.Token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		jsn, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal(jsn, &p[i])
		if err != nil {
			return "", err
		}

		// Need support for playlists < 100 songs
		if i == 0 {
			playlistLength = math.Trunc(float64(p[0].Total/100)) + 1
		}
	}
	return "", nil
}
