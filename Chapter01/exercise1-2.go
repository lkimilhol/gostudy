package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := " ", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	for i := 1; i < len(os.Args); i++ {
		fmt.Print(i)
		fmt.Println(" " + os.Args[i])
	}
}
