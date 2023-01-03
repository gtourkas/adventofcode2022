package adventofcode2022

import (
	"fmt"
	"sort"
	"strconv"
)

func calcSumOfIndicesForOrderedPairs(lines []string) int {

	sum := 0

	for i := 0; i < len(lines); i += 3 {
		pair := lines[i : i+2]

		r := _areInOrder(pair[0], pair[1])

		if r == 1 {
			sum += (i / 3) + 1
		}

	}

	return sum
}

func getDecoderKey(lines []string) int {

	dividerPackets := []string{"[[2]]", "[[6]]"}

	packetLines := make([]string, 0)
	for _, l := range lines {
		if l != "" {
			packetLines = append(packetLines, l)
		}
	}

	packetLines = append(packetLines, dividerPackets...)

	sort.Slice(packetLines, func(i, j int) bool {
		inOrder := _areInOrder(packetLines[i], packetLines[j])
		return inOrder == 1
	})

	decoderKey := 1

	for i, p := range packetLines {
		for _, dp := range dividerPackets {
			if p == dp {
				decoderKey *= (i + 1)
			}
		}
	}

	return decoderKey
}

func _isList(o string) bool {
	return len(o) >= 2 && o[0] == '[' && o[len(o)-1] == ']'
}

func _getListItems(o string) []string {

	items := make([]string, 0)

	level := 0
	itemStart := 0
	for i := 0; i < len(o); i++ {
		if o[i] == '[' {
			level++
			if level == 1 {
				itemStart = i + 1
			}
		} else if o[i] == ']' {
			level--
			if level == 0 {
				item := o[itemStart:i]
				if item != "" {
					items = append(items, item)
				}
				itemStart = 0
			}
		} else if o[i] == ',' {
			if level == 1 {
				item := o[itemStart:i]
				if item != "" {
					items = append(items, item)
				}
				itemStart = i + 1
			}
		}
	}

	return items
}

func _areInOrder(left string, right string) int {

	inOrder := 0 //unknown

	lIsList := _isList(left)
	rIsList := _isList(right)

	// both integers
	if !lIsList && !rIsList {
		li, _ := strconv.Atoi(left)
		ri, _ := strconv.Atoi(right)
		if li < ri {
			return 1 // in order
		} else if li > ri {
			return -1 // not in order
		} else {
			return 0 // unknown
		}
	} else if lIsList && rIsList {
		leftItems := _getListItems(left)
		rightItems := _getListItems(right)
		for i := 0; i < len(leftItems) && i < len(rightItems); i++ {
			l := leftItems[i]
			r := rightItems[i]
			inOrder = _areInOrder(l, r)
			if inOrder != 0 {
				break
			}
		}
		if inOrder == 0 {
			if len(leftItems) < len(rightItems) {
				inOrder = 1 // in order
			}
			if len(leftItems) > len(rightItems) {
				inOrder = -1 // not in order
			}
		}
	} else {
		l := left
		r := right
		if !lIsList {
			l = fmt.Sprintf("[%s]", left)
		}
		if !rIsList {
			r = fmt.Sprintf("[%s]", right)
		}
		inOrder = _areInOrder(l, r)
	}

	return inOrder
}
