package libuv

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
	"unsafe"
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs libuv); -luv"
)

// ----------------------------------------------
const (
	RUN_DEFAULT RunMode = iota
	RUN_ONCE
	RUN_NOWAIT
)

const (
	LOOP_BLOCK_SIGNAL LoopOption = iota
	METRICS_IDLE_TIME
)

const (
	UV_LEAVE_GROUP Membership = iota
	UV_JOIN_GROUP
)

const (
	UNKNOWN_HANDLE HandleType = iota
	ASYNC
	CHECK
	FS_EVENT
	FS_POLL
	HANDLE
	IDLE
	NAMED_PIPE
	POLL
	PREPARE
	PROCESS
	STREAM
	TCP
	TIMER
	TTY
	UDP
	SIGNAL
	FILE
	HANDLE_TYPE_MAX
)

const (
	UNKNOWN_REQ ReqType = iota
	REQ
	CONNECT
	WRITE
	SHUTDOWN
	UDP_SEND
	FS
	WORK
	GETADDRINFO
	GETNAMEINFO
	RANDOM
	REQ_TYPE_PRIVATE
	REQ_TYPE_MAX
)

const (
	READABLE PollEvent = 1 << iota
	WRITABLE
	DISCONNECT
	PRIPRIORITIZED
)

type RunMode c.Int

type LoopOption c.Int

type Membership c.Int

type HandleType c.Int

type ReqType c.Int

type OsSock c.Int

type OsFd c.Int

type PollEvent c.Int

// ----------------------------------------------

/* Handle types. */

type Loop struct {
	Unused [0]byte
}

type Handle struct {
	Unused [96]byte
}

type Dir struct {
	Unused [0]byte
}

type Stream struct {
	Unused [264]byte
}

type Pipe struct {
	Unused [0]byte
}

type Tty struct {
	Unused [0]byte
}

type Poll struct {
	Unused [0]byte
}

type Prepare struct {
	Unused [0]byte
}

type Check struct {
	Unused [0]byte
}

type Idle struct {
	Unused [0]byte
}

type Async struct {
	Unused [0]byte
}

type Process struct {
	Unused [0]byte
}

/* Request types. */

type Req struct {
	Unused [0]byte
}

type GetAddrInfo struct {
	Unused [0]byte
}

type GetNameInfo struct {
	Unused [0]byte
}

type Shutdown struct {
	Unused [0]byte
}

type Write struct {
	Unused [192]byte
}

type Connect struct {
	Unused [0]byte
}

type Buf struct {
	Base *c.Char
	Len  uintptr
} // ----------------------------------------------

/* Function type */

// llgo:type C
type MallocFunc func(size uintptr) c.Pointer

// llgo:type C
type ReallocFunc func(ptr c.Pointer, size uintptr) c.Pointer

// llgo:type C
type CallocFunc func(count uintptr, size uintptr) c.Pointer

// llgo:type C
type FreeFunc func(ptr c.Pointer)

// llgo:type C
type AllocCb func(handle *Handle, suggestedSize uintptr, buf *Buf)

// llgo:type C
type ReadCb func(stream *Stream, nread c.Long, buf *Buf)

// llgo:type C
type WriteCb func(req *Write, status c.Int)

// llgo:type C
type GetaddrinfoCb func(req *GetAddrInfo, status c.Int, res *net.AddrInfo)

// llgo:type C
type GetnameinfoCb func(req *GetNameInfo, status c.Int, hostname *c.Char, service *c.Char)

// llgo:type C
type ConnectionCb func(server *Stream, status c.Int)

// llgo:type C
type ShutdownCb func(req *Shutdown, status c.Int)

// llgo:type C
type WalkCb func(handle *Handle, arg c.Pointer)

// llgo:type C
type PollCb func(handle *Poll, status c.Int, events c.Int)

// ----------------------------------------------

//go:linkname UvVersion C.uv_version
func UvVersion() c.Uint

//go:linkname UvVersionString C.uv_version_string
func UvVersionString() *c.Char

//go:linkname UvLibraryShutdown C.uv_library_shutdown
func UvLibraryShutdown()

//go:linkname UvReplaceAllocator C.uv_replace_allocator
func UvReplaceAllocator(mallocFunc MallocFunc, reallocFunc ReallocFunc, callocFunc CallocFunc, freeFunc FreeFunc) c.Int

// Version returns the version of libuv.
func Version() int {
	return int(UvVersion())
}

// VersionString returns the version string of libuv.
func VersionString() string {
	return c.GoString(UvVersionString())
}

// LibraryShutdown shuts down the libuv library.
func LibraryShutdown() {
	UvLibraryShutdown()
}

