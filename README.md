# mergeflight
[![GoDoc](https://godoc.org/github.com/kkHAIKE/mergeflight?status.svg)](https://godoc.org/github.com/kkHAIKE/mergeflight)

It is similar to [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight), but it is used to merge different parameters into batch calls with a count window and time window.

This is useful for making batch RPC requests or DB queries, or for other slow I/O functions.

# usage
see [godoc](https://godoc.org/github.com/kkHAIKE/mergeflight)
