package sortpeople

import (
	"fmt"
	"sort"
)

func SortPeople(names []string, heights []int) ([]string, error) {

	if len(names) != len(heights) {
		return nil, fmt.Errorf("length mismatch: names=%d, heights=%d", len(names), len(heights))
	}

	if len(names) == 0 {
		return []string{}, nil
	}

	for i := 0; i < len(names)-1; i++ {

		for j := i + 1; j < len(heights); j++ {
			if heights[i] < heights[j] {
				tempHeight := heights[i]
				heights[i] = heights[j]
				heights[j] = tempHeight
				tempName := names[i]
				names[i] = names[j]
				names[j] = tempName
			}
		}
	}

	return names, nil
}

// SortPeopleOptimized - Optimized implementation using Go's built-in sort
func SortPeopleOptimized(names []string, heights []int) ([]string, error) {
	if len(names) != len(heights) {
		return nil, fmt.Errorf("length mismatch: names=%d, heights=%d", len(names), len(heights))
	}

	if len(names) == 0 {
		return []string{}, nil
	}

	// Create pairs of (name, height) for sorting
	type Person struct {
		name   string
		height int
	}

	people := make([]Person, len(names))
	for i := range names {
		people[i] = Person{names[i], heights[i]}
	}

	// Sort by height (descending) using Go's optimized sort
	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})

	// Extract names in sorted order
	result := make([]string, len(people))
	for i, person := range people {
		result[i] = person.name
	}

	return result, nil
}
