package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

var numbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// digitAtPosition returns the digit at potision i in string line, from 1 to 9
// This works for both digits 1-9 and strings one-nine
// Returns 0 in case of it not being a digit
func digitAtPosition(i int, line string) int {
	if unicode.IsDigit(rune(line[i])) {
		return int(rune(line[i]) - '0')
	}
	for j, num := range numbers {
		if len(line)-i < len(num) {
			continue
		} else if line[i:i+len(num)] == num {
			return j + 1
		}
	}
	return 0
}
func SolveDay01() {
	f, err := os.Open("inputs/input_day_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		fmt.Println(line)
		i := 0
		j := len(line) - 1
		var digitI, digitJ int
		for {
			digitI = digitAtPosition(i, line)
			if digitI == 0 {
				i++
			}

			digitJ = digitAtPosition(j, line)
			if digitJ == 0 {
				j--
			}

			if digitI != 0 && digitJ != 0 {
				break
			}
		}
		sum += digitI*10 + digitJ
	}
	fmt.Printf("Result: %d\n", sum)
}
