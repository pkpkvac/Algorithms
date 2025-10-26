package validpalindromeii

import (
	"errors"
	"fmt"
	"strings"
)

func ValidPalindromeII(s string) (bool, error) {

	// handle empty string, spaces, numbers, or capitalized characters
	if len(s) == 0 || strings.ContainsAny(s, " ") || strings.ContainsAny(s, "0123456789") || strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false, errors.New("invalid input")
	}

	r := len(s) - 1

	for l := 0; l < r; l++ {

		fmt.Printf("Val%c\n", s[l])

	}

	// 	Hint
	// Either skip one character from left

	// Or skip one character from right

	// If any skipping returns true then it's a valid palindrome.
	// 	Use the standard two-pointer approach, with a left pointer at the start and a right pointer at the end of the string.
	// Iterate, moving the pointers inwards as long as s[left] and s[right] match.
	// If we encounter a mismatch, this is our "fork in the road." Now we can use our one allowed deletion. We immediately stop the main loop and check two possibilities:
	// Does the string become a palindrome if we delete the character at left? (i.e., we check the substring from left + 1 to right).

	// OR, does the string become a palindrome if we delete the character at right? (i.e., we check the substring from left to right - 1).

	// If the main loop finishes without finding any mismatches, the string is already a palindrome, so we return true.

	return true, nil
}
