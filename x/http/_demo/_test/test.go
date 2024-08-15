package main

import "net/http"

func main() {
	http.Post("https://jsonplaceholder.typicode.com/posts", "application/json; charset=UTF-8", nil)
}
