package main

import (
	"fmt"
	"io"
	"sync"

	"github.com/goplus/llgo/x/net/http"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, ":", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	resp.Body.Close()
}

func main() {
	var wait sync.WaitGroup
	//resp, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp.Status)
	for i := 0; i < 50; i++ {
		wait.Add(1)
		go worker(i, &wait)
	}
	wait.Wait()
	fmt.Println("All done")
}
