package lemonadechange

func lemonadeChange(bills []int) bool {

	fives := 0
	tens := 0
	// twenny := 0

	for _, bill := range bills {

		// if it's a 5, you just add five

		switch bill {

		case 5:
			fives += 1

		case 10:

			if fives == 0 {
				return false
			}
			fives -= 1
			tens += 1

		case 20:

			if tens > 0 && fives > 0 {
				//
				tens -= 1
				fives -= 1
			} else if fives > 2 {
				fives -= 3
			} else {
				return false
			}

		}

		// if it's a 10, you need to give 5 back

		// if it's a 20, you need to give either

		// a 10 and a 5 back

		// or 3 5's back

		// greedy will give back the biggest bills first

	}

	return true

}
