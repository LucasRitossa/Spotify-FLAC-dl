package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type UserContent struct {
	spotifyUrl string
	ARLKEY     string `json:"ARL"`
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

func getARL() (UserContent, error) {
	var config UserContent
	configFile, err := os.Open("config.json")
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}

func (u *UserContent) SetUrl(url string) {
	u.spotifyUrl = url
}

func (u *UserContent) GetPlaylist(p []Playlist) {
	fmt.Println(p)
}

func (u *UserContent) GetSpotifyPlaylist(p []Playlist) (string, error) {
	var req *http.Request
	var nextUrl string
	playlistLength := 1

  fmt.Println(len(p))
	for i := 0; i < playlistLength; i++ {

		if p[0].Total == 0 {
			fmt.Println("1 get")
			req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/4MRFMcazMrU660XGjoyjhp/tracks?market=US&fields=items(track(name%2Calbum(name)))&offset=0", nil)
		} else {
			fmt.Println("2 get")
			req, _ = http.NewRequest("GET", nextUrl, nil)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer BQAxTgFjcWtzAL1BsRWL_Xhrgel_7XH1ScOYLXYDQPXUAA9z7DPRRpXBjRCFil7I7q0j7_-LfU5YCJwn6XwJ9-obqpPkRDjhYjewxuCfq7PKC5aggnZs-k71E-3vmutAfj90uTCY87sR5dJ5p_L8Y63JGqjAGJ3T12FU_BMAWCJB")
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
  }
		// unmarshal not mapping correctly, need all this logic to cut p.total into a single int, then allocate
		// a slice of the correct nice (use less memory), could instead just allocate an array like 10 which is
		// longer then the max spotify playlist (lame)

		// once array allocated, need to copy old array (playlistBackup) into new array (p)
		// we need an else unmarshal because we dont want to reallocate everything if its allready there
	return "", nil
}
