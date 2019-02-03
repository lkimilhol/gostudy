package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //변수 정의 -> "" int -> 0
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
