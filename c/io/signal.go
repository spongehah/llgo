package io

import (
	_ "unsafe"
)

type Signal struct {
	Unused [0]byte
}

type SignalCb func(handle *Signal, sigNum int)

//go:linkname UvSignalInit C.uv_signal_init
func UvSignalInit(loop *Loop, handle *Signal) int

// llgo:link (*Signal).Start C.uv_signal_start
func (s *Signal) Start(cb SignalCb, signum int) int {
	return 0
}

// llgo:link (*Signal).StartOneshot C.uv_signal_start_oneshot
func (s *Signal) StartOneshot(handle *Signal, cb SignalCb, signum int) int {
	return 0
}

func (s *Signal) Init(loop *Loop) int {
	return UvSignalInit(loop, s)
}
