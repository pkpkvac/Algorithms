package longestrepeatingcharacterreplacement

import "errors"

func LongestRepeatingCharacterReplacement(s string, k int) (int, error) {

	// Input validation
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}

	// only allow uppercase English letters as per problem constraints
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return 0, errors.New("string must contain only uppercase English letters")
		}
	}

	// initialize a frequency map
	freqMap := make(map[string]int)
	for _, char := range s {
		freqMap[string(char)] = 0
	}

	// initialize a left pointer
	l := 0
	// initialize a right pointer
	r := 0

	// initialize a maximum sequence length
	maxSeq := 0

	windowSize := 0

	for r < len(s) {
		// add the character at the right pointer to the frequency map
		freqMap[string(s[r])]++

		maxFreq := GetMaxFrequency(freqMap)

		windowSize = r - l + 1

		// tricky part: if the window size minus the max frequency is greater than k, we need to shrink the window from the left
		// this is because we can only replace k characters, so if the window size minus the max frequency is greater than k, we need to shrink the window from the left
		if windowSize-maxFreq > k {
			// shrink the window from the left
			freqMap[string(s[l])]--
			l++

		}

		windowSize = r - l + 1
		maxSeq = max(maxSeq, windowSize)
		r++

	}

	return maxSeq, nil
}

func GetMaxFrequency(freqMap map[string]int) int {
	maxFreq := 0
	for _, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return maxFreq
}
