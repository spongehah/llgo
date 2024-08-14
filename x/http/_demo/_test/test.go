package main

import "net/http"

func main() {
	http.Post("http://example.com", "application/json", nil)
}
