package libuv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
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

type File struct {
	loop *Loop
	req  *Fs
}

// llgo:type C
type FsCb func(req *Fs)

// llgo:type C
type FsEventCb func(handle *FsEvent, filename string, events c.Int, status c.Int)

// llgo:type C
type FsPollCb func(handle *FsPoll, status c.Int, events c.Int)

/* Request types. */

/* None of the above. */

// ----------------------------------------------

//go:linkname FsReqCleanup C.uv_fs_req_cleanup
func FsReqCleanup(req *Fs)

//go:linkname DefaultLoop C.uv_default_loop
func DefaultLoop() *Loop

//go:linkname FsOpen C.uv_fs_open
func FsOpen(loop *Loop, req *Fs, path string, flags c.Int, mode c.Int, cb FsCb) c.Int

//go:linkname FsClose C.uv_fs_close
func FsClose(loop *Loop, req *Fs, file c.Int, cb FsCb) c.Int

//go:linkname FsRead C.uv_fs_read
func FsRead(loop *Loop, req *Fs, file c.Int, buf []byte, offset c.Int, cb FsCb) c.Int

//go:linkname FsWrite C.uv_fs_write
func FsWrite(loop *Loop, req *Fs, file c.Int, buf []byte, offset c.Int, cb FsCb) c.Int

