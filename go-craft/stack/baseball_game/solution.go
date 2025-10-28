package baseballgame

import (
	"errors"
	"fmt"
	"strconv"
)

func BaseballGame(operations []string) (int, error) {

	// Input validation
	if len(operations) == 0 {
		return 0, errors.New("empty operations")
	}

	// initialize a stack
	stack := []int{}

	for _, operation := range operations {

		stack = processOperation(stack, operation)

	}

	fmt.Println("final stack", stack)
	// Sum all remaining scores
	sum := 0
	for _, score := range stack {
		sum += score
	}

	return sum, nil
}

func processOperation(stack []int, operation string) []int {
	switch operation {

	case "D":
		// take the top element,
		// double it, append it to the stack
		if len(stack) > 0 {
			lastScore := stack[len(stack)-1]
			stack = append(stack, lastScore*2)
			return stack
		}

	case "C":
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
			return stack
		}

	case "+":
		if len(stack) >= 2 {
			lastTwoSum := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, lastTwoSum)
			return stack
		}

	default:

		if num, err := strconv.Atoi(operation); err == nil {
			if num >= -30000 && num <= 30000 {
				stack = append(stack, num)
				return stack
			}
		}

	}

	return stack
}
