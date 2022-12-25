package adventofcode2022

import "testing"

func Test_sumSizeOfDirsOverInSample(t *testing.T) {
	lines := readStringFile("day7_sample")
	act := sumSizeOfDirsUnder(lines, 100000)
	exp := 95437
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_sumSizeOfDirsOverInInput(t *testing.T) {
	lines := readStringFile("day7_input")
	act := sumSizeOfDirsUnder(lines, 100000)
	exp := 1582412
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getSizeOfDirToLeaveSpaceInSample(t *testing.T) {
	lines := readStringFile("day7_sample")
	act := getSizeOfDirToDelete(lines, 70000000, 30000000)
	exp := 24933642
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getSizeOfDirToLeaveSpaceInInput(t *testing.T) {
	lines := readStringFile("day7_input")
	act := getSizeOfDirToDelete(lines, 70000000, 30000000)
	exp := 3696336
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
