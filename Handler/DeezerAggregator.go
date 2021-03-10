package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (u *UserContent) GetDeezerLinks(p []Playlist) error {
	var req *http.Request
  var jsonPush []DeezerLinks

  jsonPush = make([]DeezerLinks, 10000, 10000) // Any way to make this dynamic?

	for i := range p {
		queryArtistName := strings.ReplaceAll(p[i].Items[0].Track.Artists[0].Name, " ", "+")
		queryTrack := strings.ReplaceAll(p[i].Items[0].Track.Name, " ", "+")
		query := "https://api.deezer.com/search?q=" + `artist:"` + queryArtistName + `",track:"` + queryTrack + `"`
		fmt.Println("QUERY: ", query)

		fmt.Println("INDEX: ", i)
		req, _ = http.NewRequest("GET", query, nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		jsn, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(jsn, &jsonPush[i])

		if err != nil {
			return err
		}

		 fmt.Println("JSON: ", jsonPush[i])
	}
  u.finalLinks = jsonPush
	return nil
}
