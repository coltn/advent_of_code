package main


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


	for _, id_range := range ranges {
		fmt.Println(id_range)
	}
}
