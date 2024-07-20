package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
)

func isGo(pid uint32) bool {
	// Ignore "system" processes (MacOS reports some processes as having pid 0)
	if pid == 0 {
		return false
	}

	path, err := processPath(pid)
	// Ignore the process if could not determine its file path
	if err != nil {
		return false
	}

	// Parse file as Go binary, and see if it succeeded
	_, err = buildinfo.ReadFile(path)
	return err == nil
}

func main() {
	if len(os.Args) > 1 {
		fmt.Fprintln(os.Stderr, "This version of gops does not accept any arguments.")
		os.Exit(2)
	}

	pids, err := processes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read list of processes: %v\n", err)
		os.Exit(2)
	}

	for _, pid := range pids {
		if isGo(pid) {
			fmt.Println(pid)
		}
	}
}
