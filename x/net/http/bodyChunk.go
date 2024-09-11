package http

import (
	"errors"
	"io"

	"github.com/goplus/llgo/c/libuv"
)

type bodyChunk struct {
	chunk       []byte
	readCh      chan []byte
	asyncHandle *libuv.Async

	done chan struct{}

	rerr error
}

var (
	errClosedBodyChunk = errors.New("bodyChunk: read/write on closed body")
)

func newBodyChunk(asyncHandle *libuv.Async) *bodyChunk {
	return &bodyChunk{
		readCh:      make(chan []byte, 1),
		done:        make(chan struct{}),
		asyncHandle: asyncHandle,
	}
}

func (bc *bodyChunk) Read(p []byte) (n int, err error) {
	for n < len(p) {
		if len(bc.chunk) == 0 {
			select {
			case chunk, ok := <-bc.readCh:
				if !ok {
					if n > 0 {
						return n, nil
					}
					return 0, bc.readCloseError()
				}
				bc.chunk = chunk
				bc.asyncHandle.Send()
			case <-bc.done:
				if n > 0 {
					return n, nil
				}
				return 0, io.EOF
			}
		}

		copied := copy(p[n:], bc.chunk)
		n += copied
		bc.chunk = bc.chunk[copied:]
	}

	return n, nil
}

func (bc *bodyChunk) Close() error {
	return bc.closeRead(nil)
}

func (bc *bodyChunk) readCloseError() error {
	if rerr := bc.rerr; rerr != nil {
		return rerr
	}
	return errClosedBodyChunk
}

func (bc *bodyChunk) closeRead(err error) error {
	if err == nil {
		err = io.EOF
	}
	bc.rerr = err
	close(bc.done)
	return nil
}
