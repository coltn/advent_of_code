package main

/* day03/part01 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day03 "aoc/2025_go/day03/pkg"
)


func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	joltages, err := day03.FindJoltages(scanner)
	if err != nil {
		log.Fatal("Failed to FindJoltages(): %v", err)
	}

	answer := 0
	for _, joltage := range joltages {
		answer += joltage
	}


	fmt.Println(answer)
}
