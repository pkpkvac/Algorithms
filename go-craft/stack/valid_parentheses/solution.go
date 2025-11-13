package validparentheses

func isValid(s string) bool {

	stack := []string{}

	for _, char := range s {

		switch string(char) {
		case "(":
			stack = append(stack, string(char))
		case "[":
			stack = append(stack, string(char))
		case "{":
			stack = append(stack, string(char))
		case ")":
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "]":
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != "[" {
				return false
			}
			stack = stack[:len(stack)-1]

		case "}":
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != "{" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
