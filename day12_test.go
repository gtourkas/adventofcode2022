package adventofcode2022

import (
	"testing"
)

func Test_findShortestPathInSample_StartFromDesignatedPoint(t *testing.T) {
	lines := readStringFile("day12_sample")
	act := findShortestPath(lines, false)
	exp := 31
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_findShortestPathInInput_StartFromDesignatedPoint(t *testing.T) {
	lines := readStringFile("day12_input")
	act := findShortestPath(lines, false)
	exp := 370
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_findShortestPathInSample_StartFromLowestElevation(t *testing.T) {
	lines := readStringFile("day12_sample")
	act := findShortestPath(lines, true)
	exp := 29
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_findShortestPathInInput_StartFromLowestElevation(t *testing.T) {
	lines := readStringFile("day12_input")
	act := findShortestPath(lines, true)
	exp := 363
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
