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

	wd := os.Getcwd(c.Alloca(os.PATH_MAX), os.PATH_MAX)
	c.Printf(c.Str("Current working directory: %s\n"), wd)

	// Check if file exists
	fileInfo := new(os.StatT)
	existResult := os.Stat(c.Str(FILE_NAME), fileInfo)
	if existResult != 0 {
		c.Printf(c.Str("Error: File '%s' does not exist or cannot be accessed. Error code: %d\n"), c.Str(FILE_NAME), existResult)
		return
	}
	c.Printf(c.Str("File '%s' exists and is accessible\n"), c.Str(FILE_NAME))

	// Try synchronous read
	fd := os.Open(c.Str(FILE_NAME), os.O_RDONLY, 0)
	if fd < 0 {
		c.Printf(c.Str("Error opening file synchronously: %d\n"), fd)
		return
	}
	syncBuffer := make([]c.Char, BUFFER_SIZE)
	bytesRead := os.Read(fd, unsafe.Pointer(&syncBuffer[0]), BUFFER_SIZE)
	if bytesRead < 0 {
		c.Printf(c.Str("Error reading file synchronously: %d\n"), bytesRead)
	} else {
		c.Printf(c.Str("Synchronous read successful. Bytes read: %d\n"), bytesRead)
		c.Printf(c.Str("File content: %.*s\n"), c.Int(bytesRead), &syncBuffer[0])
	}
	os.Close(fd)

	c.Printf(c.Str("Before FsOpen\n"))
	open := libuv.FsOpen(loop, &openReq, c.Str(FILE_NAME), os.O_RDONLY, 0, onOpen)
	c.Printf(c.Str("After FsOpen, result: %d\n"), open)

	c.Printf(c.Str("Before Run\n"))

	// Manually run the event loop a few times
	for i := 0; i < 5; i++ {
		run := libuv.UvRun(loop, libuv.RUN_NOWAIT)
		c.Printf(c.Str("Run iteration %d, result: %d\n"), i, run)
		if run == 0 {
			break
		}
	}

	c.Printf(c.Str("After manual runs\n"))

	// Final run to ensure all callbacks are processed
	run := libuv.UvRun(loop, libuv.RUN_DEFAULT)
	c.Printf(c.Str("Final Run, result: %d\n"), run)

	libuv.FsReqCleanup(&openReq)
	libuv.FsReqCleanup(&readReq)
	libuv.FsReqCleanup(&closeReq)
	loop.Close()
	c.Printf(c.Str("End of main\n"))
}

func onOpen(req *libuv.Fs) {
	c.Printf(c.Str("onOpen called, result: %d\n"), req.GetResult())
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Error opening file: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
		loop.Stop()
		return
	}
	iov = libuv.InitBuf((*c.Char)(unsafe.Pointer(&buffer[0])), c.Uint(unsafe.Sizeof(buffer)))
	read := libuv.FsRead(loop, &readReq, libuv.UvFile(req.GetResult()), []libuv.Buf{iov}, 1, -1, onRead)
	if read != 0 {
		c.Printf(c.Str("Error in FsRead: %s (code: %d)\n"), libuv.Strerror(c.Int(read)), read)
		loop.Stop()
		return
	}
	c.Printf(c.Str("FsRead called, result: %d\n"), read)
}

func onRead(req *libuv.Fs) {
	c.Printf(c.Str("onRead called, result: %d\n"), req.GetResult())
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Read error: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
		loop.Stop()
	} else if req.GetResult() == 0 {
		c.Printf(c.Str("End of file reached\n"))
		close := libuv.FsClose(loop, &closeReq, libuv.UvFile(openReq.GetResult()), onClose)
		if close != 0 {
			c.Printf(c.Str("Error in FsClose: %s (code: %d)\n"), libuv.Strerror(c.Int(close)), close)
			loop.Stop()
			return
		}
		c.Printf(c.Str("FsClose called, result: %d\n"), close)
	} else {
		c.Printf(c.Str("Read %d bytes\n"), req.GetResult())
		c.Printf(c.Str("Read content: %.*s\n"), c.Int(req.GetResult()), (*c.Char)(unsafe.Pointer(&buffer[0])))
		loop.Stop()
	}
}

func onClose(req *libuv.Fs) {
	c.Printf(c.Str("onClose called, result: %d\n"), req.GetResult())
	if req.GetResult() < 0 {
		c.Fprintf(c.Stderr, c.Str("Error closing file: %s\n"), libuv.Strerror(c.Int(req.GetResult())))
	} else {
		c.Printf(c.Str("\nFile closed successfully.\n"))
	}
	loop.Stop()
}
