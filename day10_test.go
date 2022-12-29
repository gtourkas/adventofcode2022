package adventofcode2022

import "testing"

func Test_calcSumOfSignalStrengthsInSample(t *testing.T) {
	lines := readStringFile("day10_sample")
	act := calcSumOfSignalStrengths(lines, []int{20, 60, 100, 140, 180, 220})
	exp := 13140
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcSumOfSignalStrengthsInInput(t *testing.T) {
	lines := readStringFile("day10_input")
	act := calcSumOfSignalStrengths(lines, []int{20, 60, 100, 140, 180, 220})
	exp := 12640
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_printCRTInSample(t *testing.T) {
	// NOTE: not a real test as we're visually inspecting the output
	lines := readStringFile("day10_sample")
	printCRT(lines)
}

func Test_printCRTInInput(t *testing.T) {
	// NOTE: not a real test as we're visually inspecting the output
	lines := readStringFile("day10_input")
	printCRT(lines)
	// EHBZLRJR
}
