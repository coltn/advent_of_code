package main

/* day04/part01 */


import (
	"bufio"
	"fmt"
	"log"
	"os"

	 day04 "aoc/2025_go/day04/pkg"
)


func main() {

	f, err := os.Open("../small_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	fmt.Println(day04.FindRolls(reader))

	




}
