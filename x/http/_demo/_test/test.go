package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
)

func main() {
	host := "localhost"
	port := "8080"
	var hints net.AddrInfo
	c.Memset(c.Pointer(&hints), 0, unsafe.Sizeof(hints))
	hints.Family = syscall.AF_UNSPEC
	hints.SockType = syscall.SOCK_STREAM

	var res *net.AddrInfo
	status := net.Getaddrinfo(c.AllocaCStr(host), c.AllocaCStr(port), &hints, &res)
	if status != 0 {
		fmt.Println("getaddrinfo error")
		return
	}
	fmt.Println("end")
}
