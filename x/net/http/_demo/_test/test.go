package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/textproto"
	"net/url"
	"strconv"
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/libuv"
	cnet "github.com/goplus/llgo/c/net"
	"github.com/goplus/llgo/c/syscall"
	"github.com/goplus/llgoexamples/rust/hyper"
)

func main() {
	t := &Transport{}
	t.loop = libuv.LoopNew()
	t.async = &libuv.Async{}
	t.exec = hyper.NewExecutor()

	t.loop.Async(t.async, nil)

	checker := &libuv.Check{}
	libuv.InitCheck(t.loop, checker)
	(*libuv.Handle)(c.Pointer(checker)).SetData(c.Pointer(t))
	checker.Start(readWriteLoop)

	go t.loop.Run(libuv.RUN_DEFAULT)

	host := "www.baidu.com"
	port := "80"
	conn := new(connData)
	libuv.InitTcp(t.loop, &conn.TcpHandle)
	(*libuv.Handle)(c.Pointer(&conn.TcpHandle)).SetData(c.Pointer(conn))

	var hints cnet.AddrInfo
	c.Memset(c.Pointer(&hints), 0, unsafe.Sizeof(hints))
	hints.Family = syscall.AF_UNSPEC
	hints.SockType = syscall.SOCK_STREAM

	var res *cnet.AddrInfo
	status := cnet.Getaddrinfo(c.AllocaCStr(host), c.AllocaCStr(port), &hints, &res)
	if status != 0 {
		fmt.Printf("getaddrinfo error\n")
	}

	(*libuv.Req)(c.Pointer(&conn.ConnectReq)).SetData(c.Pointer(conn))
	status = libuv.TcpConnect(&conn.ConnectReq, &conn.TcpHandle, res.Addr, onConnect)
	if status != 0 {
		fmt.Printf("connect error: %s\n", c.GoString(libuv.Strerror(libuv.Errno(status))))
	}
	cnet.Freeaddrinfo(res)

	// Hookup the IO
	hyperIo := newIoWithConnReadWrite(conn)
	// We need an executor generally to poll futures
	// Prepare client options
	opts := hyper.NewClientConnOptions()
	opts.Exec(t.exec)
	// send the handshake
	handshakeTask := hyper.Handshake(hyperIo, opts)
	var handshake = handshake
	handshakeTask.SetUserdata(c.Pointer(uintptr(handshake)))
	resc := make(chan responseAndError, 1)
	t.resc = resc

	// Send the request to readWriteLoop().
	t.exec.Push(handshakeTask)
	t.async.Send()

	select {
	case res := <-resc:
		resp := res.res
		fmt.Println(resp.Status)
		//resp.PrintHeaders()
		if resp.Body == nil {
			fmt.Println("respBody(reader) is nil")
			return
		} else {
			fmt.Println("respBody(reader) is present")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}
		fmt.Println(string(body))
	}
}

const debugSwitch = false

type Transport struct {
	exec       *hyper.Executor
	loop       *libuv.Loop
	async      *libuv.Async
	resc       chan responseAndError
	bodyWriter *io.PipeWriter
	nwrite     uintptr
}

// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).
type incomparable [0]func()

// responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.
type responseAndError struct {
	_   incomparable
	res *Response // else use this response (see res method)
	err error
}

type Header map[string][]string

type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           Header
	Body             io.ReadCloser
	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	//Form             url.Values
	//PostForm         url.Values
	//MultipartForm    *multipart.Form
	Trailer    Header
	RemoteAddr string
	RequestURI string
	//TLS              *tls.ConnectionState
	Cancel   <-chan struct{}
	Response *Response
}

type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	//Trailer          Header
	Request *Request
}

type connData struct {
	TcpHandle     libuv.Tcp
	ConnectReq    libuv.Connect
	ReadBuf       libuv.Buf
	ReadBufFilled uintptr
	nwrite        int64 // bytes written(Replaced from persistConn's nwrite)
	ReadWaker     *hyper.Waker
	WriteWaker    *hyper.Waker
}

// taskId The unique identifier of the next task polled from the executor
type taskId c.Int

const (
	notSet taskId = iota
	handshake
	read
	readDone
)

