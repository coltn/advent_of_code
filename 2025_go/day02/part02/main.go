package main

/* day02/part02 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day02 "aoc/2025_go/day02/pkg"
)


func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	ranges, err := day02.ParseInputToRanges(reader)
	if err != nil {
		log.Fatalf("Could not parse file into ranges: %v", err)
	}



	invalid := day02.FindInvalidIdRepeating(ranges)

	answer := 0
	for _, id := range invalid {
		answer += id
	}

	fmt.Println(answer)
}
