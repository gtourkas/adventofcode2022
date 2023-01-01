package adventofcode2022

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"sort"
)

func findShortestPath(lines []string, startFromLowestElevation bool) int {

	pointsWithLowestElevation := make([]string, 0)

	// build the hill graphs; non-directional and weighted on the distance

	getKey := func(x int, y int) string {
		return fmt.Sprintf("%d-%d", x, y)
	}

	g := graph.New(graph.StringHash, graph.Directed())

	dstKey := ""

	normalizeHeight := func(c uint8) uint8 {
		if c == 'S' {
			return 'a'
		}
		if c == 'E' {
			return 'z'
		}
		return c
	}

	maxX := len(lines[0])
	maxY := len(lines)
	for y := 0; y < maxY; y++ {
		l := lines[y]
		for x := 0; x < maxX; x++ {

			key := getKey(x, y)

			g.AddVertex(key)

			height := int(normalizeHeight(l[x]))

			if l[x] == 'S' || (startFromLowestElevation && l[x] == 'a') {
				pointsWithLowestElevation = append(pointsWithLowestElevation, key)
			}
			if l[x] == 'E' {
				dstKey = key
			}

			if x > 0 {
				leftHeight := int(normalizeHeight(l[x-1]))
				// left to right
				diff := height - leftHeight
				if diff <= 1 {
					g.AddEdge(getKey(x-1, y), key, graph.EdgeWeight(1))
				}
				// right to left
				diff = leftHeight - height
				if diff <= 1 {
					g.AddEdge(key, getKey(x-1, y), graph.EdgeWeight(1))
				}
			}

			if y > 0 {
				topHeight := int(normalizeHeight(lines[y-1][x]))
				// top to bottom
				diff := height - topHeight
				if diff <= 1 {
					g.AddEdge(getKey(x, y-1), key, graph.EdgeWeight(1))
				}
				// bottom to top
				diff = topHeight - height
				if diff <= 1 {
					g.AddEdge(key, getKey(x, y-1), graph.EdgeWeight(1))
				}
			}

		}
	}

	// find the shortest path
	pathsLengths := make([]int, 0)
	for _, srcKey := range pointsWithLowestElevation {
		path, err := graph.ShortestPath(g, srcKey, dstKey)
		if err == nil {
			pathsLengths = append(pathsLengths, len(path)-1)
		}
	}

	sort.Slice(pathsLengths, func(i int, j int) bool { return pathsLengths[i] < pathsLengths[j] })

	return pathsLengths[0]
}
