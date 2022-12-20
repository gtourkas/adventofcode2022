package adventofcode2022

import cb "github.com/emirpasic/gods/queues/circularbuffer"

func findFirstMarker(input string, last int) int {
	lastCharacters := cb.New(last)
	for i, char := range input {

		lastCharacters.Enqueue(char)

		it := lastCharacters.Iterator()
		prevChars := make(map[int32]bool, last)

		uniqueChars := true
		for it.Next() {
			prevChar := it.Value()
			_, ok := prevChars[prevChar.(int32)]
			if ok {
				uniqueChars = false
				break
			}
			prevChars[prevChar.(int32)] = true
		}

		if uniqueChars && i >= last-1 {
			return i + 1
		}
	}
	return 0
}
