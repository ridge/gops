# gops replacement for GoLand

GoLand needs `gops` utility to list all Go processes for debugger.

This utility is a replacement of a small subset used by GoLand, namely:

- GoLand invokes `gops` without arguments.
- GoLand reads the output, and ignores everything except first word (PID).

## Why replacement

`gops` uses github.com/keybase/go-ps that uses cgo under macOS.

github.com/mitchellh/go-ps does not use cgo, but also does not have
Path() method used by `gops`.

To avoid cgo we have to reimplement pieces of `go-ps`, unfortunately.

## Legal

Copyright Tectonic Labs Ltd.

Licensed under [Apache 2.0](LICENSE) license.

Authors:
- [Mikhail Gusarov](https://github.com/dottedmag)
