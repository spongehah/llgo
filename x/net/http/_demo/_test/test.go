package main

import "sync"

var once sync.Once

var count = true

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(func() {
		count = false
	})
	if count {
		println("count is true")
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}
