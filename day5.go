package adventofcode2022

import (
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	as "github.com/emirpasic/gods/stacks/arraystack"
	"regexp"
	"strconv"
)

func _parseStackLines(lines []string) []as.Stack {
	var stacks []as.Stack
	for i := len(lines) - 1; i >= 0; i-- {
		l := lines[i]
		if i == len(lines)-1 {
			stacksCount := (len(l) + 1) / 4
			stacks = make([]as.Stack, stacksCount)
			for j := 0; j < stacksCount; j++ {
				stacks[j] = *as.New()
			}
		}
		k := 0
		for j := 0; j < len(l); j += 4 {
			crate := l[j+1 : j+2]
			if crate != " " {
				stacks[k].Push(crate)
			}
			k++
		}
	}
	return stacks
}

type movement struct {
	crates int
	from   int
	to     int
}

func _parseMovementLines(lines []string) []movement {
	var re = regexp.MustCompile(`(?m)^move\s(?P<crates>\d+)\sfrom\s(?P<from>\d+)\sto\s(?P<to>\d+)$`)

	movements := make([]movement, len(lines))
	for i, l := range lines {

		match := re.FindStringSubmatch(l)
		result := make(map[string]string)
		for j, name := range re.SubexpNames() {
			if j != 0 && name != "" {
				result[name] = match[j]
			}
		}

		movements[i] = movement{}
		movements[i].crates, _ = strconv.Atoi(result["crates"])
		movements[i].from, _ = strconv.Atoi(result["from"])
		movements[i].to, _ = strconv.Atoi(result["to"])
	}
	return movements
}

func moveOneByOne(movement movement, stacks []as.Stack) {
	for c := 0; c < movement.crates; c++ {
		crateOut, _ := stacks[movement.from-1].Pop()
		stacks[movement.to-1].Push(crateOut)
	}
}

func moveAllTogether(movement movement, stacks []as.Stack) {
	queue := aq.New()
	for c := 0; c < movement.crates; c++ {
		crateOut, _ := stacks[movement.from-1].Pop()
		queue.Enqueue(crateOut)
	}
	it := queue.Iterator()
	for it.End(); it.Prev(); {
		stacks[movement.to-1].Push(it.Value())
	}
}

func getTopStacks(lines []string, move func(movement movement, stacks []as.Stack)) string {

	// parse stacks and movements
	emptyLine := 0
	for i, l := range lines {
		if l == "" {
			emptyLine = i
			break
		}
	}
	stacks := _parseStackLines(lines[:emptyLine-1])
	movements := _parseMovementLines(lines[emptyLine+1:])

	// run the movements
	for _, m := range movements {
		move(m, stacks)
	}

	// get crates from the top of each stack
	topCrates := ""
	for _, s := range stacks {
		crate, _ := s.Peek()
		topCrates += crate.(string)
	}

	return topCrates
}
