package Handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
)

type DeezerLinks struct {
	Data []struct {
		Link string `json:"link"`
	} `json:"data"`
}

type UserContent struct {
	songCount  int
	spotifyUrl string
	Token struct {
		SpotifyToken string `json:"Token"`
		Arl   string `json:"ARL"`
	}
	finalLinks []DeezerLinks
}

type Playlist struct {
	Items []struct {
		Track struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"items"`
	Next  string `json:"next"`
	Total int    `json:"total"`
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
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
	jsonParser.Decode(&u.Token)
}

func (u *UserContent) PrintPlaylist(p []Playlist) {
	fmt.Println(p)
}

func (u *UserContent) GetSpotifyPlaylist(p []Playlist) error {
	var req *http.Request
	var playlistLength float64 = 1

	for i := 0; i < int(playlistLength); i++ {

		if i == 0 {
			req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/"+u.spotifyUrl+"/tracks?market=US&fields=items(track(name%2Cartists(name)%2C))%2Cnext%2Ctotal", nil)
			fmt.Println("GET", "https://api.spotify.com/v1/playlists/"+u.spotifyUrl+"/tracks?market=US&fields=items(track(name%2Cartists(name)%2C))%2Cnext%2Ctotal")
		} else {
			req, _ = http.NewRequest("GET", p[i-1].Next, nil)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+u.Token.SpotifyToken)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		jsn, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(jsn, &p[i])
		if err != nil {
			return err
		}
		if p[i].Error.Status != 0 {
			return errors.New(p[i].Error.Message)
		}

		// Need support for playlists < 100 songs
		if i == 0 {
			if playlistLength > 100 {
				playlistLength = math.Trunc(float64(p[0].Total/100)) + 1
			} else {
				playlistLength = 1
			}
		}
	}
	return nil
}
