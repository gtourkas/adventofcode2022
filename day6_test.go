package adventofcode2022

import (
	"fmt"
	"testing"
)

func Test_findFirstMarkerLast4InSample(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			act := findFirstMarker(tt.input, 4)
			if act != tt.expected {
				t.Errorf("expected %d but actual is %d", tt.expected, act)
			}
		})
	}
}

func Test_findFirstMarkerLast4InInput(t *testing.T) {
	lines := readStringFile("day6_input")
	act := findFirstMarker(lines[0], 4)
	exp := 1760
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_findFirstMarkerLast14InSample(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s,%d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			act := findFirstMarker(tt.input, 14)
			if act != tt.expected {
				t.Errorf("expected %d but actual is %d", tt.expected, act)
			}
		})
	}
}

func Test_findFirstMarkerLast14InInput(t *testing.T) {
	lines := readStringFile("day6_input")
	act := findFirstMarker(lines[0], 14)
	exp := 2974
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
