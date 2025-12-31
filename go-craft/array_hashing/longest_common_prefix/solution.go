package longestcommonprefix

func longestCommonPrefix(strs []string) string {

	lcp := strs[0]

	for i, word := range strs {

		for j, letter := range word {

			// if len(word)

			if lcp[j] != byte(letter) {
				// we can break from this loop, knowing lcp
				if j == 0 {
					return ""
				}
				lcp = lcp[:j-1]
				break
			}
		}

	}

}
