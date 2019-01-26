package main

import (
	"fmt"
	"os"
)

/**
연습문제 1.1
echo 프로그램을 수정해 호출한 명령인 os.Args[0]도 같이 출력하라.
*/

func main() {
	s, sep := " ", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
