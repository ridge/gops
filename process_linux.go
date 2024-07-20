package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func processes() ([]uint32, error) {
	des, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}
	out := []uint32{}
	for _, de := range des {
		pid, err := strconv.ParseUint(de.Name(), 10, 32)
		if err != nil { // There are non-process entries in /proc, skip them
			continue
		}
		out = append(out, uint32(pid))
	}
	return out, nil
}

func processPath(pid uint32) (string, error) {
	return filepath.EvalSymlinks(fmt.Sprintf("/proc/%d/exe", pid))
}
