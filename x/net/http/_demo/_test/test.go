package main

func main() {
	ch := make(chan struct{}, 1)

	close(ch)

	select {
	case <-ch:
		println(1)
	}
}
