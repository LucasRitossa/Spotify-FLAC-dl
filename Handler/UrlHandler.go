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
	playlistBackup := make([]Playlist, 1, 1)
	playlistLength := 1

	for i := 0; i < playlistLength; i++ {

		if len(p) == 1 {
			fmt.Println("1 get")
			req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/4MRFMcazMrU660XGjoyjhp/tracks?market=US&fields=items(track(name%2Calbum(name)))%2Cnext%2Ctotal&offset=0", nil)
		} else {
			fmt.Println("2 get")
			req, _ = http.NewRequest("GET", nextUrl, nil)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer BQADfXGLNG1N4_cEBRMCKQ9BvXvvqnIYQPRXLGyTGmnOk4CraCeLpo7A-0CfmHkxBMAsN6vUtUNOMacC8BhlQUfev6tQfBovsdDN89OeWnCrSBCaxE5ChFSXvepejAno3UPVDitbJiWD-vwA9i1mLAzV6IPMXrfUU-rYdl9csndI")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}

		jsn, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

    // unmarshal not mapping correctly, need all this logic to cut p.total into a single int, then allocate 
    // a slice of the correct nice (use less memory), could instead just allocate an array like 10 which is
    // longer then the max spotify playlist (lame)


    // once array allocated, need to copy old array (playlistBackup) into new array (p)
    // we need an else unmarshal because we dont want to reallocate everything if its allready there
		fmt.Println(i)
		if playlistLength == 1 {
			err = json.Unmarshal(jsn, &playlistBackup[i])
			if err != nil {
				return "", err
			}
      fmt.Println("PLAYLIST[0]: " + fmt.Sprint(p[0]))
			playlistLength = p[0].Total / 100
      fmt.Println("PLAYLIST LENGTH: " + fmt.Sprint(playlistLength))
			p = make([]Playlist, playlistLength, playlistLength)
			p[0] = playlistBackup[0]
		} else {
			err = json.Unmarshal(jsn, &p[i])
			if err != nil {
				return "", err
			}
		}

	}
	//this should return an array of Playlist pages
	return "", nil
}
