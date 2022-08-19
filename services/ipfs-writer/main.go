package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"io"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "https://api.pinata.cloud/pinning/pinFileToIPFS"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open("./images/test.png")
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("./images/test.png"))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}

	_ = writer.WriteField("pinataOptions", "{\"cidVersion\": 1}")
	_ = writer.WriteField("pinataMetadata", "{\"name\": \"MyFile\", \"keyvalues\": {\"company\": \"Pinata\"}}")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	authBearer := fmt.Sprintf("Bearer %s", os.Getenv("PINATA_TOKEN"))

	req.Header.Add("Authorization", authBearer)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
	  fmt.Println(err)
	  return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}