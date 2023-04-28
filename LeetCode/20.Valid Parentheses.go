package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	m := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	temp := ""
	for _, v := range s {
		if strings.Contains("([{", string(v)) {
			temp += string(v)
		} else if len(temp) != 0 {
			if string(v) != m[temp[len(temp)-1:]] {
				return false
			}

			temp = temp[:len(temp)-1]
		} else {
			return false
		}
	}

	if len(temp) != 0 {
		return false
	}

	return true
}

func main() {
	result1 := isValid("()")
	fmt.Println(result1)

	result2 := isValid("()[]{}")
	fmt.Println(result2)

	result3 := isValid("(]")
	fmt.Println(result3)

	result4 := isValid("{[]}")
	fmt.Println(result4)

	result5 := isValid("]")
	fmt.Println(result5)

	result6 := isValid("((")
	fmt.Println(result6)
}
