package adventofcode2022

import (
	"strconv"
	"strings"
)

func calcUnitsOfSandsAtRest(lines []string, hasFloor bool) int {

	caveSize := 1000
	prodX, prodY := 500, 0

	parsePoint := func(p string) (int, int) {
		xy := strings.Split(p, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		return x, y
	}

	minInt := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}

	maxInt := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// build the cave
	cave := make([][]rune, caveSize)
	for x := 0; x < len(cave); x++ {
		cave[x] = make([]rune, caveSize)
	}

	maxY := 0
	for _, l := range lines {
		pairs := strings.Split(l, " -> ")

		fromX, fromY := parsePoint(pairs[0])

		for i := 1; i < len(pairs); i++ {

			tillX, tillY := parsePoint(pairs[i])

			if maxInt(fromY, tillY) > maxY {
				maxY = maxInt(fromY, tillY)
			}

			for x := minInt(fromX, tillX); x <= maxInt(fromX, tillX); x++ {
				for y := minInt(fromY, tillY); y <= maxInt(fromY, tillY); y++ {
					cave[x][y] = '#'
				}
			}
			fromX, fromY = tillX, tillY
		}
	}

	if hasFloor {
		for x := 0; x < len(cave); x++ {
			cave[x][maxY+2] = '#'
		}
	}

	// start sand production
	atRest := 0
	produce := true
	for produce {
		x, y := prodX, prodY
		for {
			// down
			if cave[x][y+1] != 0 {
				// left
				if cave[x-1][y+1] != 0 {
					// right
					if cave[x+1][y+1] != 0 {
						// can't move any further - put at rest
						// if putting at rest stops at production point, stop producing
						cave[x][y] = 'Ο'
						atRest++
						if x == prodX && y == prodY {
							produce = false
						}
						break
					} else {
						x++
						y++
					}
				} else {
					x--
					y++
				}
			} else {
				y++
				if y == caveSize-1 {
					produce = false
					break
				}
			}
		}
	}

	return atRest
}
