package stack

import (
	"fmt"
)

func isValid(s string) bool {

	if s == "" {
		return true
	}

	d := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}
	for _, c := range s {
		if len(stack) == 0 {
			if _, ok := d[c]; ok {
				return false
			} else {
				stack = append(stack, c)
				continue
			}
		}
		if k, ok := d[c]; ok {
			if stack[len(stack)-1] == k {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	b := isValid("[][][[[]]")
	fmt.Println("isValid: ", b)
}
