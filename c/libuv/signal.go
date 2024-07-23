package libuv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/* Handle types. */

type Signal struct {
	Unused [0]byte
}

// ----------------------------------------------

/* Function type */

// llgo:type C
type SignalCb func(handle *Signal, sigNum c.Int)

// ----------------------------------------------

/* Signal related functions and method. */

//go:linkname UvSignalInit C.uv_signal_init
func UvSignalInit(loop *Loop, handle *Signal) c.Int

//go:linkname UvSignalStart C.uv_signal_start
func UvSignalStart(handle *Signal, cb SignalCb, signum c.Int) c.Int

//go:linkname UvSignalStartOneshot C.uv_signal_start_oneshot
func UvSignalStartOneshot(handle *Signal, cb SignalCb, signum c.Int) c.Int

// Start starts the signal handle with a callback.
func (s *Signal) Start(cb SignalCb, signum int) int {
	return int(UvSignalStart(s, cb, c.Int(signum)))
}

// StartOneshot starts the signal handle with a callback.
func (s *Signal) StartOneshot(handle *Signal, cb SignalCb, signum int) int {
	return int(UvSignalStartOneshot(s, cb, c.Int(signum)))
}

// Init initializes the signal handle.
func (s *Signal) Init(loop *Loop) int {
	return int(UvSignalInit(loop, s))
}
