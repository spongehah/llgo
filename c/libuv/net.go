package libuv

import (
	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/net"
	_ "unsafe"
)

const (
	/* Used with uv_tcp_bind, when an IPv6 address is used. */
	TCP_IPV6ONLY TcpFlags = 1
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

type TcpFlags c.Int

type UdpFlags c.Int

// ----------------------------------------------

/* Handle types. */

type Tcp struct {
	Unused [264]byte
}

type Udp struct {
	Unused [0]byte
}

/* Request types. */

type UdpSend struct {
	Unused [0]byte
}

// ----------------------------------------------

/* Function type */

// llgo:type C
type CloseCb func(handle *Handle)

// llgo:type C
type ConnectCb func(req *Connect, status c.Int)

// llgo:type C
type UdpSendCb func(req *UdpSend, status c.Int)

// llgo:type C
type UdpRecvCb func(handle *Udp, nread c.Long, buf *Buf, addr *net.SockAddr, flags c.Uint)

// ----------------------------------------------

/* Tcp related function and method */

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

// UvInitTcp initializes the tcp handle.
func UvInitTcp(loop *Loop, tcp *Tcp) int {
	return int(InitTcp(loop, tcp))
}

// UvInitTcpEx initializes the tcp handle with flags.
func UvInitTcpEx(loop *Loop, tcp *Tcp, flags c.Uint) int {
	return int(InitTcpEx(loop, tcp, flags))
}

// UvOpen opens the tcp handle.
func (tcp *Tcp) UvOpen(sock OsSock) int {
	return int(tcp.Open(sock))
}

// UvNodelay sets the nodelay option.
func (tcp *Tcp) UvNodelay(enable int) int {
	return int(tcp.Nodelay(c.Int(enable)))
}

// UvKeepAlive sets the keepalive option.
func (tcp *Tcp) UvKeepAlive(enable int, delay uint) int {
	return int(tcp.KeepAlive(c.Int(enable), c.Uint(delay)))
}

// UvSimultaneousAccepts sets the simultaneous accepts option.
func (tcp *Tcp) UvSimultaneousAccepts(enable int) int {
	return int(tcp.SimultaneousAccepts(c.Int(enable)))
}

// UvBind binds the tcp handle to the address.
func (tcp *Tcp) UvBind(addr *net.SockAddr, flags uint) int {
	return int(tcp.Bind(addr, c.Uint(flags)))
}

// UvGetsockname gets the socket name.
func (tcp *Tcp) UvGetsockname(name *net.SockAddr, nameLen *int) int {
	return int(tcp.Getsockname(name, (*c.Int)(c.Pointer(nameLen))))
}

// UvGetpeername gets the peer name.
func (tcp *Tcp) UvGetpeername(name *net.SockAddr, nameLen *int) int {
	return int(tcp.Getpeername(name, (*c.Int)(c.Pointer(nameLen))))
}

// UvCloseReset closes the tcp handle with reset.
func (tcp *Tcp) UvCloseReset(closeCb CloseCb) int {
	return int(tcp.CloseReset(closeCb))
}

// UvTcpConnect connects the tcp handle to the address.
func UvTcpConnect(req *Connect, tcp *Tcp, addr *net.SockAddr, connectCb ConnectCb) int {
	return int(TcpConnect(req, tcp, addr, connectCb))
}

// ----------------------------------------------

/* Udp related function and method */

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

// llgo:link (*Udp).SetBroadcast C.uv_udp_set_broadcast
func (udp *Udp) SetBroadcast(on c.Int) c.Int {
	return 0
}

// llgo:link (*Udp).SetTTL C.uv_udp_set_ttl
func (udp *Udp) SetTTL(ttl c.Int) c.Int {
	return 0
}

//go:linkname Send C.uv_udp_send
func Send(req *UdpSend, udp *Udp, bufs *Buf, nbufs c.Uint, addr *net.SockAddr, sendCb UdpSendCb) c.Int

// llgo:link (*Udp).TrySend C.uv_udp_try_send
func (udp *Udp) TrySend(bufs *Buf, nbufs c.Uint, addr *net.SockAddr) c.Int {
	return 0
}

// llgo:link (*Udp).StartRecv C.uv_udp_recv_start
func (udp *Udp) StartRecv(allocCb AllocCb, recvCb UdpRecvCb) c.Int {
	return 0
}

// llgo:link (*Udp).UsingRecvmmsg C.uv_udp_using_recvmmsg
func (udp *Udp) UsingRecvmmsg() c.Int {
	return 0
}

// llgo:link (*Udp).StopRecv C.uv_udp_recv_stop
func (udp *Udp) StopRecv() c.Int {
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

// UvInitUdp initializes the udp handle.
func UvInitUdp(loop *Loop, udp *Udp) int {
	return int(InitUdp(loop, udp))
}

// UvInitUdpEx initializes the udp handle with flags.
func UvInitUdpEx(loop *Loop, udp *Udp, flags uint) int {
	return int(InitUdpEx(loop, udp, c.Uint(flags)))
}

// UvOpen opens the udp handle.
func (udp *Udp) UvOpen(sock OsSock) int {
	return int(udp.Open(sock))
}

// UvBind binds the udp handle to the address.
func (udp *Udp) UvBind(addr *net.SockAddr, flags uint) int {
	return int(udp.Bind(addr, c.Uint(flags)))
}

// UvConnect connects the udp handle to the address.
func (udp *Udp) UvConnect(addr *net.SockAddr) int {
	return int(udp.Connect(addr))
}

// UvGetpeername gets the peer name.
func (udp *Udp) UvGetpeername(name *net.SockAddr, nameLen *int) int {
	return int(udp.Getpeername(name, (*c.Int)(c.Pointer(nameLen))))
}

// UvGetsockname gets the socket name.
func (udp *Udp) UvGetsockname(name *net.SockAddr, nameLen *int) int {
	return int(udp.Getsockname(name, (*c.Int)(c.Pointer(nameLen))))
}

// UvSetMembership sets the membership.
func (udp *Udp) UvSetMembership(multicastAddr, interfaceAddr string, membership Membership) int {
	return int(udp.SetMembership(c.AllocaCStr(multicastAddr), c.AllocaCStr(interfaceAddr), membership))
}

// UvSourceMembership sets the source membership.
func (udp *Udp) UvSourceMembership(multicastAddr, interfaceAddr, sourceAddr string, membership Membership) int {
	return int(udp.SourceMembership(c.AllocaCStr(multicastAddr), c.AllocaCStr(interfaceAddr), c.AllocaCStr(sourceAddr), membership))
}

// UvSetMulticastLoop sets the multicast loop.
func (udp *Udp) UvSetMulticastLoop(on int) int {
	return int(udp.SetMulticastLoop(c.Int(on)))
}

// UvSetMulticastTTL sets the multicast ttl.
func (udp *Udp) UvSetMulticastTTL(ttl int) int {
	return int(udp.SetMulticastTTL(c.Int(ttl)))
}

// UvSetMulticastInterface sets the multicast interface.
func (udp *Udp) UvSetMulticastInterface(interfaceAddr string) int {
	return int(udp.SetMulticastInterface(c.AllocaCStr(interfaceAddr)))
}

// UvSetBroadcast sets the broadcast.
func (udp *Udp) UvSetBroadcast(on int) int {
	return int(udp.SetBroadcast(c.Int(on)))
}

// UvSetTTL sets the ttl.
func (udp *Udp) UvSetTTL(ttl int) int {
	return int(udp.SetTTL(c.Int(ttl)))
}

// UvSend sends the data.
func UvSend(req *UdpSend, udp *Udp, bufs *Buf, nbufs uint, addr *net.SockAddr, sendCb UdpSendCb) int {
	return int(Send(req, udp, bufs, c.Uint(nbufs), addr, sendCb))
}

// UvTrySend tries to send the data.
func (udp *Udp) UvTrySend(bufs *Buf, nbufs uint, addr *net.SockAddr) int {
	return int(udp.TrySend(bufs, c.Uint(nbufs), addr))
}

// UvStartRecv starts to receive the data.
func (udp *Udp) UvStartRecv(allocCb AllocCb, recvCb UdpRecvCb) int {
	return int(udp.StartRecv(allocCb, recvCb))
}

// UvUsingRecvmmsg returns if using recvmmsg.
func (udp *Udp) UvUsingRecvmmsg() int {
	return int(udp.UsingRecvmmsg())
}

// UvStopRecv stops to receive the data.
func (udp *Udp) UvStopRecv() int {
	return int(udp.StopRecv())
}

// UvGetSendQueueSize gets the send queue size.
func (udp *Udp) UvGetSendQueueSize() uintptr {
	return udp.GetSendQueueSize()
}

// UvGetSendQueueCount gets the send queue count.
func (udp *Udp) UvGetSendQueueCount() uintptr {
	return udp.GetSendQueueCount()
}

// ----------------------------------------------

//go:linkname Ip4Addr C.uv_ip4_addr
func Ip4Addr(ip *c.Char, port c.Int, addr *net.SockaddrIn) c.Int

//go:linkname Ip6Addr C.uv_ip6_addr
func Ip6Addr(ip *c.Char, port c.Int, addr *net.SockaddrIn6) c.Int

//go:linkname Ip4Name C.uv_ip4_name
func Ip4Name(src *net.SockaddrIn, dst *c.Char, size uintptr) c.Int

//go:linkname Ip6Name C.uv_ip6_name
func Ip6Name(src *net.SockaddrIn6, dst *c.Char, size uintptr) c.Int

//go:linkname IpName C.uv_ip_name
func IpName(src *net.SockAddr, dst *c.Char, size uintptr) c.Int

//go:linkname InetNtop C.uv_inet_ntop
func InetNtop(af c.Int, src c.Pointer, dst *c.Char, size uintptr) c.Int

//go:linkname InetPton C.uv_inet_pton
func InetPton(af c.Int, src *c.Char, dst c.Pointer) c.Int

// UvIp4Addr fills in the sockaddr_in struct with the address.
func UvIp4Addr(ip string, port int, addr *net.SockaddrIn) int {
	return int(Ip4Addr(c.AllocaCStr(ip), c.Int(port), addr))
}

// UvIp6Addr fills in the sockaddr_in6 struct with the address.
func UvIp6Addr(ip string, port int, addr *net.SockaddrIn6) int {
	return int(Ip6Addr(c.AllocaCStr(ip), c.Int(port), addr))
}

// UvIp4Name fills in the ip address string.
func UvIp4Name(src *net.SockaddrIn, dst string) int {
	return int(Ip4Name(src, c.AllocaCStr(dst), uintptr(len(dst))))
}

// UvIp6Name fills in the ip address string.
func UvIp6Name(src *net.SockaddrIn6, dst string) int {
	return int(Ip6Name(src, c.AllocaCStr(dst), uintptr(len(dst))))
}

// UvIpName fills in the ip address string.
func UvIpName(src *net.SockAddr, dst string) int {
	return int(IpName(src, c.AllocaCStr(dst), uintptr(len(dst))))
}

// UvInetNtop converts the network address to the presentation format.
func UvInetNtop(af int, src c.Pointer, dst string) int {
	return int(InetNtop(c.Int(af), src, c.AllocaCStr(dst), uintptr(len(dst))))
}

// UvInetPton converts the presentation format to the network address.
func UvInetPton(af int, src string, dst c.Pointer) int {
	return int(InetPton(c.Int(af), c.AllocaCStr(src), dst))
}
