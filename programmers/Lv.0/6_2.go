// 직각삼각형 출력하기
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
