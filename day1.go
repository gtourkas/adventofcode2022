package adventofcode2022

import (
	"sort"
	"strconv"
)

func getMaxCalories(lines []string) int {
	maxCalories := 0

	curCalories := 0
	for _, l := range lines {
		if l == "" {
			if curCalories > maxCalories {
				maxCalories = curCalories
			}
			curCalories = 0
		} else {
			lineCalories, _ := strconv.Atoi(l)
			curCalories += lineCalories
		}
	}

	return maxCalories

}

func getSumOfTopCalories(lines []string, top int) int {
	var calories []int

	curCalories := 0
	i := 0
	for _, l := range lines {
		if l == "" {
			calories = append(calories, curCalories)
			curCalories = 0
			i++
		} else {
			lineCalories, _ := strconv.Atoi(l)
			curCalories += lineCalories
		}
	}
	calories = append(calories, curCalories)

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	sum := 0
	for i = 0; i < top; i++ {
		sum += calories[i]
	}
	return sum
}
