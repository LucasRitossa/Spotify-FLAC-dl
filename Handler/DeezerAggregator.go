package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
type Deezer struct {
	Data []struct {
		ID                    int    `json:"id"`
		Readable              bool   `json:"readable"`
		Title                 string `json:"title"`
		TitleShort            string `json:"title_short"`
		TitleVersion          string `json:"title_version"`
		Link                  string `json:"link"`
		Duration              int    `json:"duration"`
		Rank                  int    `json:"rank"`
		ExplicitLyrics        bool   `json:"explicit_lyrics"`
		ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
		ExplicitContentCover  int    `json:"explicit_content_cover"`
		Preview               string `json:"preview"`
		Md5Image              string `json:"md5_image"`
		Artist                struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Link          string `json:"link"`
			Picture       string `json:"picture"`
			PictureSmall  string `json:"picture_small"`
			PictureMedium string `json:"picture_medium"`
			PictureBig    string `json:"picture_big"`
			PictureXl     string `json:"picture_xl"`
			Tracklist     string `json:"tracklist"`
			Type          string `json:"type"`
		} `json:"artist"`
		Album struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Cover       string `json:"cover"`
			CoverSmall  string `json:"cover_small"`
			CoverMedium string `json:"cover_medium"`
			CoverBig    string `json:"cover_big"`
			CoverXl     string `json:"cover_xl"`
			Md5Image    string `json:"md5_image"`
			Tracklist   string `json:"tracklist"`
			Type        string `json:"type"`
		} `json:"album"`
		Type string `json:"type"`
	} `json:"data"`
	Total int `json:"total"`
}


func (u *UserContent) GetDeezerLinks(p []Playlist) error {
	var req *http.Request

	for i := range p {
		queryArtistName := strings.ReplaceAll(p[i].Items[0].Track.Artists[0].Name, " ", "+")
		queryTrack := strings.ReplaceAll(p[i].Items[0].Track.Name, " ", "+")
		query := "https://api.deezer.com/search?q=" + `artist:"` + queryArtistName + `",track:"` + queryTrack + `"`
    fmt.Println("QUERY: ",query)

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
		fmt.Println("JSON DEEZER: ", string(jsn))

		err = json.Unmarshal(jsn, &u.DeezerURL)
		if err != nil {
			return err
		}
	}
	return nil
}
