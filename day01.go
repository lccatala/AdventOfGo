package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func solveDay01() {
	f, err := os.Open("input_day_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		j := len(line) - 1
		var charI rune
		var charJ rune
		for {
			charI = rune(line[i])
			isDigitI := unicode.IsDigit(charI)
			if !isDigitI {
				i++
			}

			charJ = rune(line[j])
			isDigitJ := unicode.IsDigit(charJ)
			if !isDigitJ {
				j--
			}

			if isDigitI && isDigitJ {
				break
			}
		}
		sum += int(charI-'0')*10 + int(charJ-'0')
	}
	fmt.Printf("Result: %d\n", sum)
}

func float(r rune) {
	panic("unimplemented")
}
