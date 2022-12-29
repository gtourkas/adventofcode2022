package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

func _abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func _key(x int, y int) string {
	return fmt.Sprintf("%d#%d", x, y)
}

type _point struct {
	x int
	y int
}

func _visualize(rope []_point, size int) {
	printed := make(map[string]bool)
	for y := -size; y < size; y++ {
		for x := -size; x < size; x++ {
			key := _key(x, y)
			for k := 0; k < len(rope); k++ {
				if rope[k].x == x && rope[k].y == y {
					if _, ok := printed[key]; !ok {
						fmt.Print(k)
						printed[key] = true
					}
				}
			}
			if _, ok := printed[key]; !ok {
				fmt.Print(".")
				printed[key] = true
			}
		}
		fmt.Println()
	}
}

func _calcNewCoordinate(diff int) int {
	// movement is at most by one regardless of the difference at any direction
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}
	return 0
}

func countTailVisits(lines []string, knotLength int) int {

	rope := make([]_point, knotLength)

	tailVisits := make(map[string]bool)
	tailVisits[_key(0, 0)] = true

	for _, l := range lines {
		parts := strings.Split(l, " ")

		dir := parts[0]
		positions, _ := strconv.Atoi(parts[1])

		for p := 1; p <= positions; p++ {

			// move head
			switch dir {
			case "R":
				rope[0].x++
				break
			case "L":
				rope[0].x--
				break
			case "U":
				rope[0].y--
				break
			case "D":
				rope[0].y++
				break
			}

			for k := 1; k < len(rope); k++ {

				// check if adjacent (manhattan distance)
				diffX := rope[k-1].x - rope[k].x
				diffY := rope[k-1].y - rope[k].y
				isAdjacent := _abs(diffX) <= 1 &&
					_abs(diffY) <= 1

				if !isAdjacent {
					rope[k].x += _calcNewCoordinate(diffX)
					rope[k].y += _calcNewCoordinate(diffY)
				}
			}
			tailVisits[_key(rope[len(rope)-1].x, rope[len(rope)-1].y)] = true

		}
	}

	return len(tailVisits)
}
