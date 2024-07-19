package uv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs libuv); -luv"
)

type LoopOption c.Int

const (
	LOOP_BLOCK_SIGNAL LoopOption = iota
	METRICS_IDLE_TIME
)

type RunMode c.Int

const (
	RUN_DEFAULT RunMode = iota
	RUN_ONCE
	RUN_NOWAIT
)

/* Handle types. */
type LoopT struct {
	Unused [0]byte
}

type HandleT struct {
}

type DirT struct {
}

type StreamT struct {
}

type TcpT struct {
}

type UdpT struct {
}

type PipeT struct {
}

type TtyT struct {
}

type PollT struct {
}

type TimerT struct {
}

type PrepareT struct {
}

type CheckT struct {
}

type IdleT struct {
}

type AsyncT struct {
}

type ProcessT struct {
}

type FsEventT struct {
}

type FsPollT struct {
}

type SignalT struct {
}

/* Request types. */

/* None of the above. */

// ----------------------------------------------

//go:linkname Version C.uv_version
func Version() c.Uint

//go:linkname VersionString C.uv_version_string
func VersionString() *c.Char

//go:linkname LibraryShutdown C.uv_library_shutdown
func LibraryShutdown()

//go:linkname ReplaceAllocator C.uv_replace_allocator
func ReplaceAllocator() c.Int

// ----------------------------------------------

/* LoopT related function and method */

//go:linkname DefaultLoop C.uv_default_loop
func DefaultLoop() *LoopT

//go:linkname LoopSize C.uv_loop_size
func LoopSize() uintptr

// llgo:link (*LoopT).Init C.uv_loop_init
func (loop *LoopT) Init() c.Int {
	return 0
}

// llgo:link (*LoopT).Close C.uv_loop_close
func (loop *LoopT) Close() c.Int {
	return 0
}

// llgo:link (*LoopT).Alive C.uv_loop_alive
func (loop *LoopT) Alive() c.Int {
	return 0
}

// llgo:link (*LoopT).Configure C.uv_loop_configure
func (loop *LoopT) Configure(option LoopOption, __llgo_va_list ...any) c.Int {
	return 0
}

// llgo:link (*LoopT).Fork C.uv_loop_fork
func (loop *LoopT) Fork() c.Int {
	return 0
}

// llgo:link (*LoopT).Run C.uv_run
func (loop *LoopT) Run(mode RunMode) c.Int {
	return 0
}

// llgo:link (*LoopT).Stop C.uv_stop
func (loop *LoopT) Stop() {}

// llgo:link (*LoopT).UpdateTime C.uv_update_time
func (loop *LoopT) UpdateTime() {}

// llgo:link (*LoopT).Now C.uv_now
func (loop *LoopT) Now() uint64 {
	return 0
}

// llgo:link (*LoopT).BackendFd C.uv_backend_fd
func (loop *LoopT) BackendFd() c.Int {
	return 0
}

// llgo:link (*LoopT).BackendTimeout C.uv_backend_timeout
func (loop *LoopT) BackendTimeout() c.Int {
	return 0
}

// llgo:link (*LoopT).Now C.uv_now

// llgo:link (*LoopT).Now C.uv_now

// llgo:link (*LoopT).Now C.uv_now

// ----------------------------------------------

/* HandleT related function and method */

// llgo:link (*HandleT).Ref C.uv_ref
func (handle *HandleT) Ref() {}

// llgo:link (*HandleT).Unref C.uv_unref
func (handle *HandleT) Unref() {}

// llgo:link (*HandleT).HasRef C.uv_has_ref
func (handle *HandleT) HasRef() c.Int {
	return 0
}

// llgo:link (*HandleT).Fork C.uv_loop_fork

// llgo:link (*HandleT).Fork C.uv_loop_fork
// llgo:link (*HandleT).Fork C.uv_loop_fork
// llgo:link (*HandleT).Fork C.uv_loop_fork
// llgo:link (*HandleT).Fork C.uv_loop_fork
