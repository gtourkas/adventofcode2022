package adventofcode2022

import (
	"strconv"
	"strings"
)

func _parsePair(pair string) (int, int) {
	numbers := strings.Split(pair, "-")
	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])
	return a, b
}

func isContained(fa int, fb int, sa int, sb int) bool {
	if (fa >= sa && fb <= sb) || (sa >= fa && sb <= fb) {
		return true
	}
	return false
}

func isOverlapping(fa int, fb int, sa int, sb int) bool {
	if isContained(fa, fb, sa, sb) {
		return true
	}
	if (fa >= sa && fa <= sb) || (fb >= sa && fb <= sb) {
		return true
	}
	return false
}

func countContainedPairs(lines []string, cond func(int, int, int, int) bool) int {
	containedPairs := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		pairs := strings.Split(line, ",")
		fa, fb := _parsePair(pairs[0])
		sa, sb := _parsePair(pairs[1])
		if cond(fa, fb, sa, sb) {
			containedPairs++
		}
	}
	return containedPairs
}
