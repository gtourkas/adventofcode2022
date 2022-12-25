package adventofcode2022

import (
	"fmt"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"math"
	"strconv"
	"strings"
)

func _calcPathsSizes(lines []string) map[string]int {
	const (
		CDPrefix = "$ cd"
		LSPrefix = "$ ls"
		DirCmd   = "dir"
	)

	pathStack := lls.New()
	currentPath := ""
	currentPathDepth := 0
	pathsSizes := make(map[string]int)
	for _, l := range lines {
		if strings.HasPrefix(l, CDPrefix) {
			dirName := strings.TrimSpace(strings.Replace(l, CDPrefix, "", 1))
			if dirName == ".." {
				pathStack.Pop()
				o, _ := pathStack.Peek()
				currentPath = o.(string)
				currentPathDepth--
			} else {
				if dirName == "/" {
					dirName = ""
					currentPath = "/"
				} else {
					if currentPath != "/" {
						currentPath = fmt.Sprintf("%s/%s", currentPath, dirName)
					} else {
						currentPath = fmt.Sprintf("/%s", dirName)
					}
					currentPathDepth++
				}
				pathStack.Push(currentPath)
			}
		} else if !strings.HasPrefix(l, DirCmd) && !strings.HasPrefix(l, LSPrefix) {
			rowParts := strings.Split(l, " ")
			rowSize, _ := strconv.Atoi(rowParts[0])
			_, ok := pathsSizes[currentPath]
			if ok {
				pathsSizes[currentPath] += rowSize
			} else {
				pathsSizes[currentPath] = rowSize
			}
			// add rawSize to parent paths
			if currentPathDepth > 0 {
				it := pathStack.Iterator()
				for it.Next() {
					o := it.Value()
					path := o.(string)
					if path != currentPath {
						pathsSizes[path] += rowSize
					}
				}
			}
		}
	}

	return pathsSizes
}

func sumSizeOfDirsUnder(lines []string, maxSize int) int {

	pathsSizes := _calcPathsSizes(lines)

	sum := 0
	for path := range pathsSizes {
		pathSize := pathsSizes[path]
		if pathSize <= maxSize {
			sum += pathSize
		}
	}

	return sum
}

func getSizeOfDirToDelete(lines []string, fsSize int, unusedSpaceNeeded int) int {

	pathsSizes := _calcPathsSizes(lines)

	unusedSpace := fsSize - pathsSizes["/"]
	sizeOfDirToDelete := math.MaxInt32
	for path := range pathsSizes {
		if path == "/" {
			continue
		}
		pathSize := pathsSizes[path]
		if unusedSpace+pathSize >= unusedSpaceNeeded {
			if pathSize < sizeOfDirToDelete {
				sizeOfDirToDelete = pathSize
			}
		}
	}

	return sizeOfDirToDelete
}
