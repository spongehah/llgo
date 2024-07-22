package io

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs libuv); -luv"
)

// ----------------------------------------------

const (
	LOOP_BLOCK_SIGNAL LoopOption = iota
	METRICS_IDLE_TIME
)

const (
	RUN_DEFAULT RunMode = iota
	RUN_ONCE
	RUN_NOWAIT
)

const (
	UV_LEAVE_GROUP Membership = iota
	UV_JOIN_GROUP
)

type LoopOption c.Int

type RunMode c.Int

type OsSock c.Int

type Membership c.Int

// ----------------------------------------------

/* Handle types. */
type Loop struct {
	Unused [0]byte
}

type Handle struct {
}

type Dir struct {
}

type Stream struct {
}

type Tcp struct {
}

type Udp struct {
}

type Pipe struct {
}

type Tty struct {
}

type Poll struct {
}

type Timer struct {
}

type Prepare struct {
}

type Check struct {
}

type Idle struct {
}

type Async struct {
}

type Process struct {
}

type FsEvent struct {
}

type FsPoll struct {
}

type Signal struct {
}

/* Request types. */

type GetAddrInfo struct {
}

type GetNameInfo struct {
}

type Connect struct {
}

type UdpSend struct {
}

/* None of the above. */

// ----------------------------------------------

type Buf struct {
	base *c.Char
	len  uintptr
}

// ----------------------------------------------

// llgo:type C
type MallocFunc func(size uintptr) c.Pointer

// llgo:type C
type ReallocFunc func(ptr c.Pointer, size uintptr) c.Pointer

// llgo:type C
type CallocFunc func(count uintptr, size uintptr) c.Pointer

// llgo:type C
type FreeFunc func(ptr c.Pointer)

// ----------------------------------------------

// llgo:type C
type AllocCb func(handle *Handle, suggestedSize uintptr, buf *Buf)

// llgo:type C
type GetaddrinfoCb func(req *GetAddrInfo, status c.Int, res *net.AddrInfo)

// llgo:type C
type GetnameinfoCb func(req *GetNameInfo, status c.Int, hostname *c.Char, service *c.Char)

// ----------------------------------------------

//go:linkname Version C.uv_version
func Version() c.Uint

//go:linkname VersionString C.uv_version_string
func VersionString() *c.Char

//go:linkname LibraryShutdown C.uv_library_shutdown
func LibraryShutdown()

//go:linkname ReplaceAllocator C.uv_replace_allocator
func ReplaceAllocator(mallocFunc MallocFunc, reallocFunc ReallocFunc, callocFunc CallocFunc, freeFunc FreeFunc) c.Int

// ----------------------------------------------

/* LoopT related function and method */

//go:linkname DefaultLoop C.uv_default_loop
func DefaultLoop() *Loop

//go:linkname LoopSize C.uv_loop_size
func LoopSize() uintptr

// llgo:link (*Loop).Init C.uv_loop_init
func (loop *Loop) Init() c.Int {
	return 0
}

// llgo:link (*Loop).Close C.uv_loop_close
func (loop *Loop) Close() c.Int {
	return 0
}

// llgo:link (*Loop).Alive C.uv_loop_alive
func (loop *Loop) Alive() c.Int {
	return 0
}

// llgo:link (*Loop).Configure C.uv_loop_configure
func (loop *Loop) Configure(option LoopOption, __llgo_va_list ...any) c.Int {
	return 0
}

// llgo:link (*Loop).Fork C.uv_loop_fork
func (loop *Loop) Fork() c.Int {
	return 0
}

// llgo:link (*Loop).Run C.uv_run
func (loop *Loop) Run(mode RunMode) c.Int {
	return 0
}

// llgo:link (*Loop).Stop C.uv_stop
func (loop *Loop) Stop() {}

// llgo:link (*Loop).UpdateTime C.uv_update_time
func (loop *Loop) UpdateTime() {}

// llgo:link (*Loop).Now C.uv_now
func (loop *Loop) Now() uint64 {
	return 0
}

// llgo:link (*Loop).BackendFd C.uv_backend_fd
func (loop *Loop) BackendFd() c.Int {
	return 0
}

// llgo:link (*Loop).BackendTimeout C.uv_backend_timeout
func (loop *Loop) BackendTimeout() c.Int {
	return 0
}

// ----------------------------------------------

/* HandleT related function and method */

// llgo:link (*Handle).Ref C.uv_ref
func (handle *Handle) Ref() {}

// llgo:link (*Handle).Unref C.uv_unref
func (handle *Handle) Unref() {}

// llgo:link (*Handle).HasRef C.uv_has_ref
func (handle *Handle) HasRef() c.Int {
	return 0
}

// ----------------------------------------------

/* Getaddrinfo related function and method */

//go:linkname Getaddrinfo C.uv_getaddrinfo
func Getaddrinfo(loop *Loop, req *GetAddrInfo, getaddrinfoCb GetaddrinfoCb, node *c.Char, service *c.Char, hints *net.AddrInfo) c.Int

//go:linkname Freeaddrinfo C.uv_freeaddrinfo
func Freeaddrinfo(addrInfo *net.AddrInfo)

// ----------------------------------------------

/* Getnameinfo related function and method */

//go:linkname Getnameinfo C.uv_getnameinfo
func Getnameinfo(loop *Loop, req *GetNameInfo, getnameinfoCb GetnameinfoCb, addr *net.SockAddr, flags c.Int) c.Int
