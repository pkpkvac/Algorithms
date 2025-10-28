package validpalindromeii

import "errors"

func ValidPalindromeII(s string) (bool, error) {

	// Input validation
	if len(s) == 0 {
		return false, errors.New("empty string")
	}

	// Only allow lowercase English letters as per problem constraints
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false, errors.New("string must contain only lowercase English letters")
		}
	}

	l := 0
	r := len(s) - 1
	var res bool

	for l < r {

		if s[l] == s[r] {
			r--
			l++
		} else {
			res = isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
			return res, nil
		}

	}

	return true, nil
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

}

func isPalindrome(s string, l int, r int) bool {

	for l < r {
		if s[l] != s[r] {
			return false
		}
		r--
		l++
	}
	return true
}