// ReplaceAllocator replaces the allocator functions used by libuv.
func ReplaceAllocator(mallocFunc MallocFunc, reallocFunc ReallocFunc, callocFunc CallocFunc, freeFunc FreeFunc) int {
	return int(UvReplaceAllocator(mallocFunc, reallocFunc, callocFunc, freeFunc))
}

// ----------------------------------------------

// llgo:link (*Shutdown).Shutdown C.uv_shutdown
func (shutdown *Shutdown) Shutdown(stream *Stream, shutdownCb ShutdownCb) c.Int {
	return 0
}

// ----------------------------------------------

/* Handle related function and method */

// llgo:link (*Handle).Ref C.uv_ref
func (handle *Handle) Ref() {}

// llgo:link (*Handle).Unref C.uv_unref
func (handle *Handle) Unref() {}

// llgo:link (*Handle).HasRef C.uv_has_ref
func (handle *Handle) HasRef() c.Int {
	return 0
}

//go:linkname HandleSize C.uv_handle_size
func HandleSize(handleType HandleType) uintptr

// llgo:link (*Handle).GetType C.uv_handle_get_type
func (handle *Handle) GetType() HandleType {
	return 0
}

//go:linkname HandleTypeName C.uv_handle_type_name
func HandleTypeName(handleType HandleType) *c.Char

// llgo:link (*Handle).GetData C.uv_handle_get_data
func (handle *Handle) GetData() c.Pointer {
	return nil
}

// llgo:link (*Handle).GetLoop C.uv_handle_get_loop
func (handle *Handle) GetLoop() *Loop {
	return nil
}

// llgo:link (*Handle).SetData C.uv_handle_set_data
func (handle *Handle) SetData(data c.Pointer) {}

// llgo:link (*Handle).IsActive C.uv_is_active
func (handle *Handle) IsActive() c.Int {
	return 0
}

// llgo:link (*Handle).Close C.uv_close
func (handle *Handle) Close(closeCb CloseCb) {}

// llgo:link (*Handle).SendBufferSize C.uv_send_buffer_size
func (handle *Handle) SendBufferSize(value *c.Int) c.Int {
	return 0
}

// llgo:link (*Handle).RecvBufferSize C.uv_recv_buffer_size
func (handle *Handle) RecvBufferSize(value *c.Int) c.Int {
	return 0
}

// llgo:link (*Handle).Fileno C.uv_fileno
func (handle *Handle) Fileno(fd *OsFd) c.Int {
	return 0
}

//go:linkname UvPipe C.uv_pipe
func UvPipe(fds [2]UvFile, readFlags c.Int, writeFlags c.Int) c.Int {
	return 0
}

//go:linkname Socketpair C.uv_socketpair
func Socketpair(_type c.Int, protocol c.Int, socketVector [2]OsSock, flag0 c.Int, flag1 c.Int) c.Int {
	return 0
}

// llgo:link (*Handle).IsClosing C.uv_is_closing
func (handle *Handle) IsClosing() c.Int {
	return 0
}

// UvRef increases the reference count of the handle.
func (handle *Handle) UvRef() {
	handle.Ref()
}

// UvUnRef decreases the reference count of the handle.
func (handle *Handle) UvUnRef() {
	handle.Unref()
}

// UvHasRef returns the reference count of the handle.
func (handle *Handle) UvHasRef() int {
	return int(handle.HasRef())
}

// UvHandleSize returns the size of the handle.
func UvHandleSize(handleType HandleType) uintptr {
	return HandleSize(handleType)
}

// UvGetType returns the type of the handle.
func (handle *Handle) UvGetType() HandleType {
	return handle.GetType()
}

// UvHandleTypeName returns the name of the handle type.
func UvHandleTypeName(handleType HandleType) string {
	return c.GoString(HandleTypeName(handleType))
}

// UvGetData returns the data of the handle.
func (handle *Handle) UvGetData() c.Pointer {
	return handle.GetData()
}

// UvGetLoop returns the loop of the handle.
func (handle *Handle) UvGetLoop() *Loop {
	return handle.GetLoop()
}

// UvSetData sets the data of the handle.
func (handle *Handle) UvSetData(data c.Pointer) {
	handle.SetData(data)
}

// UvIsActive returns the status of the handle.
func (handle *Handle) UvIsActive() int {
	return int(handle.IsActive())
}

// UvClose closes the handle.
func (handle *Handle) UvClose(closeCb CloseCb) {
	handle.Close(closeCb)
}

