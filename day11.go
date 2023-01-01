package adventofcode2022

import (
	aq "github.com/emirpasic/gods/queues/arrayqueue"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id            int
	startingItems *aq.Queue
	operation     operation
	test          test
}

type operation struct {
	operand1 string
	operand2 string
	operator string
}

type test struct {
	divisibleBy          int
	onTrueThrowToMonkey  int
	onFalseThrowToMonkey int
}

func calcMonkeyBusiness(lines []string, rounds int, divideWorryLevel int) int {

	useLCM := false
	lcm := 1
	if divideWorryLevel == 0 {
		useLCM = true
	}

	var monkeys []monkey
	for l := 0; l < len(lines); l += 7 {
		entry := strings.Join(lines[l:l+6], "\n")

		var re = regexp.MustCompile(`(?m)^Monkey\s(?P<id>\d+):\s{3}Starting\sitems:(?P<startingItems>\s(\d+)(,\s*\d+)*)\s{3}Operation:\s(?P<operationResult>\w+)\s=\s(?P<operand1>\w+)\s(?P<operator>[+\-*\/])\s(?P<operand2>\w+)\s{3}Test:\sdivisible\sby\s(?P<divisibleBy>\d+)\s{5}If\strue:\sthrow\sto\smonkey\s(?P<onTrueThrowToMonkey>\d+)\s{5}If\sfalse:\sthrow\sto\smonkey\s(?P<onFalseThrowToMonkey>\d+)$`)

		monkey := monkey{}

		match := re.FindStringSubmatch(entry)
		result := make(map[string]string)
		for j, name := range re.SubexpNames() {
			if j != 0 && name != "" {
				result[name] = match[j]
			}
		}

		monkey.id, _ = strconv.Atoi(result["id"])

		startItems := strings.Split(result["startingItems"], ",")
		monkey.startingItems = aq.New()
		for _, s := range startItems {
			i, _ := strconv.Atoi(strings.TrimSpace(s))
			monkey.startingItems.Enqueue(i)
		}

		monkey.operation = operation{
			operand1: result["operand1"],
			operator: result["operator"],
			operand2: result["operand2"],
		}

		monkey.test = test{}
		monkey.test.divisibleBy, _ = strconv.Atoi(result["divisibleBy"])
		monkey.test.onTrueThrowToMonkey, _ = strconv.Atoi(result["onTrueThrowToMonkey"])
		monkey.test.onFalseThrowToMonkey, _ = strconv.Atoi(result["onFalseThrowToMonkey"])

		monkeys = append(monkeys, monkey)

		if useLCM {
			lcm *= monkey.test.divisibleBy
		}
	}

	inspections := make([]int, len(monkeys))

	for r := 1; r <= rounds; r++ {
		for _, m := range monkeys {
			values := m.startingItems.Values()
			for _, o := range values {

				i := o.(int)

				worryLevel := 0

				// monkey inspects the item

				var operand1, operand2 int
				if m.operation.operand1 == "old" {
					operand1 = i
				} else {
					operand1, _ = strconv.Atoi(m.operation.operand1)
				}
				if m.operation.operand2 == "old" {
					operand2 = i
				} else {
					operand2, _ = strconv.Atoi(m.operation.operand2)
				}

				switch m.operation.operator {
				case "+":
					worryLevel = operand1 + operand2
					break
				case "*":
					worryLevel = operand1 * operand2
					break
				}

				// monkey gets bored with the item
				if useLCM {
					worryLevel %= lcm
				} else {
					worryLevel /= divideWorryLevel
				}

				// test and throw to another monkey
				m.startingItems.Dequeue()
				if worryLevel%m.test.divisibleBy == 0 {
					monkeys[m.test.onTrueThrowToMonkey].startingItems.Enqueue(worryLevel)
				} else {
					monkeys[m.test.onFalseThrowToMonkey].startingItems.Enqueue(worryLevel)
				}

				inspections[m.id]++
			} // startingItems
		} // monkeys

	} // rounds

	sort.Slice(inspections, func(i int, j int) bool { return inspections[i] > inspections[j] })

	return inspections[0] * inspections[1]
}
