package io

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
	_ "unsafe"
)

/*
 * UDP support.
 */
const (
	/* Disables dual stack mode. */
	UDP_IPV6ONLY UdpFlags = 1
	/*
	 * Indicates message was truncated because read buffer was too small. The
	 * remainder was discarded by the OS. Used in uv_udp_recv_cb.
	 */
	UDP_PARTIAL UdpFlags = 2
	/*
	 * Indicates if SO_REUSEADDR will be set when binding the handle.
	 * This sets the SO_REUSEPORT socket flag on the BSDs and OS X. On other
	 * Unix platforms, it sets the SO_REUSEADDR flag.  What that means is that
	 * multiple threads or processes can bind to the same address without error
	 * (provided they all set the flag) but only the last one to bind will receive
	 * any traffic, in effect "stealing" the port from the previous listener.
	 */
	UDP_REUSEADDR UdpFlags = 4
	/*
	 * Indicates that the message was received by recvmmsg, so the buffer provided
	 * must not be freed by the recv_cb callback.
	 */
	UDP_MMSG_CHUNK UdpFlags = 8
	/*
	 * Indicates that the buffer provided has been fully utilized by recvmmsg and
	 * that it should now be freed by the recv_cb callback. When this flag is set
	 * in uv_udp_recv_cb, nread will always be 0 and addr will always be NULL.
	 */
	UDP_MMSG_FREE UdpFlags = 16
	/*
	 * Indicates if IP_RECVERR/IPV6_RECVERR will be set when binding the handle.
	 * This sets IP_RECVERR for IPv4 and IPV6_RECVERR for IPv6 UDP sockets on
	 * Linux. This stops the Linux kernel from suppressing some ICMP error
	 * messages and enables full ICMP error reporting for faster failover.
	 * This flag is no-op on platforms other than Linux.
	 */
	UDP_LINUX_RECVERR UdpFlags = 32
	/*
	 * Indicates that recvmmsg should be used, if available.
	 */
	UDP_RECVMMSG UdpFlags = 256
)

type UdpFlags c.Int

// llgo:type C
type UdpSendCb func(req *UdpSend, status c.Int)

// llgo:type C
type UdpRecvCb func(handle *Udp, nread c.Long, buf *Buf, addr *net.SockAddr, flags c.Uint)

/* UdpT related function and method */

//go:linkname InitUdp C.uv_udp_init
func InitUdp(loop *Loop, udp *Udp) c.Int

//go:linkname InitUdpEx C.uv_udp_init_ex
func InitUdpEx(loop *Loop, udp *Udp, flags c.Uint) c.Int

// llgo:link (*Udp).Open C.uv_udp_open
func (udp *Udp) Open(sock OsSock) c.Int {
	return 0
}

// llgo:link (*Udp).Bind C.uv_udp_bind
func (udp *Udp) Bind(addr *net.SockAddr, flags c.Uint) c.Int {
	return 0
}

// llgo:link (*Udp).Connect C.uv_udp_connect
func (udp *Udp) Connect(addr *net.SockAddr) c.Int {
	return 0
}

// llgo:link (*Udp).Getpeername C.uv_udp_getpeername
func (udp *Udp) Getpeername(name *net.SockAddr, nameLen *c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).Getsockname C.uv_udp_getsockname
func (udp *Udp) Getsockname(name *net.SockAddr, nameLen *c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).SetMembership C.uv_udp_set_membership
func (udp *Udp) SetMembership(multicastAddr *c.Char, interfaceAddr *c.Char, membership Membership) c.Int {
	return 0
}

// llgo:link (*Udp).SourceMembership C.uv_udp_set_source_membership
func (udp *Udp) SourceMembership(multicastAddr *c.Char, interfaceAddr *c.Char, sourceAddr *c.Char, membership Membership) c.Int {
	return 0
}

// llgo:link (*Udp).SetMulticastLoop C.uv_udp_set_multicast_loop
func (udp *Udp) SetMulticastLoop(on c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).SetMulticastTTL C.uv_udp_set_multicast_ttl
func (udp *Udp) SetMulticastTTL(ttl c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).SetMulticastInterface C.uv_udp_set_multicast_interface
func (udp *Udp) SetMulticastInterface(interfaceAddr *c.Char) c.Int {
	return 0
}

// llgo:link (*Udp).SAetBroadcast C.uv_udp_set_broadcast
func (udp *Udp) SAetBroadcast(on c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).SetTTL C.uv_udp_set_ttl
func (udp *Udp) SetTTL(ttl c.Int) c.Int {
	return 0
}

//go:linkname Send C.uv_udp_send
func Send(req *UdpSend, udp *Udp, bufs []Buf, nbufs c.Uint, addr *net.SockAddr, sendCb UdpSendCb) c.Int

// llgo:link (*Udp).TrySend C.uv_udp_try_send
func (udp *Udp) TrySend(bufs []Buf, nbufs c.Uint, addr *net.SockAddr) c.Int {
	return 0
}

// llgo:link (*Udp).RecvStart C.uv_udp_recv_start
func (udp *Udp) RecvStart(allocCb AllocCb, recvCb UdpRecvCb) c.Int {
	return 0
}

// llgo:link (*Udp).UsingRecvmmsg C.uv_udp_using_recvmmsg
func (udp *Udp) UsingRecvmmsg() c.Int {
	return 0
}

// llgo:link (*Udp).RecvStop C.uv_udp_recv_stop
func (udp *Udp) RecvStop() c.Int {
	return 0
}

// llgo:link (*Udp).GetSendQueueSize C.uv_udp_get_send_queue_size
func (udp *Udp) GetSendQueueSize() uintptr {
	return 0
}

// llgo:link (*Udp).GetSendQueueCount C.uv_udp_get_send_queue_count
func (udp *Udp) GetSendQueueCount() uintptr {
	return 0
}
