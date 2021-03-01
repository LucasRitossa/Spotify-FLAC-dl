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
	var playlistLength float64 = 1

	for i := 0; i < int(playlistLength); i++ {

		if i == 0 {
			req, _ = http.NewRequest("GET", "https://api.spotify.com/v1/playlists/4MRFMcazMrU660XGjoyjhp/tracks?market=US&fields=items(track(name%2Calbum(name)))%2Cnext%2Ctotal&offset=0", nil)
		} else {
			req, _ = http.NewRequest("GET", p[i-1].Next, nil)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer BQCHs9Atbz2a-r_SZYkdeWfee_54xusFhrSw1VHD5oI1cxvlE6JZDJTSEzOp5e_WD2MhtITzSdNyladK2ot59FfbvAz-Q8p0FZTsMWW1PXSKbnCFY9iZEDYWT21uXkD99OXf1S-c0M_of9c3YcuibvZp3WhDCptddzAufwI8qrHw")
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

		if i == 0 {
			playlistLength = math.Trunc(float64(p[0].Total/100)) + 1
		}
	}
	// unmarshal not mapping correctly, need all this logic to cut p.total into a single int, then allocate
	// a slice of the correct nice (use less memory), could instead just allocate an array like 10 which is
	// longer then the max spotify playlist (lame)

	// once array allocated, need to copy old array (playlistBackup) into new array (p)
	// we need an else unmarshal because we dont want to reallocate everything if its allready there
	return "", nil
}