// UvSendBufferSize sets the send buffer size of the handle.
func (handle *Handle) UvSendBufferSize(value *c.Int) int {
	return int(handle.SendBufferSize(value))
}

// UvRecvBufferSize sets the receive buffer size of the handle.
func (handle *Handle) UvRecvBufferSize(value *c.Int) int {
	return int(handle.RecvBufferSize(value))
}

// UvFileno returns the file descriptor of the handle.
func (handle *Handle) UvFileno(fd *OsFd) int {
	return int(handle.Fileno(fd))
}

// UvUvPipe creates a pipe.
func UvUvPipe(fds [2]UvFile, readFlags int, writeFlags int) int {
	return int(UvPipe(fds, c.Int(readFlags), c.Int(writeFlags)))
}

// UvSocketpair creates a socket pair.
func UvSocketpair(_type int, protocol int, socketVector [2]OsSock, flag0 int, flag1 int) int {
	return int(Socketpair(c.Int(_type), c.Int(protocol), socketVector, c.Int(flag0), c.Int(flag1)))
}

// UvIsClosing returns the status of the handle.
func (handle *Handle) UvIsClosing() int {
	return int(handle.IsClosing())
}

// ----------------------------------------------

/* Req related function and method */

//go:linkname ReqSize C.uv_req_size
func ReqSize(reqType ReqType) uintptr

// llgo:link (*Req).GetData C.uv_req_get_data
func (req *Req) GetData() c.Pointer {
	return nil
}

// llgo:link (*Req).SetData C.uv_handle_set_data
func (req *Req) SetData(data c.Pointer) {}

// llgo:link (*Req).GetType C.uv_req_get_type
func (req *Req) GetType() ReqType {
	return 0
}

//go:linkname TypeName C.uv_req_type_name
func TypeName(reqType ReqType) *c.Char

// UvReqSize returns the size of the request.
func UvReqSize(reqType ReqType) uintptr {
	return ReqSize(reqType)
}

// UvGetData returns the data of the request.
func (req *Req) UvGetData() c.Pointer {
	return req.GetData()
}

// UvSetData sets the data of the request.
func (req *Req) UvSetData(data c.Pointer) {
	req.SetData(data)
}

// UvGetType returns the type of the request.
func (req *Req) UvGetType() ReqType {
	return req.GetType()
}

// UvTypeName returns the name of the request type.
func UvTypeName(reqType ReqType) string {
	return c.GoString(TypeName(reqType))
}

// ----------------------------------------------

/* Stream related function and method */

// llgo:link (*Stream).GetWriteQueueSize C.uv_stream_get_write_queue_size
func (stream *Stream) GetWriteQueueSize() uintptr {
	return 0
}

// llgo:link (*Stream).Listen C.uv_listen
func (stream *Stream) Listen(backlog c.Int, connectionCb ConnectionCb) c.Int {
	return 0
}

// llgo:link (*Stream).Accept C.uv_accept
func (server *Stream) Accept(client *Stream) c.Int {
	return 0
}

// llgo:link (*Stream).StartRead C.uv_read_start
func (stream *Stream) StartRead(allocCb AllocCb, readCb ReadCb) c.Int {
	return 0
}

// llgo:link (*Stream).StopRead C.uv_read_stop
func (stream *Stream) StopRead() c.Int {
	return 0
}

// llgo:link (*Write).Write C.uv_write
func (req *Write) Write(stream *Stream, bufs *Buf, nbufs c.Uint, writeCb WriteCb) c.Int {
	return 0
}

// llgo:link (*Write).Write2 C.uv_write2
func (req *Write) Write2(stream *Stream, bufs *Buf, nbufs c.Uint, sendStream *Stream, writeCb WriteCb) c.Int {
	return 0
}

// llgo:link (*Stream).TryWrite C.uv_try_write
func (stream *Stream) TryWrite(bufs *Buf, nbufs c.Uint) c.Int {
	return 0
}

// llgo:link (*Stream).TryWrite2 C.uv_try_write2
func (stream *Stream) TryWrite2(bufs *Buf, nbufs c.Uint, sendStream *Stream) c.Int {
	return 0
}

// llgo:link (*Stream).IsReadable C.uv_is_readable
func (stream *Stream) IsReadable() c.Int {
	return 0
}

// llgo:link (*Stream).IsWritable C.uv_is_writable
func (stream *Stream) IsWritable() c.Int {
	return 0
}

// llgo:link (*Stream).SetBlocking C.uv_stream_set_blocking
func (stream *Stream) SetBlocking(blocking c.Int) c.Int {
	return 0
}

