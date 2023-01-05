package adventofcode2022

import "testing"

func Test_calcUnitsOfSandsInRestInSample_NoFloor(t *testing.T) {
	lines := readStringFile("day14_sample")
	act := calcUnitsOfSandsAtRest(lines, false)
	exp := 24
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcUnitsOfSandsInRestInInput_NoFloor(t *testing.T) {
	lines := readStringFile("day14_input")
	act := calcUnitsOfSandsAtRest(lines, false)
	exp := 674
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcUnitsOfSandsInRestInSample_Floor(t *testing.T) {
	lines := readStringFile("day14_sample")
	act := calcUnitsOfSandsAtRest(lines, true)
	exp := 93
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcUnitsOfSandsInRestInInput_Floor(t *testing.T) {
	lines := readStringFile("day14_input")
	act := calcUnitsOfSandsAtRest(lines, true)
	exp := 24958
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
