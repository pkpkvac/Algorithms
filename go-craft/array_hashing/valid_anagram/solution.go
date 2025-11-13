package validanagram

import (
	"fmt"
	"maps"
)

func isAnagram(s string, t string) bool {

	m1 := map[string]int{}
	m2 := map[string]int{}

	for _, char := range s {

		m1[string(char)]++

	}

	for _, char := range t {

		m2[string(char)]++

	}

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)

	return maps.Equal(m1, m2)
}
