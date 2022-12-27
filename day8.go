package adventofcode2022

import "fmt"

func _countIfVisible(grid [][]uint8,
	x int,
	y int,
	tallest *uint8,
	visibleMap map[string]bool) {
	cur := grid[x][y]
	if cur > *tallest {
		key := fmt.Sprintf("%d-%d", x, y)
		if _, ok := visibleMap[key]; !ok {
			visibleMap[key] = true
		}
		*tallest = cur
	}
}

func _buildGrid(lines []string) ([][]uint8, int, int) {
	maxX := len(lines[0])
	maxY := len(lines)
	grid := make([][]uint8, maxX)
	for x := 0; x < maxX; x++ {
		grid[x] = make([]uint8, maxY)
		for y := 0; y < maxY; y++ {
			grid[x][y] = lines[y][x]
		}
	}
	return grid, maxX, maxY
}

func countVisibleTrees(lines []string) int {

	grid, maxX, maxY := _buildGrid(lines)

	var tallest uint8 = 0
	visibleMap := make(map[string]bool)

	// top to bottom
	for x := 0; x < maxX; x++ {
		tallest = 0
		for y := 0; y < maxY; y++ {
			_countIfVisible(grid, x, y, &tallest, visibleMap)
		}
	}
	// bottom to top
	for x := 0; x < maxX; x++ {
		tallest = 0
		for y := maxY - 1; y >= 0; y-- {
			_countIfVisible(grid, x, y, &tallest, visibleMap)
		}
	}
	// left to right
	for y := 0; y < maxY; y++ {
		tallest = 0
		for x := 0; x < maxX; x++ {
			_countIfVisible(grid, x, y, &tallest, visibleMap)
		}
	}
	// right to left
	for y := 0; y < maxY; y++ {
		tallest = 0
		for x := maxX - 1; x >= 0; x-- {
			_countIfVisible(grid, x, y, &tallest, visibleMap)
		}
	}

	return len(visibleMap)
}

func calcHighestScenicScore(lines []string) int {

	grid, maxX, maxY := _buildGrid(lines)

	highestScenicScore := 1
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {

			cur := grid[x][y]

			// up (same x, decreasing y)
			upScenicScore := 0
			for i := y - 1; i >= 0; i-- {
				next := grid[x][i]
				if cur > next {
					upScenicScore++
				} else {
					upScenicScore++
					break
				}
			}

			// down (same x, increasing y)
			downScenicScore := 0
			for i := y + 1; i < maxY; i++ {
				next := grid[x][i]
				if cur > next {
					downScenicScore++
				} else {
					downScenicScore++
					break
				}
			}

			// left (same y, decreasing x)
			leftScenicScore := 0
			for i := x - 1; i >= 0; i-- {
				next := grid[i][y]
				if cur > next {
					leftScenicScore++
				} else {
					leftScenicScore++
					break
				}
			}

			// right (same y, increasing x)
			rightScenicScore := 0
			for i := x + 1; i < maxX; i++ {
				next := grid[i][y]
				if cur > next {
					rightScenicScore++
				} else {
					rightScenicScore++
					break
				}
			}

			totalScenicScore := upScenicScore * downScenicScore * leftScenicScore * rightScenicScore
			if totalScenicScore > highestScenicScore {
				highestScenicScore = totalScenicScore
			}
		}
	}

	return highestScenicScore
}
