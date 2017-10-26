# shm

[![Build Status](https://travis-ci.org/hidez8891/shm.svg?branch=master)](https://travis-ci.org/hidez8891/shm)

shm is Golang shared memory library.

## Example

```go
w, _ := shm.Create("shm_name", 256)
defer w.Close()

r, _ := shm.Open("shm_name", 256)
defer r.Close()

wbuf := []byte("Hello World")
w.Write(wbuf)

rbuf := make([]byte, len(wbuf))
r.Read(rbuf)
// rbuf == wbuf
```
