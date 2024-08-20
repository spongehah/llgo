package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	url := "http://httpbin.org/post"
	filePath := "/Users/spongehah/go/src/llgo/x/http/_demo/get/get.go"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post(url, writer.FormDataContentType(), body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(respBody))
}