func readWriteLoop(checker *libuv.Check) {
	t := (*Transport)((*libuv.Handle)(c.Pointer(checker)).GetData())

	const debugReadWriteLoop = true // Debug switch provided for developers

	// The polling state machine!
	// Poll all ready tasks and act on them...
	for {
		task := t.exec.Poll()
		if debugSwitch {
			println("polling")
		}
		if task == nil {
			return
		}
		taskId := taskId(uintptr(task.Userdata()))
		if debugReadWriteLoop {
			println("taskId: ", taskId)
		}
		switch taskId {
		case handshake:
			if debugReadWriteLoop {
				println("write")
			}

			err := checkTaskType(task, handshake)
			if err != nil {
				fmt.Println(err)
				return
			}

			client := (*hyper.ClientConn)(task.Value())
			task.Free()

			taskId = read
			// Prepare the hyper request
			hyperReq := hyper.NewRequest()

			// Set the request line, default HTTP/1.1
			if hyperReq.SetMethod(&[]byte("GET")[0], c.Strlen(c.Str("GET"))) != hyper.OK {
				fmt.Printf("error setting method %s\n", "GET")
			}
			ruri := "/"
			if hyperReq.SetURI(&[]byte(ruri)[0], c.Strlen(c.AllocaCStr(ruri))) != hyper.OK {
				fmt.Printf("error setting uri %s\n", ruri)
			}
			if hyperReq.SetVersion(c.Int(hyper.HTTPVersion11)) != hyper.OK {
				fmt.Printf("error setting httpversion %s\n", "HTTP/1.1")
			}

			// Set the request headers
			reqHeaders := hyperReq.Headers()
			host := "www.baidu.com"
			if reqHeaders.Set(&[]byte("Host")[0], c.Strlen(c.Str("Host")), &[]byte(host)[0], c.Strlen(c.AllocaCStr(host))) != hyper.OK {
				fmt.Printf("error setting header: Host: %s\n", host)
			}
			// Send it!
			sendTask := client.Send(hyperReq)
			var read = read
			sendTask.SetUserdata(c.Pointer(uintptr(read)))
			sendRes := t.exec.Push(sendTask)
			if sendRes != hyper.OK {
				err = errors.New("failed to send the request")
			}

			if debugReadWriteLoop {
				println("write end")
			}
		case read:
			if debugReadWriteLoop {
				println("read")
			}

			err := checkTaskType(task, read)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Take the results
			hyperResp := (*hyper.Response)(task.Value())
			task.Free()

			rp := hyperResp.ReasonPhrase()
			rpLen := hyperResp.ReasonPhraseLen()

			resp := &Response{}
			// Parse the first line of the response.
			resp.Status = strconv.Itoa(int(hyperResp.Status())) + " " + c.GoString((*int8)(c.Pointer(rp)), rpLen)
			resp.StatusCode = int(hyperResp.Status())
			version := int(hyperResp.Version())
			resp.ProtoMajor, resp.ProtoMinor = splitTwoDigitNumber(version)
			resp.Proto = fmt.Sprintf("HTTP/%d.%d", resp.ProtoMajor, resp.ProtoMinor)
			println("read resp line")

			headers := hyperResp.Headers()
			headers.Foreach(appendToResponseHeader, c.Pointer(resp))
			println("read resp headers")

			resp.Body, t.bodyWriter = io.Pipe()
			respBody := hyperResp.Body()

			// No longer need the response
			hyperResp.Free()

			bodyForeachTask := respBody.Foreach(appendToResponseBody, c.Pointer(t))
			var readDone = readDone
			bodyForeachTask.SetUserdata(c.Pointer(uintptr(readDone)))
			t.exec.Push(bodyForeachTask)
			println("read resp body")

			t.resc <- responseAndError{res: resp}

			if debugReadWriteLoop {
				println("read end")
			}
		case readDone:
			// A background task of reading the response body is completed
			if debugReadWriteLoop {
				println("readDone")
			}
			if t.bodyWriter != nil {
				println("close body writer")
				t.bodyWriter.Close()
			}
			err := checkTaskType(task, readDone)
			if err != nil {
				fmt.Println(err)
				return
			}

			if task.Type() != hyper.TaskEmpty {
				fmt.Println("not empty task")
				return
			}
			// free the task
			task.Free()

			//testHookReadLoopBeforeNextRead()
			if debugReadWriteLoop {
				println("readDone end")
			}
		case notSet:
			if debugReadWriteLoop {
				println("notSet")
			}
			// A background task for hyper_client completed...
			task.Free()
		}
	}
}

// ----------------------------------------------------------

// appendToResponseHeader (HeadersForEachCallback) prints each header to the console
func appendToResponseHeader(userdata c.Pointer, name *uint8, nameLen uintptr, value *uint8, valueLen uintptr) c.Int {
	resp := (*Response)(userdata)
	nameStr := c.GoString((*int8)(c.Pointer(name)), nameLen)
	valueStr := c.GoString((*int8)(c.Pointer(value)), valueLen)

	if resp.Header == nil {
		resp.Header = make(Header)
	}
	resp.Header.Add(nameStr, valueStr)
	return hyper.IterContinue
}

