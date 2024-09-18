package main

import "github.com/goplus/llgo/c"

func main() {
	count := c.Sysconf(c.Int(58))
	println(count)
}
