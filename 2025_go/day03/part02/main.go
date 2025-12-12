package main

/* day03/part02 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day03 "aoc/2025_go/day03/pkg"
)

const AMOUNT int = 12

func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	joltages, err := day03.FindAmountOfJoltages(scanner, AMOUNT)
	if err != nil {
		log.Fatalf("Failed to FindJoltages(): %v\n", err)
	}

	answer := 0
	for _, joltage := range joltages {
		answer += joltage
	}


	fmt.Println(answer)
}
