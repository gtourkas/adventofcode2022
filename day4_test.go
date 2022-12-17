package adventofcode2022

import "testing"

func Test_countContainedPairsInSample(t *testing.T) {
	lines := readStringFile("day4_sample")
	act := countContainedPairs(lines, isContained)
	exp := 2
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countContainedPairsInInput(t *testing.T) {
	lines := readStringFile("day4_input")
	act := countContainedPairs(lines, isContained)
	exp := 595
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countOverlappingPairsInSample(t *testing.T) {
	lines := readStringFile("day4_sample")
	act := countContainedPairs(lines, isOverlapping)
	exp := 4
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countOverlappingPairsInInput(t *testing.T) {
	lines := readStringFile("day4_input")
	act := countContainedPairs(lines, isOverlapping)
	exp := 952
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
