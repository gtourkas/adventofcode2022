package adventofcode2022

import (
	"testing"
)

func Test_getTopStacks_MoveOneByOne_InSample(t *testing.T) {
	lines := readStringFile("day5_sample")
	act := getTopStacks(lines, moveOneByOne)
	exp := "CMZ"
	if exp != act {
		t.Errorf("expected %s but actual is %s", exp, act)
	}
}

func Test_getTopStacks_MoveOneByOne_InInput(t *testing.T) {
	lines := readStringFile("day5_input")
	act := getTopStacks(lines, moveOneByOne)
	exp := "FZCMJCRHZ"
	if exp != act {
		t.Errorf("expected %s but actual is %s", exp, act)
	}
}

func Test_getTopStacks_MoveAllTogether_InSample(t *testing.T) {
	lines := readStringFile("day5_sample")
	act := getTopStacks(lines, moveAllTogether)
	exp := "MCD"
	if exp != act {
		t.Errorf("expected %s but actual is %s", exp, act)
	}
}

func Test_getTopStacks_MoveAllTogether_InInput(t *testing.T) {
	lines := readStringFile("day5_input")
	act := getTopStacks(lines, moveAllTogether)
	exp := "JSDHQMZGF"
	if exp != act {
		t.Errorf("expected %s but actual is %s", exp, act)
	}
}
