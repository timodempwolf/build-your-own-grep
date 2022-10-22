package main

import (
	"bufio"
	// Uncomment this to pass the first stage
	// "bytes"
	"fmt"
	"io"
	"os"
)

// Usage: echo <input_text> | your_grep.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	ok, err := match(os.Stdin, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(4)
	}

	code := 0

	if !ok {
		code = 1
	}

	os.Exit(code)
}

func match(r io.Reader, pattern string) (bool, error) {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	var ok bool
	s := bufio.NewScanner(r)

	for s.Scan() {
		// Uncomment this to pass the first stage
		// line := s.Bytes()
		//
		// if bytes.IndexAny(line, pattern) != -1 {
		// 	ok = true
		// }
	}

	if err := s.Err(); err != nil {
		return false, fmt.Errorf("scan input: %w", err)
	}

	return ok, nil
}
