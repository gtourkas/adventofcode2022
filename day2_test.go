package adventofcode2022

import "testing"

func Test_calcTotalScorePart1InSample(t *testing.T) {
	lines := readStringFile("day2_sample")
	act := calcTotalScorePart1(lines)
	exp := 15
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcTotalScorePart1InInput(t *testing.T) {
	lines := readStringFile("day2_input")
	act := calcTotalScorePart1(lines)
	exp := 11449
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcTotalScorePart2InSample(t *testing.T) {
	lines := readStringFile("day2_sample")
	act := calcTotalScorePart2(lines)
	exp := 12
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcTotalScorePart2InInput(t *testing.T) {
	lines := readStringFile("day2_input")
	act := calcTotalScorePart2(lines)
	exp := 13187
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
