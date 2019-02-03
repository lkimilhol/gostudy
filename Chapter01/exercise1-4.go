package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
It is print file's name and path
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2L %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, fileNames map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fileNames[f.Name()]++
	}
	// NOTE: input.Err()에서의 잠재적 오류는 무시한다.
}
