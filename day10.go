package adventofcode2022

import (
	"fmt"
	"strconv"
	"strings"
)

func calcSumOfSignalStrengths(lines []string, cycles []int) int {
	sumOfSignalStrengths := 0
	prevCycle := 0
	cycle := 0
	cyclesIndex := 0
	registry := 1
	for _, l := range lines {
		var cmd, param string
		cmd = l
		parts := strings.Split(l, " ")
		if len(parts) == 2 {
			cmd, param = parts[0], parts[1]
		}
		prevCycle = cycle
		if cmd == "noop" {
			cycle++
		}
		if cmd == "addx" {
			cycle += 2
		}

		if prevCycle < cycles[cyclesIndex] && cycle >= cycles[cyclesIndex] {
			sumOfSignalStrengths += cycles[cyclesIndex] * registry
			cyclesIndex++
		}

		if cmd == "addx" {
			x, _ := strconv.Atoi(param)
			registry += x
		}

		if cyclesIndex == len(cycles) {
			break
		}

	}
	return sumOfSignalStrengths
}

func printCRT(lines []string) {
	cycle := 0
	prevCycle := 0
	registry := 1
	for _, l := range lines {
		var cmd, param string
		cmd = l
		parts := strings.Split(l, " ")
		if len(parts) == 2 {
			cmd, param = parts[0], parts[1]
		}
		prevCycle = cycle
		if cmd == "noop" {
			cycle++
		}
		if cmd == "addx" {
			cycle += 2
		}

		for c := prevCycle + 1; c <= cycle; c++ {

			lc := c % 40
			if lc >= registry && lc <= registry+2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			if lc == 0 {
				fmt.Println()
			}
		}

		if cmd == "addx" {
			x, _ := strconv.Atoi(param)
			registry += x
		}

	}
}
