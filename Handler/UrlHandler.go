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

func (p *Playlist) GetPlaylist() {
	fmt.Println(p)
}

func (u *UserContent) GetSpotifyPlaylist(p *Playlist, currentUrl string) (string, error) {
	var req *http.Request
    
	if currentUrl == "" {
		req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/4MRFMcazMrU660XGjoyjhp/tracks?market=US&fields=items(track(name%2Calbum(name)))%2Cnext%2Ctotal&offset=0", nil)
	} else {
		req, _ = http.NewRequest("GET", currentUrl, nil)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer BQAYKJ-oXk_y4DVOaykyISMlXUVvs1w4MxRzgvYFM5tVtgApjkjD5UGa5Yr-N7giIx5fJR5HbSe7KGavG36OB91ENDQtbHn7cylBoNbcMDK4g4TtE1B35YWqTXQ8SkBEzYm_A7-SuYSrIuHm6G1LQv9y67gGXYdDdTydk5Wm0IQS")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	jsn, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(jsn, &p)
	if err != nil {
		return "", err
	}

	return "", nil
}
