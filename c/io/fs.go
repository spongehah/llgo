package io

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs libuv); -luv"
)

const (
	DirentUnknown DirentType = iota
	DirentFile
	DirentDir
	DirentLink
	DirentFifo
	DirentSocket
	DirentChar
	DirentBlock
)

type DirentType c.Int

/* Handle types. */

type Fs struct {
	Unused [0]byte
}

type FsEvent struct {
	Unused [0]byte
}

type FsPoll struct {
	Unused [0]byte
}

type Dirent struct {
	Name string
	Type DirentType
}

type FsCb func(req *Fs)

type FsEventCb func(handle *FsEvent, filename string, events int, status int)

type FsPollCb func(handle *FsPoll, status int, events int)

/* Request types. */

/* None of the above. */

// ----------------------------------------------
//
//go:linkname FsReqCleanup C.uv_fs_req_cleanup
func FsReqCleanup(req *Fs)

//go:linkname FsOpen C.uv_fs_open
func FsOpen(loop *Loop, req *Fs, path string, flags int, mode int, cb FsCb) int

//go:linkname FsClose C.uv_fs_close
func FsClose(loop *Loop, req *Fs, file int, cb FsCb) int

//go:linkname FsRead C.uv_fs_read
func FsRead(loop *Loop, req *Fs, file int, buf []byte, offset int, cb FsCb) int

//go:linkname FsWrite C.uv_fs_write
func FsWrite(loop *Loop, req *Fs, file int, buf []byte, offset int, cb FsCb) int

//go:linkname FsUnlink C.uv_fs_unlink
func FsUnlink(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsMkdir C.uv_fs_mkdir
func FsMkdir(loop *Loop, req *Fs, path string, mode int, cb FsCb) int

//go:linkname FsMkdtemp C.uv_fs_mkdtemp
func FsMkdtemp(loop *Loop, req *Fs, tpl string, cb FsCb) int

//go:linkname FsMkStemp C.uv_fs_mkstemp
func FsMkStemp(loop *Loop, req *Fs, tpl string, cb FsCb) int

//go:linkname FsRmdir C.uv_fs_rmdir
func FsRmdir(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsStat C.uv_fs_stat
func FsStat(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsFstat C.uv_fs_fstat
func FsFstat(loop *Loop, req *Fs, file int, cb FsCb) int

//go:linkname FsRename C.uv_fs_rename
func FsRename(loop *Loop, req *Fs, path string, newPath string, cb FsCb) int

//go:linkname FsFsync C.uv_fs_fsync
func FsFsync(loop *Loop, req *Fs, file int, cb FsCb) int

//go:linkname FsFdatasync C.uv_fs_fdatasync
func FsFdatasync(loop *Loop, req *Fs, file int, cb FsCb) int

//go:linkname FsFtruncate C.uv_fs_ftruncate
func FsFtruncate(loop *Loop, req *Fs, file int, offset int, cb FsCb) int

//go:linkname FsSendfile C.uv_fs_sendfile
func FsSendfile(loop *Loop, req *Fs, outFd int, inFd int, inOffset int, length int, cb FsCb) int

//go:linkname FsAccess C.uv_fs_access
func FsAccess(loop *Loop, req *Fs, path string, flags int, cb FsCb) int

//go:linkname FsChmod C.uv_fs_chmod
func FsChmod(loop *Loop, req *Fs, path string, mode int, cb FsCb) int

//go:linkname FsFchmod C.uv_fs_fchmod
func FsFchmod(loop *Loop, req *Fs, file int, mode int, cb FsCb) int

//go:linkname FsUtime C.uv_fs_utime
func FsUtime(loop *Loop, req *Fs, path string, atime int, mtime int, cb FsCb) int

//go:linkname FsFutime C.uv_fs_futime
func FsFutime(loop *Loop, req *Fs, file int, atime int, mtime int, cb FsCb) int

//go:linkname FsLutime C.uv_fs_lutime
func FsLutime(loop *Loop, req *Fs, path string, atime int, mtime int, cb FsCb) int

//go:linkname FsLink C.uv_fs_link
func FsLink(loop *Loop, req *Fs, path string, newPath string, cb FsCb) int

//go:linkname FsSymlink C.uv_fs_symlink
func FsSymlink(loop *Loop, req *Fs, path string, newPath string, flags int, cb FsCb) int

//go:linkname FsReadlink C.uv_fs_read
func FsReadlink(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsRealpath C.uv_fs_realpath
func FsRealpath(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsCopyfile C.uv_fs_copyfile
func FsCopyfile(loop *Loop, req *Fs, path string, newPath string, flags int, cb FsCb) int

//go:linkname FsScandir C.uv_fs_scandir
func FsScandir(loop *Loop, req *Fs, path string, flags int, cb FsCb) int

//go:linkname FsScandirNext C.uv_fs_scandir_next
func FsScandirNext(req *Fs, ent *Dirent) int

//go:linkname FsOpenDir C.uv_fs_opendir
func FsOpenDir(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsReaddir C.uv_fs_readdir
func FsReaddir(loop *Loop, req *Fs, dir int, cb FsCb) int

//go:linkname FsCloseDir C.uv_fs_closedir
func FsCloseDir(loop *Loop, req *Fs) int

//go:linkname FsStatfs C.uv_fs_statfs
func FsStatfs(loop *Loop, req *Fs, path string, cb FsCb) int

//go:linkname FsEventInit C.uv_fs_event_init
func FsEventInit(loop *Loop, handle *FsEvent) int

//go:linkname FsEventStart C.uv_fs_event_start
func FsEventStart(handle *FsEvent, cb FsEventCb, path string, flags int) int

//go:linkname FsEventStop C.uv_fs_event_stop
func FsEventStop(handle *FsEvent) int

//go:linkname FsEventClose C.uv_fs_event_close
func FsEventClose(handle *FsEvent) int

//go:linkname FsEventGetpath C.uv_fs_event_getpath
func FsEventGetpath(handle *FsEvent) string

//go:linkname FsPollInit C.uv_fs_poll_init
func FsPollInit(loop *Loop, handle *FsPoll) int

//go:linkname FsPollStart C.uv_fs_poll_start
func FsPollStart(handle *FsPoll, cb FsPollCb, path string, interval uint) int

//go:linkname FsPollStop C.uv_fs_poll_stop
func FsPollStop(handle *FsPoll) int

//go:linkname FsPollClose C.uv_fs_poll_close
func FsPollClose(handle *FsPoll) int

//go:linkname FsPollGetPath C.uv_fs_poll_getpath
func FsPollGetPath(handle *FsPoll) string

//go:linkname FsChown C.uv_fs_chown
func FsChown(loop *Loop, req *Fs, path string, uid int, gid int, cb FsCb) int

//go:linkname FsFchown C.uv_fs_fchown
func FsFchown(loop *Loop, req *Fs, file int, uid int, gid int, cb FsCb) int

//go:linkname FsLchown C.uv_fs_lchown
func FsLchown(loop *Loop, req *Fs, path string, uid int, gid int, cb FsCb) int

//go:linkname FsLstat C.uv_fs_lstat
func FsLstat(loop *Loop, req *Fs, path string, cb FsCb) int
