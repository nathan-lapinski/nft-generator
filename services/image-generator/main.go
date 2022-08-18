package main

import (
	"fmt"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeDataToFile() {
	f, err := os.Create("/tmp/dat1")
	check(err)

	defer f.Close()

	d1 := []byte{115, 111, 109, 101, 10}
	n1, err := f.Write(d1)
	check(err)
	fmt.Printf("wrote %d bytes\n", n1)
	f.Sync()
}

func generateImage(w http.ResponseWriter, req *http.Request) {
	writeDataToFile()
	fmt.Fprintf(w, "Hello from image gen\n")
}

func main() {
	http.HandleFunc("/generate", generateImage)
	http.ListenAndServe(":8090", nil)
}