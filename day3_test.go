package adventofcode2022

import (
	"testing"
)

func Test_calcSumOfPrioritiesInSample(t *testing.T) {
	lines := readStringFile("day3_sample")
	act := calcSumOfPriorities(lines)
	exp := 157
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcSumOfPrioritiesInFile(t *testing.T) {
	lines := readStringFile("day3_input")
	act := calcSumOfPriorities(lines)
	exp := 7727
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcSumOfPrioritiesPerGroupInSample(t *testing.T) {
	lines := readStringFile("day3_sample")
	act := calcSumOfPrioritiesForGroups(lines, 3)
	exp := 70
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcSumOfPrioritiesPerGroupInFile(t *testing.T) {
	lines := readStringFile("day3_input")
	act := calcSumOfPrioritiesForGroups(lines, 3)
	exp := 2609
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
