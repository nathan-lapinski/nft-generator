package main

import (
	"bytes"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"time"

	"github.com/aofei/cameron"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateImage(w http.ResponseWriter, req *http.Request) {
	// instead of writing buffer to the filesystem, it might be better to keep it in memory and send directly to IPFS service
	buf := bytes.Buffer{}
    png.Encode(&buf, cameron.Identicon([]byte(req.RequestURI), 540, 60))
    w.Header().Set("Content-Type", "image/png")

	filename := fmt.Sprintf("%s%d%s", "./images/dat", time.Now().Unix(), ".png")
	f, err := os.Create(filename)
	check(err)

	defer f.Close()
	buf.WriteTo(f)
	fmt.Fprintf(w, "image generated successfully\n")
}

func main() {
	http.HandleFunc("/generate", generateImage)
	http.ListenAndServe(":8090", nil)
}