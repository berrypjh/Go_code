package main

import (
	"fmt"
)

func isValid(s string) bool {
    return true
}

func main() {
	result1 := isValid("()")
	fmt.Println(result1)

	result2 := isValid("()[]{}")
	fmt.Println(result2)

	result3 := isValid("(]")
	fmt.Println(result3)
}
