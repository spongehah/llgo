package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.Header.Set()
	resp, _ := client.Do(req)
	fmt.Println(resp.Status, "read bytes: ", resp.ContentLength)
	defer resp.Body.Close()
}