//go:linkname FsUnlink C.uv_fs_unlink
func FsUnlink(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsMkdir C.uv_fs_mkdir
func FsMkdir(loop *Loop, req *Fs, path string, mode c.Int, cb FsCb) c.Int

//go:linkname FsMkdtemp C.uv_fs_mkdtemp
func FsMkdtemp(loop *Loop, req *Fs, tpl string, cb FsCb) c.Int

//go:linkname FsMkStemp C.uv_fs_mkstemp
func FsMkStemp(loop *Loop, req *Fs, tpl string, cb FsCb) c.Int

//go:linkname FsRmdir C.uv_fs_rmdir
func FsRmdir(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsStat C.uv_fs_stat
func FsStat(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsFstat C.uv_fs_fstat
func FsFstat(loop *Loop, req *Fs, file c.Int, cb FsCb) c.Int

//go:linkname FsRename C.uv_fs_rename
func FsRename(loop *Loop, req *Fs, path string, newPath string, cb FsCb) c.Int

//go:linkname FsFsync C.uv_fs_fsync
func FsFsync(loop *Loop, req *Fs, file c.Int, cb FsCb) c.Int

//go:linkname FsFdatasync C.uv_fs_fdatasync
func FsFdatasync(loop *Loop, req *Fs, file c.Int, cb FsCb) c.Int

//go:linkname FsFtruncate C.uv_fs_ftruncate
func FsFtruncate(loop *Loop, req *Fs, file c.Int, offset c.Int, cb FsCb) c.Int

//go:linkname FsSendfile C.uv_fs_sendfile
func FsSendfile(loop *Loop, req *Fs, outFd c.Int, inFd c.Int, inOffset c.Int, length c.Int, cb FsCb) c.Int

//go:linkname FsAccess C.uv_fs_access
func FsAccess(loop *Loop, req *Fs, path string, flags c.Int, cb FsCb) c.Int

//go:linkname FsChmod C.uv_fs_chmod
func FsChmod(loop *Loop, req *Fs, path string, mode c.Int, cb FsCb) c.Int

//go:linkname FsFchmod C.uv_fs_fchmod
func FsFchmod(loop *Loop, req *Fs, file c.Int, mode c.Int, cb FsCb) c.Int

//go:linkname FsUtime C.uv_fs_utime
func FsUtime(loop *Loop, req *Fs, path string, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsFutime C.uv_fs_futime
func FsFutime(loop *Loop, req *Fs, file c.Int, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsLutime C.uv_fs_lutime
func FsLutime(loop *Loop, req *Fs, path string, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsLink C.uv_fs_link
func FsLink(loop *Loop, req *Fs, path string, newPath string, cb FsCb) c.Int

//go:linkname FsSymlink C.uv_fs_symlink
func FsSymlink(loop *Loop, req *Fs, path string, newPath string, flags c.Int, cb FsCb) c.Int

//go:linkname FsReadlink C.uv_fs_read
func FsReadlink(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsRealpath C.uv_fs_realpath
func FsRealpath(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsCopyfile C.uv_fs_copyfile
func FsCopyfile(loop *Loop, req *Fs, path string, newPath string, flags c.Int, cb FsCb) c.Int

//go:linkname FsScandir C.uv_fs_scandir
func FsScandir(loop *Loop, req *Fs, path string, flags c.Int, cb FsCb) c.Int

//go:linkname FsScandirNext C.uv_fs_scandir_next
func FsScandirNext(req *Fs, ent *Dirent) c.Int

//go:linkname FsOpenDir C.uv_fs_opendir
func FsOpenDir(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsReaddir C.uv_fs_readdir
func FsReaddir(loop *Loop, req *Fs, dir c.Int, cb FsCb) c.Int

//go:linkname FsCloseDir C.uv_fs_closedir
func FsCloseDir(loop *Loop, req *Fs) c.Int

//go:linkname FsStatfs C.uv_fs_statfs
func FsStatfs(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsChown C.uv_fs_chown
func FsChown(loop *Loop, req *Fs, path string, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsFchown C.uv_fs_fchown
func FsFchown(loop *Loop, req *Fs, file c.Int, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsLchown C.uv_fs_lchown
func FsLchown(loop *Loop, req *Fs, path string, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsLstat C.uv_fs_lstat
func FsLstat(loop *Loop, req *Fs, path string, cb FsCb) c.Int

//go:linkname FsEventInit C.uv_fs_event_init
func FsEventInit(loop *Loop, handle *FsEvent) c.Int

//go:linkname FsEventStart C.uv_fs_event_start
func FsEventStart(handle *FsEvent, cb FsEventCb, path string, flags c.Int) c.Int

//go:linkname FsEventStop C.uv_fs_event_stop
func FsEventStop(handle *FsEvent) c.Int

//go:linkname FsEventClose C.uv_fs_event_close
func FsEventClose(handle *FsEvent) c.Int

//go:linkname FsEventGetpath C.uv_fs_event_getpath
func FsEventGetpath(handle *FsEvent) string

//go:linkname FsPollInit C.uv_fs_poll_init
func FsPollInit(loop *Loop, handle *FsPoll) c.Int

//go:linkname FsPollStart C.uv_fs_poll_start
func FsPollStart(handle *FsPoll, cb FsPollCb, path string, interval uint) c.Int

//go:linkname FsPollStop C.uv_fs_poll_stop
func FsPollStop(handle *FsPoll) c.Int

//go:linkname FsPollClose C.uv_fs_poll_close
func FsPollClose(handle *FsPoll) c.Int

//go:linkname FsPollGetPath C.uv_fs_poll_getpath
func FsPollGetPath(handle *FsPoll) string

//TODO: Implemnt uv_poll_init_socket

// Cleanup cleans up the file system request.
func (f *File) Cleanup() {
	FsReqCleanup(f.req)
}

// Open opens a file specified by the path with given flags and mode, and returns a file descriptor.
func (f *File) Open(path string, flags int, mode int, cb FsCb) int {
	return FsOpen(f.loop, f.req, path, flags, mode, cb)
}

// Close closes a file descriptor.
func (f *File) Close(file int, cb FsCb) int {
	return FsClose(f.loop, f.req, file, cb)
}

// Read reads data from a file descriptor into a buffer at a specified offset.
func (f *File) Read(file int, buf []byte, offset int, cb FsCb) int {
	return FsRead(f.loop, f.req, file, buf, offset, cb)
}

// Write writes data to a file descriptor from a buffer at a specified offset.
func (f *File) Write(file int, buf []byte, offset int, cb FsCb) int {
	return FsWrite(f.loop, f.req, file, buf, offset, cb)
}

// Unlink deletes a file specified by the path.
func (f *File) Unlink(path string, cb FsCb) int {
	return FsUnlink(f.loop, f.req, path, cb)
}

// Mkdir creates a new directory at the specified path with a specified mode.
func (f *File) Mkdir(path string, mode int, cb FsCb) int {
	return FsMkdir(f.loop, f.req, path, mode, cb)
}

// Mkdtemp creates a temporary directory with a template path.
func (f *File) Mkdtemp(tpl string, cb FsCb) int {
	return FsMkdtemp(f.loop, f.req, tpl, cb)
}

// MkStemp creates a temporary file from a template path.
func (f *File) MkStemp(tpl string, cb FsCb) int {
	return FsMkStemp(f.loop, f.req, tpl, cb)
}

// Rmdir removes a directory specified by the path.
func (f *File) Rmdir(path string, cb FsCb) int {
	return FsRmdir(f.loop, f.req, path, cb)
}

// Stat retrieves status information about the file specified by the path.
func (f *File) Stat(path string, cb FsCb) int {
	return FsStat(f.loop, f.req, path, cb)
}

// Fstat retrieves status information about a file descriptor.
func (f *File) Fstat(file int, cb FsCb) int {
	return FsFstat(f.loop, f.req, file, cb)
}

// Rename renames a file from the old path to the new path.
func (f *File) Rename(path string, newPath string, cb FsCb) int {
	return FsRename(f.loop, f.req, path, newPath, cb)
}

// Fsync synchronizes a file descriptor's state with storage device.
func (f *File) Fsync(file int, cb FsCb) int {
	return FsFsync(f.loop, f.req, file, cb)
}

// Fdatasync synchronizes a file descriptor's data with storage device.
func (f *File) Fdatasync(file int, cb FsCb) int {
	return FsFdatasync(f.loop, f.req, file, cb)
}

// Ftruncate truncates a file to a specified length.
func (f *File) Ftruncate(file int, offset int, cb FsCb) int {
	return FsFtruncate(f.loop, f.req, file, offset, cb)
}

// Sendfile sends data from one file descriptor to another.
func (f *File) Sendfile(outFd int, inFd int, inOffset int, length int, cb FsCb) int {
	return FsSendfile(f.loop, f.req, outFd, inFd, inOffset, length, cb)
}

// Access checks the access permissions of a file specified by the path.
func (f *File) Access(path string, flags int, cb FsCb) int {
	return FsAccess(f.loop, f.req, path, flags, cb)
}

// Chmod changes the permissions of a file specified by the path.
func (f *File) Chmod(path string, mode int, cb FsCb) int {
	return FsChmod(f.loop, f.req, path, mode, cb)
}

// Fchmod changes the permissions of a file descriptor.
func (f *File) Fchmod(file int, mode int, cb FsCb) int {
	return FsFchmod(f.loop, f.req, file, mode, cb)
}

// Utime updates the access and modification times of a file specified by the path.
func (f *File) Utime(path string, atime int, mtime int, cb FsCb) int {
	return FsUtime(f.loop, f.req, path, atime, mtime, cb)
}

// Futime updates the access and modification times of a file descriptor.
func (f *File) Futime(file int, atime int, mtime int, cb FsCb) int {
	return FsFutime(f.loop, f.req, file, atime, mtime, cb)
}

// Lutime updates the access and modification times of a file specified by the path, even if the path is a symbolic link.
func (f *File) Lutime(path string, atime int, mtime int, cb FsCb) int {
	return FsLutime(f.loop, f.req, path, atime, mtime, cb)
}

// Link creates a new link to an existing file.
func (f *File) Link(path string, newPath string, cb FsCb) int {
	return FsLink(f.loop, f.req, path, newPath, cb)
}

// Symlink creates a symbolic link from the path to the new path.
func (f *File) Symlink(path string, newPath string, flags int, cb FsCb) int {
	return FsSymlink(f.loop, f.req, path, newPath, flags, cb)
}

// Readlink reads the target of a symbolic link.
func (f *File) Readlink(path string, cb FsCb) int {
	return FsReadlink(f.loop, f.req, path, cb)
}

// Realpath resolves the absolute path of a file.
func (f *File) Realpath(path string, cb FsCb) int {
	return FsRealpath(f.loop, f.req, path, cb)
}

// Copyfile copies a file from the source path to the destination path.
func (f *File) Copyfile(path string, newPath string, flags int, cb FsCb) int {
	return FsCopyfile(f.loop, f.req, path, newPath, flags, cb)
}

// Scandir scans a directory for entries.
func (f *File) Scandir(path string, flags int, cb FsCb) int {
	return FsScandir(f.loop, f.req, path, flags, cb)
}

// OpenDir opens a directory specified by the path.
func (f *File) OpenDir(path string, cb FsCb) int {
	return FsOpenDir(f.loop, f.req, path, cb)
}

// Readdir reads entries from an open directory.
func (f *File) Readdir(dir int, cb FsCb) int {
	return FsReaddir(f.loop, f.req, dir, cb)
}

// CloseDir closes an open directory.
func (f *File) CloseDir() int {
	return FsCloseDir(f.loop, f.req)
}

// Statfs retrieves file system status information.
func (f *File) Statfs(path string, cb FsCb) int {
	return FsStatfs(f.loop, f.req, path, cb)
}

// Chown Change file ownership
func (f *File) Chown(path string, uid int, gid int, cb FsCb) int {
	return FsChown(f.loop, f.req, path, uid, gid, cb)
}

// Fchown Change file ownership by file descriptor
func (f *File) Fchown(file int, uid int, gid int, cb FsCb) int {
	return FsFchown(f.loop, f.req, file, uid, gid, cb)
}

// Lchown Change file ownership (symlink)
func (f *File) Lchown(path string, uid int, gid int, cb FsCb) int {
	return FsLchown(f.loop, f.req, path, uid, gid, cb)
}

// Lstat Get file status (symlink)
func (f *File) Lstat(path string, cb FsCb) int {
	return FsLstat(f.loop, f.req, path, cb)
}

// Init Initialize a file event handle
func (e *FsEvent) Init(loop *Loop) int {
	return FsEventInit(loop, e)
}

// Start listening for file events
func (e *FsEvent) Start(cb FsEventCb, path string, flags int) int {
	return FsEventStart(e, cb, path, flags)
}

// Stop listening for file events
func (e *FsEvent) Stop() int {
	return FsEventStop(e)
}

// Close the file event handle
func (e *FsEvent) Close() int {
	return FsEventClose(e)
}

// GetPath Get the path of the file event
func (e *FsEvent) GetPath() string {
	return FsEventGetpath(e)
}

// Init Initialize a file poll handle
func (p *FsPoll) Init(loop *Loop) int {
	return FsPollInit(loop, p)
}

// Start polling for file changes
func (p *FsPoll) Start(cb FsPollCb, path string, interval uint) int {
	return FsPollStart(p, cb, path, interval)
}

// Stop polling for file changes
func (p *FsPoll) Stop() int {
	return FsPollStop(p)
}

// Close the file poll handle
func (p *FsPoll) Close() int {
	return FsPollClose(p)
}

// GetPath Get the path of the file poll
func (p *FsPoll) GetPath() string {
	return FsPollGetPath(p)
}
