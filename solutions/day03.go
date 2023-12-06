package solutions

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func SolveDay03() {
	content, err := os.ReadFile("inputs/input_day_03.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	lines := strings.Split(string(content), "\n")
	for i := 0; i < len(lines); {
		fmt.Printf("line %q\n", lines[i])
		for j := 0; j < len(lines[i]); {
			if unicode.IsDigit(rune(lines[i][j])) {
				startNumberIndex := j
				j++
				for ; j < len(lines[i]) && unicode.IsDigit(rune(lines[i][j])); j++ {
				}

				total += getValue(lines, i, startNumberIndex, j)

				continue
			}
			j++
		}
		i++
	}
	fmt.Printf("Total: %d\n", total)
}

func checkLine(line string, start int, end int) bool {
	if start > len(line)-1 {
		return false
	}
	for j := start; j < end; j++ {
		if line[j] != '.' {
			return true
		}
	}
	return false
}

func getValue(lines []string, lineNumber, startNumberIndex, endNumberIndex int) int {
	number, _ := strconv.Atoi(lines[lineNumber][startNumberIndex:endNumberIndex])
	// fmt.Printf("number %d (%d:%d)\n", number, startNumberIndex, endNumberIndex)

	// Check upper line
	if lineNumber > 0 {
		if checkLine(lines[lineNumber-1], startNumberIndex, endNumberIndex) {
			// fmt.Println("returns in 1")
			return number
		}
	}
	// Check lower line
	if lineNumber < len(lines)-1 {
		if checkLine(lines[lineNumber+1], startNumberIndex, endNumberIndex) {
			// fmt.Println("returns in 2")
			return number
		}
	}
	// Check left side
	if startNumberIndex > 0 {
		if lineNumber > 0 {
			if lines[lineNumber-1][startNumberIndex-1] != '.' {
				// fmt.Println("returns in 3")
				return number
			}
		}
		if lineNumber < len(lines)-2 {
			fmt.Printf("%d %d\n", startNumberIndex, len(lines[lineNumber+1]))
			if lines[lineNumber+1][startNumberIndex-1] != '.' {
				// fmt.Println("returns in 4")
				return number
			}
		}
		if lines[lineNumber][startNumberIndex-1] != '.' {
			// fmt.Println("returns in 5")
			return number
		}
	}
	// Check right side
	if endNumberIndex < len(lines[lineNumber])-1 {
		if lineNumber > 0 {
			if lines[lineNumber-1][endNumberIndex] != '.' {
				// fmt.Println("returns in 6")
				return number
			}
		}
		if lineNumber < len(lines)-2 {
			if lines[lineNumber+1][endNumberIndex] != '.' {
				// fmt.Println("returns in 7")
				return number
			}
		}
		if lines[lineNumber][endNumberIndex] != '.' {
			// fmt.Println("returns in 8")
			return number
		}
	}
	// fmt.Printf("Invald %q\n", lines[lineNumber][startNumberIndex:endNumberIndex])
	return 0
}
