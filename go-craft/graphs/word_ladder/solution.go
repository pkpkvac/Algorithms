package wordladder

// hint 1:
// Consider the given problem in terms of a graph, treating strings as nodes.
// Think of a way to build edges where two strings have an edge if they differ by
// a single character. A naive approach would be to consider each pair of strings
// and check whether an edge can be formed. Can you think of an efficient way?
// For example, consider a string hot and think about the strings that can be
// formed from it by changing a single letter.

// Hint 2
// To efficiently build edges, consider transforming each word into intermediate
// states by replacing one character with a wildcard, like *. For example,
// the word hot can be transformed into *ot, h*t, and ho*. These intermediate
// states act as "parents" that connect words differing by one character.
// For instance, *ot can connect to words like cot. For each word in the list
// , generate all possible patterns by replacing each character with * and store
// the word as a child of these patterns. We can run a BFS starting from the beginWord,
// visiting other words while avoiding revisiting by using a hash set.

//  Hint 3
//  When visiting a node during BFS, if the word matches the endWord, we
//  immediately return true. Otherwise, we generate the pattern words
//  that can be formed from the current word and attempt to visit the
//  words connected to these pattern words. We add only unvisited words
//  to the queue. If we exhaust all possibilities without finding the endWord,
//  we return false.

// BFS finds the shortest path in an unweighted graph no matter how you obtain the
// neighbors, as long as you’re not adding fake edges or missing real ones

func ladderLength(beginWord string, endWord string, wordList []string) int {

	// for a word w, neighnors are computed 'a'...'z'
	// at each position i, that letter, and if the new string
	// is in the wordList, and is not w, it's a neigbor

	// need a set/or map, of wordList for O(1) access, so you're not
	// always doing the lookup
	wL := make(map[string]string)

	// queue start with the beginWord, and goes through each character
	// in the word, a-z, replacing and enqueueing that word if it's in word list
	bfsQ := []string{beginWord}

	visited := make(map[string]string)
	visited[beginWord] = beginWord

	// do not need an adjacency list, or pattern nodes in memory
	// graph exists 'implicitly', under the rule "neighnors = one letter change AND in wordList"

	for _, word := range wordList {
		wL[word] = word
	}

	count := 1

	if _, inList := wL[endWord]; !inList {
		return 0
	}

	for len(bfsQ) > 0 {

		layerSize := len(bfsQ)

		for layerSize > 0 {
			w := bfsQ[0]
			bfsQ = bfsQ[1:]

			if w == endWord {
				return count
			}

			// at each position i, try 'a'..'z'
			for i := 0; i < len(w); i++ {
				for c := byte('a'); c <= byte('z'); c++ {
					if w[i] == c {
						// if the character that is at index i is the
						// same as the new character, skip
						continue
					}

					newWord := w[:i] + string(c) + w[i+1:]

					if _, ok := visited[newWord]; !ok {
						// implies we haven't visited the word,
						// so either the word is the endWord, or can
						// bring us closer to the endWord.

						if _, inList := wL[newWord]; inList {
							// we also need this newWord to exist in the wordList

							if newWord != w {

								// the newWord is the endWord
								// we're done.
								if newWord == endWord {
									return count + 1
								}
								if newWord != w {
									// add this newWord to the queue
									bfsQ = append(bfsQ, newWord)
									// mark it as visited
									visited[newWord] = newWord
								}

							}
						}
					}

				}
			}
			layerSize--

		}
		count++

	}

	return 0
}
