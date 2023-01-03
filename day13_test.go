package adventofcode2022

import "testing"

func Test__getListItems(t *testing.T) {
	testCases := map[string][]string{
		"[1,1,3,1,1]":                           {"1", "1", "3", "1", "1"},
		"[[1],[2,3,4]]":                         {"[1]", "[2,3,4]"},
		"[[1],4]":                               {"[1]", "4"},
		"[1,[2,[3,[4,[5,6,7]]]],8,9]":           {"1", "[2,[3,[4,[5,6,7]]]]", "8", "9"},
		"[]":                                    {},
		"[[[[8,7,8,5,4],6,[4,6]]],[6,[[]]],[]]": {"[[[8,7,8,5,4],6,[4,6]]]", "[6,[[]]]", "[]"},
		"[[[8,7,8,5,4],6,[4,6]]]":               {"[[8,7,8,5,4],6,[4,6]]"},
		"[[8,7,8,5,4],6,[4,6]]]":                {"[8,7,8,5,4]", "6", "[4,6]"},
	}

	for inp, exp := range testCases {
		act := _getListItems(inp)
		if len(exp) != len(act) {
			t.Errorf("expected %v but actual is %v", exp, act)
			return
		}
		for i := range act {
			if act[i] != exp[i] {
				t.Errorf("expected %s but actual is %s at position %d", exp[i], act[i], i)
			}
		}
	}
}

func Test_calcSumOfIndicesForOrderedPairsInSample(t *testing.T) {
	lines := readStringFile("day13_sample")
	act := calcSumOfIndicesForOrderedPairs(lines)
	exp := 13
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_calcSumOfIndicesForOrderedPairsInInput(t *testing.T) {
	lines := readStringFile("day13_input")
	act := calcSumOfIndicesForOrderedPairs(lines)
	exp := 6240
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getDecoderKeyInSample(t *testing.T) {
	lines := readStringFile("day13_sample")
	act := getDecoderKey(lines)
	exp := 140
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}

func Test_getDecoderKeyInInput(t *testing.T) {
	lines := readStringFile("day13_input")
	act := getDecoderKey(lines)
	exp := 23142
	if exp != act {
		t.Errorf("expected %d but actual is %d", exp, act)
	}
}
