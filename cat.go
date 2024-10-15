package main

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	filePath := os.Args[1]

	// open(path string, mode int, perm uint32)
	fd, err := unix.Open(filePath, unix.O_RDONLY, 0) // only reading no need to add permissions
	if err != nil {
		// ENOENT: no such file or directory
		if err == unix.ENOENT {
			// log.Printf("cat: %s: No such file or directory", filePath)
			errMsg := []byte("no such file or directory")
			unix.Write(unix.Stderr, errMsg)
		}
		unix.Exit(1)
	}
	defer unix.Close(fd)

	buf := make([]byte, 1024)

	for {
		// n indicates number of bytes read, if zero we reach end of file
		n, err := unix.Read(fd, buf)
		if err != nil {
			log.Printf("Failed to read file '%s': %v", filePath, err)
			unix.Exit(1)
		}
		// eof
		if n == 0 {
			break
		}

		_, err = unix.Write(unix.Stderr, buf)
		if err != nil {
			log.Printf("Failed to write to stdout, %v", err)
			unix.Exit(1)
		}
	}

	unix.Exit(0)
}