// UvGetWriteQueueSize returns the size of the write queue of the stream.
func (stream *Stream) UvGetWriteQueueSize() uintptr {
	return stream.GetWriteQueueSize()
}

// UvListen listens on the stream.
func (stream *Stream) UvListen(backlog int, connectionCb ConnectionCb) int {
	return int(stream.Listen(c.Int(backlog), connectionCb))
}

// UvAccept accepts a connection on the stream.
func (server *Stream) UvAccept(client *Stream) int {
	return int(server.Accept(client))
}

// UvStartRead starts reading from the stream.
func (stream *Stream) UvStartRead(allocCb AllocCb, readCb ReadCb) int {
	return int(stream.StartRead(allocCb, readCb))
}

// UvStopRead stops reading from the stream.
func (stream *Stream) UvStopRead() int {
	return int(stream.StopRead())
}

// UvWrite writes to the stream.
func (req *Write) UvWrite(stream *Stream, bufs *Buf, nbufs uint, writeCb WriteCb) int {
	return int(req.Write(stream, bufs, c.Uint(nbufs), writeCb))
}

// UvWrite2 writes to the stream.
func (req *Write) UvWrite2(stream *Stream, bufs *Buf, nbufs uint, sendStream *Stream, writeCb WriteCb) int {
	return int(req.Write2(stream, bufs, c.Uint(nbufs), sendStream, writeCb))
}

// UvTryWrite writes to the stream.
func (stream *Stream) UvTryWrite(bufs *Buf, nbufs uint) int {
	return int(stream.TryWrite(bufs, c.Uint(nbufs)))
}

// UvTryWrite2 writes to the stream.
func (stream *Stream) UvTryWrite2(bufs *Buf, nbufs uint, sendStream *Stream) int {
	return int(stream.TryWrite2(bufs, c.Uint(nbufs), sendStream))
}

// UvIsReadable returns the status of the stream.
func (stream *Stream) UvIsReadable() int {
	return int(stream.IsReadable())
}

// UvIsWritable returns the status of the stream.
func (stream *Stream) UvIsWritable() int {
	return int(stream.IsWritable())
}

// UvSetBlocking sets the blocking status of the stream.
func (stream *Stream) UvSetBlocking(blocking int) int {
	return int(stream.SetBlocking(c.Int(blocking)))
}

// ----------------------------------------------

/* Loop related functions and method. */

//go:linkname UvLoopSize C.uv_loop_size
func UvLoopSize() uintptr

//go:linkname UvRun C.uv_run
func UvRun(loop *Loop, mode RunMode) c.Int

//go:linkname UvLoopAlive C.uv_loop_alive
func UvLoopAlive(loop *Loop) c.Int

//go:linkname UvLoopClose C.uv_loop_close
func UvLoopClose(loop *Loop) c.Int

//go:linkname UvLoopConfigure C.uv_loop_configure
func UvLoopConfigure(loop *Loop, option LoopOption, arg c.Int) c.Int

//go:linkname UvLoopDefault C.uv_default_loop
func UvLoopDefault() *Loop

//go:linkname UvLoopDelete C.uv_loop_delete
func UvLoopDelete(loop *Loop) c.Int

//go:linkname UvLoopFork C.uv_loop_fork
func UvLoopFork(loop *Loop) c.Int

//go:linkname UvLoopInit C.uv_loop_init
func UvLoopInit(loop *Loop) c.Int

//go:linkname UvLoopNew C.uv_loop_new
func UvLoopNew() *Loop

//go:linkname UvLoopNow C.uv_now
func UvLoopNow(loop *Loop) c.UlongLong

//go:linkname UvLoopUpdateTime C.uv_update_time
func UvLoopUpdateTime(loop *Loop)

//go:linkname UvLoopBackendFd C.uv_backend_fd
func UvLoopBackendFd(loop *Loop) c.Int

//go:linkname UvLoopBackendTimeout C.uv_backend_timeout
func UvLoopBackendTimeout(loop *Loop) c.Int

//go:linkname UvLoopWalk C.uv_walk
func UvLoopWalk(loop *Loop, walkCb WalkCb, arg c.Pointer)

// DefaultLoop returns the default loop.
func (l *Loop) DefaultLoop() *Loop {
	return UvLoopDefault()
}

// Size returns the size of the loop.
func (l *Loop) Size() uintptr {
	return UvLoopSize()
}

// Init initializes the loop.
func (l *Loop) Init() int {
	return int(UvLoopInit(l))
}

// Run runs the loop.
func (l *Loop) Run(mode RunMode) int {
	return int(UvRun(l, mode))
}

// Stop closes the loop.
func (l *Loop) Stop() int {
	return int(UvLoopClose(l))
}

