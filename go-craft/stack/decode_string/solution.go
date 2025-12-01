package decode_string

import (
	"strconv"
	"strings"
)

func reverseString(s string) string {
	bytes := []byte(s)

	l := 0
	r := len(bytes) - 1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}

	return string(bytes)

}

func decodeString(s string) string {

	stack := []string{}

	for _, char := range s {

		if char == ']' {
			// do processing

			// take top character
			// pop off stack
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			substr := ""

			for c != "[" {

				substr = c + substr

				c = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

			}
			// multiplier is next character
			mult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			i, _ := strconv.Atoi(mult)

			sStr := ""
			// substr = reverseString(substr)

			for i > 0 {
				sStr += substr
				i--
			}

			stack = append(stack, strings.TrimSpace(sStr))

		} else {
			// add it to the stack
			stack = append(stack, string(char))
		}

	}

	return strings.Join(stack, "")
}
