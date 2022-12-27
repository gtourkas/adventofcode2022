package adventofcode2022

import "testing"

func Test_countVisibleTreesInSample(t *testing.T) {
	lines := readStringFile("day8_sample")
	act := countVisibleTrees(lines)
	exp := 21
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countVisibleTreesInInput(t *testing.T) {
	lines := readStringFile("day8_input")
	act := countVisibleTrees(lines)
	exp := 1713
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcHighestScenicScoreInSample(t *testing.T) {
	lines := readStringFile("day8_sample")
	act := calcHighestScenicScore(lines)
	exp := 8
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcHighestScenicScoreInInput(t *testing.T) {
	lines := readStringFile("day8_input")
	act := calcHighestScenicScore(lines)
	exp := 268464
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
