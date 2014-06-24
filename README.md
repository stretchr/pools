pools [![GoDoc](https://godoc.org/github.com/stretchr/pools?status.png)](http://godoc.org/github.com/stretchr/pools) [![wercker status](https://app.wercker.com/status/6e99ed192553663550da37215737bdbd/s "wercker status")](https://app.wercker.com/project/bykey/6e99ed192553663550da37215737bdbd)
=====

pools is a go package for managing a suite of differently sized slices of objects backed by sync.Pool


Usage
=====

Usage is almost identical to sync.Pool:

```go
var s pools.Suite
s.New = func() interface{} {
	return MyObject{}
}
slice := s.Get(100) // returns a []MyObject of length 100
s.Put(slice) // put it back in the appropriate pool
```
