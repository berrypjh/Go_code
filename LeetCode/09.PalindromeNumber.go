package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	arr := strconv.Itoa(x)

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	result1 := isPalindrome(121)
	fmt.Println(result1)

	result2 := isPalindrome(-121)
	fmt.Println(result2)

	result3 := isPalindrome(10)
	fmt.Println(result3)
}
