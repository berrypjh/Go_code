package main

import (
	"fmt"
)

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	temp := 0
	for i := len(s) - 1; i >= 0; i-- {
		n := m[s[i]]
		if temp <= n {
			result += n
		} else {
			result -= n
		}
		temp = m[s[i]]
	}

	return result
}

func main() {
	result1 := romanToInt("III")
	fmt.Println(result1)

	result2 := romanToInt("LVIII")
	fmt.Println(result2)

	result3 := romanToInt("MCMXCIV")
	fmt.Println(result3)
}
