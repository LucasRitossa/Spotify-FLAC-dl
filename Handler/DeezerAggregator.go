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
  j := 0
  x := 0

  jsonPush = make([]DeezerLinks, 10000, 10000) // Any way to make this dynamic?

	for i := range p {
    if i == 100 {
      j++
      x = 0
    }
		queryArtistName := strings.ReplaceAll(p[j].Items[x].Track.Artists[0].Name, " ", "+")
		queryTrack := strings.ReplaceAll(p[j].Items[x].Track.Name, " ", "+")
		query := "https://api.deezer.com/search?q=" + `artist:"` + queryArtistName + `",track:"` + queryTrack + `"`
		fmt.Println("QUERY: ", query)

		fmt.Println("INDEX: ", i)
		req, _ = http.NewRequest("GET", query, nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
      fmt.Println(jsonPush)
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
     x++
	}
  u.finalLinks = jsonPush
  fmt.Print(u.finalLinks)
	return nil
}
