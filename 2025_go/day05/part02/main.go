package main

/* day05/part01 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day05 "aoc/2025_go/day05/pkg"
)


func main() {

	f, err := os.Open("../small_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	ranges, _, err := day05.LoadFileIntoMemory(reader)
	if err != nil {
		log.Fatal(err)
	}

	answer := day05.CountValidIDsInRanges(ranges)

	fmt.Println(answer)


}
