package adventofcode2022

import "testing"

func Test_countTailVisitsWith2KnotRopesInSample(t *testing.T) {
	lines := readStringFile("day9_sample")
	act := countTailVisits(lines, 2)
	exp := 13
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countTailVisitsWith2KnotRopesInInput(t *testing.T) {
	lines := readStringFile("day9_input")
	act := countTailVisits(lines, 2)
	exp := 6266
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countTailVisitsWith10KnotRopesInSample(t *testing.T) {
	lines := readStringFile("day9_sample")
	act := countTailVisits(lines, 10)
	exp := 1
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countTailVisitsWith10KnotRopesInSample2(t *testing.T) {
	lines := readStringFile("day9_sample2")
	act := countTailVisits(lines, 10)
	exp := 36
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_countTailVisitsWith10KnotRopesInInputz(t *testing.T) {
	lines := readStringFile("day9_input")
	act := countTailVisits(lines, 10)
	exp := 2369
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
