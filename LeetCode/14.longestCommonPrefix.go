package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	temp := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], temp) {
			temp = temp[:len(temp)-1]
		}
	}

	return temp
}

func main() {
	result1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(result1)

	result2 := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println(result2)
}
