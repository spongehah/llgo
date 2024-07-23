package main

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/libuv"
	"github.com/goplus/llgo/c/os"
	"unsafe"
)

const BUFFER_SIZE = 1024
const FILE_NAME = "example.txt"

var (
	loop     *libuv.Loop
	openReq  libuv.Fs
	readReq  libuv.Fs
	closeReq libuv.Fs

	buffer [BUFFER_SIZE]c.Char
	iov    libuv.Buf
)

func main() {
	c.Printf(c.Str("libuv version: %d\n"), libuv.Version())

	loop = libuv.DefaultLoop()

	libuv.FsOpen(loop, &openReq, c.Str(FILE_NAME), os.O_RDONLY, 0, onOpen)

	// Final run to ensure all callbacks are processed
	libuv.UvRun(loop, libuv.RUN_DEFAULT)

	libuv.FsReqCleanup(&openReq)
	libuv.FsReqCleanup(&readReq)
	libuv.FsReqCleanup(&closeReq)
	loop.Close()
}

func onOpen(req *libuv.Fs) {
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Error opening file: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
		loop.Stop()
		return
	}
	iov = libuv.InitBuf((*c.Char)(unsafe.Pointer(&buffer[0])), c.Uint(unsafe.Sizeof(buffer)))
	read := libuv.FsRead(loop, &readReq, libuv.UvFile(req.GetResult()), &iov, 1, -1, onRead)
	if read != 0 {
		c.Printf(c.Str("Error in FsRead: %s (code: %d)\n"), libuv.Strerror(c.Int(read)), read)
		loop.Stop()
		return
	}
}

func onRead(req *libuv.Fs) {
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Read error: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
		loop.Stop()
	} else if req.GetResult() == 0 {
		close := libuv.FsClose(loop, &closeReq, libuv.UvFile(openReq.GetResult()), onClose)
		if close != 0 {
			c.Printf(c.Str("Error in FsClose: %s (code: %d)\n"), libuv.Strerror(c.Int(close)), close)
			loop.Stop()
			return
		}
	} else {
		c.Printf(c.Str("Read %d bytes\n"), req.GetResult())
		c.Printf(c.Str("Read content: %.*s\n"), c.Int(req.GetResult()), (*c.Char)(unsafe.Pointer(&buffer[0])))
		loop.Stop()
	}
}

func onClose(req *libuv.Fs) {
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Error closing file: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
	} else {
		c.Printf(c.Str("\nFile closed successfully.\n"))
	}
	loop.Stop()
}
