package main

import (
	"fmt"
	"net/url"
)

func main() {
	URL, _ := url.Parse("http://localhost/get")
	fmt.Println(URL.String())
}
