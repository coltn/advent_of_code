package main

/* day04/part02 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day04 "aoc/2025_go/day04/pkg"
)


func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid, err := day04.LoadFileIntoMemory(scanner)
	if err != nil {
		log.Fatal(err)
	}

	answer := day04.FindRollsInMemory(grid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
