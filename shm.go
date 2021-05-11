package shm

import (
	"fmt"
	"io"
)

// Memory is shared memory struct
type Memory struct {
	m   *shmi
	pos int64
}

// Create is create shared memory
func Create(name string, size int32) (*Memory, error) {
	m, err := create(name, size)
	if err != nil {
		return nil, err
	}
	return &Memory{m, 0}, nil
}

// Open is open exist shared memory
func Open(name string, size int32) (*Memory, error) {
	m, err := open(name, size)
	if err != nil {
		return nil, err
	}
	return &Memory{m, 0}, nil
}

// Close is close & discard shared memory
func (o *Memory) Close() (err error) {
	fmt.Println("NEW CLOSE METHOD")
	// This code makes sure not to crash if the object pointer is nil - which can happen.
	if o == nil {
		return nil
	}
	if o.m != nil {
		// Refactor for simplicity.
		err = o.m.close()
		o.m = nil
	}
	return err
}

// Read is read shared memory (current position)
func (o *Memory) Read(p []byte) (n int, err error) {
	n, err = o.ReadAt(p, o.pos)
	if err != nil {
		return 0, err
	}
	o.pos += int64(n)
	return n, nil
}

// ReadAt is read shared memory (offset)
func (o *Memory) ReadAt(p []byte, off int64) (n int, err error) {
	return o.m.readAt(p, off)
}

// Seek is move read/write position at shared memory
func (o *Memory) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		offset += int64(0)
	case io.SeekCurrent:
		offset += o.pos
	case io.SeekEnd:
		offset += int64(o.m.size)
	}
	if offset < 0 || offset >= int64(o.m.size) {
		return 0, fmt.Errorf("invalid offset")
	}
	o.pos = offset
	return offset, nil
}

// Write is write shared memory (current position)
func (o *Memory) Write(p []byte) (n int, err error) {
	n, err = o.WriteAt(p, o.pos)
	if err != nil {
		return 0, err
	}
	o.pos += int64(n)
	return n, nil
}

// WriteAt is write shared memory (offset)
func (o *Memory) WriteAt(p []byte, off int64) (n int, err error) {
	return o.m.writeAt(p, off)
}