// appendToResponseBody BodyForeachCallback function: Process the response body
func appendToResponseBody(userdata c.Pointer, chunk *hyper.Buf) c.Int {
	println("appendToResponseBody start")
	t := (*Transport)(userdata)
	writer := t.bodyWriter
	if chunk == nil {
		println("appendToResponseBody chunk is nil")
	} else {
		println("appendToResponseBody chunk is present")
	}
	bufLen := chunk.Len()
	bytes := unsafe.Slice(chunk.Bytes(), bufLen)
	if writer == nil {
		println("appendToResponseBody writer is nil")
	} else {
		println("appendToResponseBody writer is present")
	}
	if bytes == nil {
		println("appendToResponseBody bytes is nil")
	} else {
		println("appendToResponseBody bytes is present")
	}
	_, err := writer.Write(bytes)
	println("write: ", bufLen, "total write: ", t.nwrite)
	if err != nil {
		fmt.Println("Error writing to response body:", err)
		writer.Close()
		return hyper.IterBreak
	}
	t.nwrite += bufLen
	println("appendToResponseBody end")
	return hyper.IterContinue
}

// onConnect is the libuv callback for a successful connection
func onConnect(req *libuv.Connect, status c.Int) {
	if debugSwitch {
		println("connect start")
		defer println("connect end")
	}
	conn := (*connData)((*libuv.Req)(c.Pointer(req)).GetData())

	if status < 0 {
		c.Fprintf(c.Stderr, c.Str("connect error: %d\n"), libuv.Strerror(libuv.Errno(status)))
		return
	}
	(*libuv.Stream)(c.Pointer(&conn.TcpHandle)).StartRead(allocBuffer, onRead)
}

// allocBuffer allocates a buffer for reading from a socket
func allocBuffer(handle *libuv.Handle, suggestedSize uintptr, buf *libuv.Buf) {
	conn := (*connData)(handle.GetData())
	if conn.ReadBuf.Base == nil {
		conn.ReadBuf = libuv.InitBuf((*c.Char)(c.Malloc(suggestedSize)), c.Uint(suggestedSize))
		//base := make([]byte, suggestedSize)
		//conn.ReadBuf = libuv.InitBuf((*c.Char)(c.Pointer(&base[0])), c.Uint(suggestedSize))
		conn.ReadBufFilled = 0
	}
	*buf = libuv.InitBuf((*c.Char)(c.Pointer(uintptr(c.Pointer(conn.ReadBuf.Base))+conn.ReadBufFilled)), c.Uint(suggestedSize-conn.ReadBufFilled))
}

// onRead is the libuv callback for reading from a socket
// This callback function is called when data is available to be read
func onRead(stream *libuv.Stream, nread c.Long, buf *libuv.Buf) {
	conn := (*connData)((*libuv.Handle)(c.Pointer(stream)).GetData())
	// If data was read (nread > 0)
	if nread > 0 {
		// Update the amount of filled buffer
		conn.ReadBufFilled += uintptr(nread)
	}
	// If there's a pending read waker
	if conn.ReadWaker != nil {
		// Wake up the pending read operation of Hyper
		conn.ReadWaker.Wake()
		// Clear the waker reference
		conn.ReadWaker = nil
	}
}

// readCallBack read callback function for Hyper library
func readCallBack(userdata c.Pointer, ctx *hyper.Context, buf *uint8, bufLen uintptr) uintptr {
	conn := (*connData)(userdata)
	// If there's data in the buffer
	if conn.ReadBufFilled > 0 {
		// Calculate how much data to copy (minimum of filled amount and requested amount)
		var toCopy uintptr
		if bufLen < conn.ReadBufFilled {
			toCopy = bufLen
		} else {
			toCopy = conn.ReadBufFilled
		}
		// Copy data from read buffer to Hyper's buffer
		c.Memcpy(c.Pointer(buf), c.Pointer(conn.ReadBuf.Base), toCopy)
		// Move remaining data to the beginning of the buffer
		c.Memmove(c.Pointer(conn.ReadBuf.Base), c.Pointer(uintptr(c.Pointer(conn.ReadBuf.Base))+toCopy), conn.ReadBufFilled-toCopy)
		// Update the amount of filled buffer
		conn.ReadBufFilled -= toCopy
		// Return the number of bytes copied
		return toCopy
	}

	// If no data in buffer, set up a waker to wait for more data
	// Free the old waker if it exists
	if conn.ReadWaker != nil {
		conn.ReadWaker.Free()
	}
	// Create a new waker
	conn.ReadWaker = ctx.Waker()
	// Return HYPER_IO_PENDING to indicate operation is pending, waiting for more data
	return hyper.IoPending
}

