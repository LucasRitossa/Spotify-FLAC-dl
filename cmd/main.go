package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"net/http"
	"os"
	"strconv"
  "github.com/gocolly/colly"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", getAudio)
	r.Get("/colly", collyTest)
	http.ListenAndServe(":8080", r)
}

func collyTest(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

  c.Visit("https://liveleak.com")
}

func getAudio(w http.ResponseWriter, r *http.Request) { // r, what client sends to server, w what server sends to client
	f := r.URL.Query().Get("file")
	if f == "" {
		http.Error(w, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + f)

	handleDownload(w, r, f)
}

func handleDownload(w http.ResponseWriter, r *http.Request, f string) {
	Openfile, err := os.Open("./media/" + f)
	defer Openfile.Close()
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+f)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)
	return

}
