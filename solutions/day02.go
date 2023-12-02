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

	totalPower := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		gameNameAndData := strings.Split(line, ":")

		var minColorCounts = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		power := 0

		games := strings.Split(gameNameAndData[1], ";")
		for _, game := range games {
			fmt.Println(game)
			colorComponents := strings.Split(game, ",")
			for _, component := range colorComponents {
				countAndColor := strings.Split(component, " ")
				count, _ := strconv.Atoi(countAndColor[1])
				if count > minColorCounts[countAndColor[2]] {
					if countAndColor[2] == "red" {
						fmt.Printf("%d > %d\n", count, minColorCounts[countAndColor[2]])
					}
					minColorCounts[countAndColor[2]] = count
				}
			}
		}
		fmt.Println(minColorCounts)
		power = minColorCounts["red"] * minColorCounts["green"] * minColorCounts["blue"]
		totalPower += power
	}
	fmt.Println(totalPower)
}
