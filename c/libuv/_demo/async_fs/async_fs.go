package main

import (
	"fmt"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/libuv"
	"github.com/goplus/llgo/c/os"
)

const BUFFER_SIZE = 1024

var (
	loop     *libuv.Loop
	openReq  libuv.Fs
	readReq  libuv.Fs
	closeReq libuv.Fs

	buffer [BUFFER_SIZE]c.Char
	iov    libuv.Buf
)

func main() {
	fmt.Printf("libuv version: %d\n", libuv.Version())

	// Initialize the loop
	loop = libuv.DefaultLoop()

	// Open the file
	openReq.Open(loop, "example.txt", os.O_RDONLY, 0, onOpen)

	// Run the loop
	loop.Run(libuv.RUN_DEFAULT)

	// Cleanup
	defer cleanup()
}

func onOpen(req *libuv.Fs) {
	// Check for errors
	if req.GetResult() < 0 {
		fmt.Printf("Error opening file: %s\n", libuv.Strerror(req.GetResult()))
		loop.Stop()
		return
	}

	// Init buffer
	iov = libuv.InitBuf(buffer[:])
	// Read the file
	readRes := readReq.Read(loop, req.GetResult(), &iov, 1, -1, onRead)
	if readRes != 0 {
		fmt.Printf("Error in FsRead: %s (code: %d)\n", libuv.Strerror(readRes), readRes)
		loop.Stop()
		return
	}
}

func onRead(req *libuv.Fs) {
	// Check for errors
	if req.GetResult() < 0 {
		fmt.Printf("Read error: %s\n", libuv.Strerror(req.GetResult()))
		loop.Stop()
	} else if req.GetResult() == 0 {
		fmt.Printf("EOF\n")
		// Close the file
		closeRes := closeReq.Close(loop, openReq.GetResult(), onClose)
		if closeRes != 0 {
			fmt.Printf("Error in FsClose: %s (code: %d)\n", libuv.Strerror(closeRes), closeRes)
			loop.Stop()
			return
		}
	} else {
		// Print the content
		fmt.Printf("Read %d bytes\n", req.GetResult())
		fmt.Printf("Read content: %.*s\n",
			req.GetResult(),
			(*[BUFFER_SIZE]byte)(unsafe.Pointer(&buffer[0]))[:])
		loop.Stop()
	}
}

func onClose(req *libuv.Fs) {
	// Check for errors
	if req.GetResult() < 0 {
		fmt.Printf("Error closing file: %s\n", libuv.Strerror(req.GetResult()))
	} else {
		fmt.Printf("\nFile closed successfully.\n")
	}
	loop.Stop()
}

func cleanup() {
	// Cleanup the requests
	libuv.FsReqCleanup(&openReq)
	libuv.FsReqCleanup(&readReq)
	libuv.FsReqCleanup(&closeReq)
	// Close the loop
	loop.Close()
}
