package main

/* day04/part01 */


import (
	"bufio"
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

	rows := day04.InitRows(scanner)
	




}
