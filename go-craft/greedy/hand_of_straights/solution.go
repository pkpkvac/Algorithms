package hand_of_straights

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {

	m := make(map[int]int)

	if len(hand) % groupSize != 0 {
		return false;
	}

	for _, card := range hand {
		m[card]++
	}

	// sort unique cards
	var uniqueCards []int
	for card := range m {
		uniqueCards = append(uniqueCards, card)
	}
	
	// sort unique cards
	sort.Ints(uniqueCards)

	// for each card, starting from the smallest
	for _, card := range uniqueCards {
		count := m[card]
		if count == 0 {
			continue
		}
		// try to form a group starting from this card
		for i := 0; i < groupSize; i++ {
			nextCard := card + i
			if m[nextCard] < count {
				return false
			}
			m[nextCard] -= count
		}
	}

	return true
}
