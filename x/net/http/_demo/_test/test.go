package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second,
	}

	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	time.Sleep(2 * time.Second)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
