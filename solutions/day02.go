package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxColorCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func SolveDay02() {
	f, err := os.Open("inputs/input_day_02.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		gameNameAndData := strings.Split(line, ":")

		gameID, _ := strconv.Atoi(strings.Split(gameNameAndData[0], " ")[1])

		games := strings.Split(gameNameAndData[1], ";")
		possible := true
		for _, game := range games {
			colorComponents := strings.Split(game, ",")
			for _, component := range colorComponents {
				countAndColor := strings.Split(component, " ")
				count, _ := strconv.Atoi(countAndColor[1])
				if count > maxColorCounts[countAndColor[2]] {
					possible = false
					break
				}
			}
		}
		if possible {
			total += gameID
		}
	}
	fmt.Println(total)
}
