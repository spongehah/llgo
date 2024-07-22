package libuv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

//go:linkname TranslateSysError C.uv_translate_sys_error
func TranslateSysError(sysErrno c.Int) c.Int

//go:linkname Strerror C.uv_strerror
func Strerror(err c.Int) *c.Char

//go:linkname StrerrorR C.uv_strerror_r
func StrerrorR(err c.Int, buf *c.Char, bufLen uintptr) *c.Char

//go:linkname ErrName C.uv_err_name
func ErrName(err c.Int) *c.Char

//go:linkname ErrNameR C.uv_err_name_r
func ErrNameR(err c.Int, buf *c.Char, bufLen uintptr) *c.Char
