package io

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
	_ "unsafe"
)

const (
	/* Used with uv_tcp_bind, when an IPv6 address is used. */
	TCP_IPV6ONLY TcpFlags = 1
)

type TcpFlags c.Int

// llgo:type C
type CloseCb func(handle *Handle)

// llgo:type C
type ConnectCb func(req *Connect, status c.Int)

/* TcpT related function and method */

//go:linkname InitTcp C.uv_tcp_init
func InitTcp(loop *Loop, tcp *Tcp) c.Int

//go:linkname InitTcpEx C.uv_tcp_init_ex
func InitTcpEx(loop *Loop, tcp *Tcp, flags c.Uint) c.Int

// llgo:link (*Tcp).Open C.uv_tcp_open
func (tcp *Tcp) Open(sock OsSock) c.Int {
	return 0
}

// llgo:link (*Tcp).Nodelay C.uv_tcp_nodelay
func (tcp *Tcp) Nodelay(enable c.Int) c.Int {
	return 0
}

// llgo:link (*Tcp).KeepAlive C.uv_tcp_keepalive
func (tcp *Tcp) KeepAlive(enable c.Int, delay c.Uint) c.Int {
	return 0
}

// llgo:link (*Tcp).SimultaneousAccepts C.uv_tcp_simultaneous_accepts
func (tcp *Tcp) SimultaneousAccepts(enable c.Int) c.Int {
	return 0
}

// llgo:link (*Tcp).Bind C.uv_tcp_bind
func (tcp *Tcp) Bind(addr *net.SockAddr, flags c.Uint) c.Int {
	return 0
}

// llgo:link (*Tcp).Getsockname C.uv_tcp_getsockname
func (tcp *Tcp) Getsockname(name *net.SockAddr, nameLen *c.Int) c.Int {
	return 0
}

// llgo:link (*Tcp).Getpeername C.uv_tcp_getpeername
func (tcp *Tcp) Getpeername(name *net.SockAddr, nameLen *c.Int) c.Int {
	return 0
}

// llgo:link (*Tcp).CloseReset C.uv_tcp_close_reset
func (tcp *Tcp) CloseReset(closeCb CloseCb) c.Int {
	return 0
}

//go:linkname TcpConnect C.uv_tcp_connect
func TcpConnect(req *Connect, tcp *Tcp, addr *net.SockAddr, connectCb ConnectCb) c.Int
