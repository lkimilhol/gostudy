package main //패키지가 main으로 되어 있어야 compile이 가능

import "fmt"

/**
Go 툴체인은 소스 프로그래과 그에 종속된 것들을 컴퓨터의 기본 기계어 명령으로 변환한다. 툴체인들은 go로 시작하는 하위 명령어를 통해 사용한다.
하위 명령 중 가장 간단한 것은 run으로 이름이 .go로 끝나는 한 개 이상의 소스 파일을 컴파일하고 라이브러리와 링크한 후 결과 실행 팡리을 구동한다.
*/

func main() {
	fmt.Println("Hello")
}

//gofmt -w를 통해서 코드 재작성 가능. 개발자는 이걸 이용해서 꾸준히 코드의 재작성을 습관처럼 들여야 한다.
