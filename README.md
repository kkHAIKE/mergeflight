# mergeflight
[![GoDoc](https://godoc.org/github.com/kkHAIKE/mergeflight?status.svg)](https://godoc.org/github.com/kkHAIKE/mergeflight)

like [singleflight](https://pkg.go.dev/golang.org/x/sync/singleflight) but use for merge diffrent parameter to batch call with count window and time window

it's useful to make batch RPC request or DB query, or other slow IO function.

# usage
see [godoc](https://godoc.org/github.com/kkHAIKE/mergeflight)
