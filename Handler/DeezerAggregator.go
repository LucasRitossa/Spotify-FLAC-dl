package Handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func (u *UserContent) DownloadAll() {
	for i := 0; i < u.songCount; i++ {
		if len(u.finalLinks[i].Data) != 0 {
      fmt.Println(len(u.finalLinks[i].Data))
			fmt.Println(u.finalLinks[i].Data[0].Link)
		}
	}
	for i := 0; i < u.songCount; i++ {
		if len(u.finalLinks[i].Data) != 0 {
			i1 := "python3"
			i2 := "-m"
			i3 := "deemix"
			i4 := "--portable"
			i5 := "-p"
			i6 := "./downloads"
			q := u.finalLinks[i].Data[0].Link
			cmd := exec.Command(i1, i2, i3, i4, q)
			cmd.Output()
			fmt.Println("DOWNLOADING: ", u.finalLinks[i].Data[0].Link)
			fmt.Println(i)
			os.WriteFile("./config/.arl", []byte(""), 0666)
			cmd = exec.Command(i1, i2, i3, i4, q, i5, i6)
			cmd.Output()
		}
	}
}

func (u *UserContent) GetDeezerLinks(p []Playlist) error {
	u.songCount = p[0].Total
	var req *http.Request
	var jsonPush []DeezerLinks
	j := 0
	x := 0

	jsonPush = make([]DeezerLinks, 10000, 10000) // Any way to make this dynamic?

	for i := 0; i < u.songCount; i++ {
		if i != 0 && i%100 == 0 {
			j++
			x = 0
		}

		queryArtistName := strings.ReplaceAll(p[j].Items[x].Track.Artists[0].Name, " ", "+")
		queryTrack := strings.ReplaceAll(p[j].Items[x].Track.Name, " ", "+")
		query := "https://api.deezer.com/search?q=" + `artist:"` + queryArtistName + `",track:"` + queryTrack + `"`
		query = strings.ReplaceAll(query, "Ö", "O") // proper unicode support?
		query = strings.ReplaceAll(query, "ö", "o")
		query = strings.ReplaceAll(query, "ü", "u")
		query = strings.ReplaceAll(query, "é", "e")
		query = strings.ReplaceAll(query, "✦", "")
		query = strings.ReplaceAll(query, "✰", "")

		fmt.Println("I-INDEX: ", i+1, ",", "X-INDEX: ", x, ",", "J-INDEX: ", j, ",", "QUERY: ", query)
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
		x++
	}
	u.finalLinks = jsonPush
	return nil
}
