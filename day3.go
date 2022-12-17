package adventofcode2022

import (
	"sort"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func calcSumOfPriorities(lines []string) int {
	sumOfPriorities := 0
	for _, l := range lines {
		compLen := len(l) / 2
		firstComp := l[:compLen]
		secondComp := l[compLen:]

		firstCompRunes := []rune(firstComp)
		secondCompRunes := []rune(secondComp)
		sort.Slice(firstCompRunes, func(i int, j int) bool { return firstCompRunes[i] < firstCompRunes[j] })
		sort.Slice(secondCompRunes, func(i int, j int) bool { return secondCompRunes[i] < secondCompRunes[j] })

		i := 0
		j := 0
		for {
			f := firstCompRunes[i]
			s := secondCompRunes[j]
			if f == s {
				sumOfPriorities += strings.Index(alphabet, string(f)) + 1
				break
			}
			if f > s {
				j++
			}
			if f < s {
				i++
			}
			if i >= compLen && j >= compLen {
				break
			}
		}
	}
	return sumOfPriorities
}

func calcSumOfPrioritiesForGroups(lines []string, group int) int {
	sumOfPriorities := 0
	for i := 0; i < len(lines); i += group {
		top := struct {
			letter uint8
			count  int
		}{
			'A',
			0,
		}
		counts := make(map[uint8]int)

		gl := 1
		for j := i; j < i+group; j++ {
			line := lines[j]
			for k := 0; k < len(line); k++ {
				l := line[k]
				if counts[l] == gl-1 {
					counts[l] += gl
				}
				if counts[l] >= top.count {
					top.count = counts[l]
					top.letter = l
				}
			}
			gl++
		}
		sumOfPriorities += strings.Index(alphabet, string(top.letter)) + 1
	}
	return sumOfPriorities
}
