package libuv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	E2BIG Errno = iota
	EACCES
	EADDRINUSE
	EADDRNOTAVAIL
	EAFNOSUPPORT
	EAGAIN
	EAI_ADDRFAMILY
	EAI_AGAIN
	EAI_BADFLAGS
	EAI_BADHINTS
	EAI_CANCELED
	EAI_FAIL
	EAI_FAMILY
	EAI_MEMORY
	EAI_NODATA
	EAI_NONAME
	EAI_OVERFLOW
	EAI_PROTOCOL
	EAI_SERVICE
	EAI_SOCKTYPE
	EALREADY
	EBADF
	EBUSY
	ECANCELED
	ECHARSET
	ECONNABORTED
	ECONNREFUSED
	ECONNRESET
	EDESTADDRREQ
	EEXIST
	EFAULT
	EFBIG
	EHOSTUNREACH
	EINTR
	EINVAL
	EIO
	EISCONN
	EISDIR
	ELOOP
	EMFILE
	EMSGSIZE
	ENAMETOOLONG
	ENETDOWN
	ENETUNREACH
	ENFILE
	ENOBUFS
	ENODEV
	ENOENT
	ENOMEM
	ENONET
	ENOPROTOOPT
	ENOSPC
	ENOSYS
	ENOTCONN
	ENOTDIR
	ENOTEMPTY
	ENOTSOCK
	ENOTSUP
	EOVERFLOW
	EPERM
	EPIPE
	EPROTO
	EPROTONOSUPPORT
	EPROTOTYPE
	ERANGE
	EROFS
	ESHUTDOWN
	ESPIPE
	ESRCH
	ETIMEDOUT
	ETXTBSY
	EXDEV
	UNKNOWN
	EOF
	ENXIO
	EMLINK
	EHOSTDOWN
	EREMOTEIO
	ENOTTY
	EFTYPE
	EILSEQ
	ESOCKTNOSUPPORT
	ENODATA
	EUNATCH
	ERRNO_MAX
)

type Errno c.Int

//go:linkname UvTranslateSysError C.uv_translate_sys_error
func UvTranslateSysError(sysErrno c.Int) c.Int

//go:linkname UvStrerror C.uv_strerror
func UvStrerror(err c.Int) *c.Char

//go:linkname UvStrerrorR C.uv_strerror_r
func UvStrerrorR(err c.Int, buf *c.Char, bufLen uintptr) *c.Char

//go:linkname UvErrName C.uv_err_name
func UvErrName(err c.Int) *c.Char

//go:linkname UvErrNameR C.uv_err_name_r
func UvErrNameR(err c.Int, buf *c.Char, bufLen uintptr) *c.Char

func (e Errno) Error() string {
	return c.GoString(UvStrerror(c.Int(e)))
}

func (e Errno) Name() string {
	return c.GoString(UvErrName(c.Int(e)))
}

func (e Errno) TranslateSysError() int {
	return int(UvTranslateSysError(c.Int(e)))
}

func Strerror(err c.Int) string {
	return c.GoString(UvStrerror(c.Int(err)))
}
