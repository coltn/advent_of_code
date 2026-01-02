package main

/* day06/part01 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day06 "aoc/2025_go/day06/pkg"
)


func main() {

	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	problems, err := day06.LoadFileIntoMemory(reader)
	if err != nil {
		log.Fatal(err)
	}


	answer := day06.HandleProblems(problems)

	fmt.Println(answer)



}
