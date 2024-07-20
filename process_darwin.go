package main

import (
	"bytes"
	"syscall"
	"unsafe"
)

// macOS syscalls
const (
	procIinfoCallLlistPIDs = 0x01
	procListPIDs           = 1

	procIinfoCallPIDInfo = 0x02
	procPIDPathInfo      = 11
)

const pidsBufLen = 1024 * 1024

func processes() ([]uint32, error) {
	out := make([]uint32, pidsBufLen, pidsBufLen)

	_, _, errno := syscall.Syscall6(
		syscall.SYS_PROC_INFO,
		procIinfoCallLlistPIDs,
		procListPIDs,
		0,
		0,
		uintptr(unsafe.Pointer(&out[0])),
		pidsBufLen*(unsafe.Sizeof(out[0])))

	if errno != 0 {
		return nil, errno
	}

	return out, nil
}

func cString(s []byte) string {
	i := bytes.IndexByte(s, 0)
	if i != -1 {
		return string(s[:i])
	}
	return string(s)
}

const maxPathSize = 4096

func processPath(pid uint32) (string, error) {
	buffer := make([]byte, maxPathSize)

	_, _, errno := syscall.Syscall6(
		syscall.SYS_PROC_INFO,
		procIinfoCallPIDInfo,
		uintptr(pid),
		procPIDPathInfo,
		0,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(maxPathSize))

	if errno != 0 {
		return "", errno
	}

	return cString(buffer), nil
}
