package Handler

import(
  "net/http"
  "io"
  "os"
  "strconv"
)

func HandleDownload(w http.ResponseWriter, r *http.Request, f string){
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


