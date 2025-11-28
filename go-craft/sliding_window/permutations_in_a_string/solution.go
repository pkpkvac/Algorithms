package permutationsinastring

func checkInclusion(s1 string, s2 string) bool {

	s1a := [26]int{0}
	s2a := [26]int{0}

	// first go through s1a
	maxWindowLen := len(s1)
	currWindowLen := 0
	for _, c := range s1 {
		s1a[c-97]++
	}

	l := 0
	r := 0

	for r < len(s2) {

		currWindowLen = r - l + 1

		s2a[s2[r]-97]++

		if maxWindowLen == currWindowLen {
			if s1a == s2a {
				return true
			}
			s2a[s2[l]-97]--
			l++
		}
		if currWindowLen > maxWindowLen {
			// check if the solution is it,
			//  if not, move l forward
			s2a[s2[l]-97]--
			l++
		}
		r++
	}

	return false
}
