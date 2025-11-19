package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {

	maxL := 0
	if len(s) == 0 {
		return maxL
	}

	m := map[byte]int{}

	l := 0
	r := 0

	for r < len(s) {

		for m[s[r]] > 0 {

			// character exists in window

			// reduce the item from m (coming from l)
			m[s[l]]--
			// move left forward
			l++

		}
		// otherwise, does not exist in window
		// include it
		m[s[r]]++

		L := r - l + 1

		maxL = max(L, maxL)
		r++
	}

	return maxL
}