// Default creates a new loop.
func (l *Loop) Default() *Loop {
	return UvLoopDefault()
}

// New creates a new loop.
func (l *Loop) New() *Loop {
	return UvLoopNew()
}

// Deprecated: use LoopClose instead.
// Delete closes the loop.
func (l *Loop) Delete() int {
	return int(UvLoopDelete(l))
}

// Alive returns the status of the loop.
func (l *Loop) Alive() int {
	return int(UvLoopAlive(l))
}

// Close closes the loop.
func (l *Loop) Close() int {
	return int(UvLoopClose(l))
}

// Configure configures the loop.
func (l *Loop) Configure(loop *Loop, option int, arg int) int {
	return int(UvLoopConfigure(l, LoopOption(c.Int(option)), c.Int(arg)))
}

// Walk walks the loop.
func (l *Loop) Walk(walkCb WalkCb, arg c.Pointer) {
	UvLoopWalk(l, walkCb, arg)
}

// Fork forks the loop.
func (l *Loop) Fork(loop *Loop) int {
	return int(UvLoopFork(l))
}

// UpdateTime updates the time of the loop.
func (l *Loop) UpdateTime() {
	UvLoopUpdateTime(l)
}

// Now returns the current time of the loop.
func (l *Loop) Now() uint64 {
	return uint64(UvLoopNow(l))
}

// BackendFd returns the backend file descriptor of the loop.
func (l *Loop) BackendFd() int {
	return int(UvLoopBackendFd(l))
}

// BackendTimeout returns the backend timeout of the loop.
func (l *Loop) BackendTimeout() int {
	return int(UvLoopBackendTimeout(l))
}

// ----------------------------------------------

/* Buf related functions and method. */

//go:linkname UvInitBuf C.uv_buf_init
func UvInitBuf(base *c.Char, len c.Uint) Buf

// InitBuf initializes a buffer with the given c.Char slice.
func InitBuf(buffer []c.Char) Buf {
	return UvInitBuf((*c.Char)(unsafe.Pointer(&buffer[0])), c.Uint(unsafe.Sizeof(buffer)))
}

// ----------------------------------------------

/* Poll related function and method */

//go:linkname PollInit C.uv_poll_init
func PollInit(loop *Loop, handle *Poll, fd OsFd) c.Int

//go:linkname PollStart C.uv_poll_start
func PollStart(handle *Poll, events c.Int, cb PollCb) c.Int

//go:linkname PollStop C.uv_poll_stop
func PollStop(handle *Poll) c.Int

// Init initializes the poll handle with the given file descriptor.
func (p *Poll) Init(loop *Loop, fd OsFd) int {
	return int(PollInit(loop, p, fd))
}

// Start starts polling the file descriptor.
func (p *Poll) Start(events int, cb PollCb) int {
	return int(PollStart(p, c.Int(events), cb))
}

// Stop stops polling the file descriptor.
func (p *Poll) Stop() int {
	return int(PollStop(p))
}

// ----------------------------------------------

/* Getaddrinfo related function and method */

//go:linkname Getaddrinfo C.uv_getaddrinfo
func Getaddrinfo(loop *Loop, req *GetAddrInfo, getaddrinfoCb GetaddrinfoCb, node *c.Char, service *c.Char, hints *net.AddrInfo) c.Int

//go:linkname Freeaddrinfo C.uv_freeaddrinfo
func Freeaddrinfo(addrInfo *net.AddrInfo)

// UvGetaddrinfo gets the address information.
func UvGetaddrinfo(loop *Loop, req *GetAddrInfo, getaddrinfoCb GetaddrinfoCb, node string, service string, hints *net.AddrInfo) int {
	return int(Getaddrinfo(loop, req, getaddrinfoCb, c.AllocaCStr(node), c.AllocaCStr(service), hints))
}

// UvFreeaddrinfo frees the address information.
func UvFreeaddrinfo(addrInfo *net.AddrInfo) {
	Freeaddrinfo(addrInfo)
}

// ----------------------------------------------

/* Getnameinfo related function and method */

//go:linkname Getnameinfo C.uv_getnameinfo
func Getnameinfo(loop *Loop, req *GetNameInfo, getnameinfoCb GetnameinfoCb, addr *net.SockAddr, flags c.Int) c.Int

// UvGetnameinfo gets the name information.
func UvGetnameinfo(loop *Loop, req *GetNameInfo, getnameinfoCb GetnameinfoCb, addr *net.SockAddr, flags int) int {
	return int(Getnameinfo(loop, req, getnameinfoCb, addr, c.Int(flags)))
}
