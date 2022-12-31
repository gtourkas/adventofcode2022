package adventofcode2022

import "testing"

func Test_calcMonkeyBusinessInSample_Rounds20_DivideWorryLevel3(t *testing.T) {
	lines := readStringFile("day11_sample")
	act := calcMonkeyBusiness(lines, 20, 3)
	exp := 10605
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcMonkeyBusinessInInput_Rounds20_DivideWorryLevel3(t *testing.T) {
	lines := readStringFile("day11_input")
	act := calcMonkeyBusiness(lines, 20, 3)
	exp := 57348
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcMonkeyBusinessInSample_Rounds2000_DivideWorryLevel1(t *testing.T) {
	lines := readStringFile("day11_sample")
	act := calcMonkeyBusiness(lines, 10000, 0)
	exp := 2713310158
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcMonkeyBusinessInInput_Rounds2000_DivideWorryLevel1(t *testing.T) {
	lines := readStringFile("day11_input")
	act := calcMonkeyBusiness(lines, 10000, 0)
	exp := 14106266886
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
