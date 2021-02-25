package Handler

import(
  "net/http"
  "fmt"
  "os/exec"
)

func ExecScraper( q string ){
  out, err := exec.Command("node", "scraper.js", q).Output()
  if err != nil{
    fmt.Println(err)
  }
  output := string(out[:])
  // data is actually only sent like 1/8 of the time?
  fmt.Println(output)
}

func GetAudio(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Query().Get("playlist")
	if f == "" {
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + f)

  ExecScraper(f)
	//HandleDownload(w, r, f)
}
