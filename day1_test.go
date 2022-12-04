package adventofcode2022

import "testing"

func Test_getMaxCaloriesInSample(t *testing.T) {
	lines := readStringFile("day1_sample")
	act := getMaxCalories(lines)
	exp := 24000
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getMaxCaloriesInInput(t *testing.T) {
	lines := readStringFile("day1_input")
	act := getMaxCalories(lines)
	exp := 67658
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getSumOfTopCaloriesInSample(t *testing.T) {
	lines := readStringFile("day1_sample")
	act := getSumOfTopCalories(lines, 3)
	exp := 45000
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getSumOfTopCaloriesInInput(t *testing.T) {
	lines := readStringFile("day1_input")
	act := getSumOfTopCalories(lines, 3)
	exp := 200158
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