// onWrite is the libuv callback for writing to a socket
// Callback function called after a write operation completes
func onWrite(req *libuv.Write, status c.Int) {
	conn := (*connData)((*libuv.Req)(c.Pointer(req)).GetData())
	// If there's a pending write waker
	if conn.WriteWaker != nil {
		// Wake up the pending write operation
		conn.WriteWaker.Wake()
		// Clear the waker reference
		conn.WriteWaker = nil
	}
}

// writeCallBack write callback function for Hyper library
func writeCallBack(userdata c.Pointer, ctx *hyper.Context, buf *uint8, bufLen uintptr) uintptr {
	conn := (*connData)(userdata)
	// Create a libuv buffer
	initBuf := libuv.InitBuf((*c.Char)(c.Pointer(buf)), c.Uint(bufLen))
	req := &libuv.Write{}
	// Associate the connection data with the write request
	(*libuv.Req)(c.Pointer(req)).SetData(c.Pointer(conn))

	// Perform the asynchronous write operation
	ret := req.Write((*libuv.Stream)(c.Pointer(&conn.TcpHandle)), &initBuf, 1, onWrite)
	// If the write operation was successfully initiated
	if ret >= 0 {
		conn.nwrite += int64(bufLen)
		// Return the number of bytes to be written
		return bufLen
	}

	// If the write operation can't complete immediately, set up a waker to wait for completion
	if conn.WriteWaker != nil {
		// Free the old waker if it exists
		conn.WriteWaker.Free()
	}
	// Create a new waker
	conn.WriteWaker = ctx.Waker()
	// Return HYPER_IO_PENDING to indicate operation is pending, waiting for write to complete
	return hyper.IoPending
}

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
// The key is case insensitive; it is canonicalized by
// CanonicalHeaderKey.
func (h Header) Add(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

func (resp *Response) PrintHeaders() {
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}

func (conn *connData) Close() error {
	if conn == nil {
		return nil
	}
	if conn.ReadWaker != nil {
		conn.ReadWaker.Free()
		conn.ReadWaker = nil
	}
	if conn.WriteWaker != nil {
		conn.WriteWaker.Free()
		conn.WriteWaker = nil
	}
	if conn.ReadBuf.Base != nil {
		c.Free(c.Pointer(conn.ReadBuf.Base))
		conn.ReadBuf.Base = nil
	}
	(*libuv.Handle)(c.Pointer(&conn.TcpHandle)).Close(nil)
	return nil
}

// newIoWithConnReadWrite creates a new IO with read and write callbacks
func newIoWithConnReadWrite(connData *connData) *hyper.Io {
	hyperIo := hyper.NewIo()
	hyperIo.SetUserdata(c.Pointer(connData))
	hyperIo.SetRead(readCallBack)
	hyperIo.SetWrite(writeCallBack)
	return hyperIo
}

// checkTaskType checks the task type
func checkTaskType(task *hyper.Task, curTaskId taskId) error {
	switch curTaskId {
	case handshake:
		if task.Type() == hyper.TaskError {
			log.Printf("[readWriteLoop::handshake]handshake task error!\n")
			return fail((*hyper.Error)(task.Value()))
		}
		if task.Type() != hyper.TaskClientConn {
			return fmt.Errorf("[readWriteLoop::handshake]unexpected task type\n")
		}
		return nil
	case read:
		if task.Type() == hyper.TaskError {
			log.Printf("[readWriteLoop::read]write task error!\n")
			return fail((*hyper.Error)(task.Value()))
		}
		if task.Type() != hyper.TaskResponse {
			c.Printf(c.Str("[readWriteLoop::read]unexpected task type\n"))
			return errors.New("[readWriteLoop::read]unexpected task type\n")
		}
		return nil
	case readDone:
		if task.Type() == hyper.TaskError {
			log.Printf("[readWriteLoop::readDone]read response body error!\n")
			return fail((*hyper.Error)(task.Value()))
		}
		return nil
	case notSet:
	}
	return errors.New("[readWriteLoop]unexpected task type\n")
}

// fail prints the error details and panics
func fail(err *hyper.Error) error {
	if err != nil {
		c.Printf(c.Str("[readWriteLoop]error code: %d\n"), err.Code())
		// grab the error details
		var errBuf [256]c.Char
		errLen := err.Print((*uint8)(c.Pointer(&errBuf[:][0])), uintptr(len(errBuf)))

		c.Printf(c.Str("[readWriteLoop]details: %.*s\n"), c.Int(errLen), c.Pointer(&errBuf[:][0]))

		// clean up the error
		err.Free()
		return fmt.Errorf("[readWriteLoop]hyper request error, error code: %d\n", int(err.Code()))
	}
	return nil
}

// splitTwoDigitNumber splits a two-digit number into two digits.
func splitTwoDigitNumber(num int) (int, int) {
	tens := num / 10
	ones := num % 10
	return tens, ones
}
