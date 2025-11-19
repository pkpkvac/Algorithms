package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {

	res := [][]string{}
	m := map[[26]int][]string{}

	for _, str := range strs {
		// each string will get its own temporary array of length 26, representing
		// the character frequencies
		// all characters are lowercase, and ascii dec a = 97

		charFreq := [26]int{}

		for _, c := range str {
			idx := c - 97
			charFreq[idx]++
		}

		m[charFreq] = append(m[charFreq], str)
	}

	for _, value := range m {
		res = append(res, value)
	}

	sort.Slice(res, func(i, j int) bool { return len(res[i]) < len(res[j]) })

	return res

}
