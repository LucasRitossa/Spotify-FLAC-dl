package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (u *UserContent) GetDeezerLinks(p []Playlist) error {
	var req *http.Request
	for i := range p {
		//		req, _ = http.NewRequest("GET", "https://api.deezer.com/search?q=artist:"+p[i].Items[0].Track.Artists[0].Name+"track:"+p[i].Items[0].Track.Name, nil)
    fmt.Println(i)
		req, _ = http.NewRequest("GET", "https://api.deezer.com/search?q=artist:%22Billy%20Joel%22%20track:%22Piano%20Man%22", nil)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		jsn, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println("JSON DEEZER: ", string(jsn))

		err = json.Unmarshal(jsn, &u.DeezerURL)
		if err != nil {
			return err
		}
	}
	return nil
}
