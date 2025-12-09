
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Uint100_t uint8
const mod = 100

func (u Uint100_t) normalize() Uint100_t {
	return Uint100_t(uint8(u) % mod)
}

func (u Uint100_t) Add(x int) Uint100_t {
	n := int(u) + x
	/* remove negative values when using % */
	n = (n % mod + mod) % mod
	return Uint100_t(n)
}

func (u Uint100_t) Sub(x int) Uint100_t {
	return u.Add(-x)
}


func main() {
	/*
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	*/
	type move struct {
		direction rune
		amount int
	}
	var movements []move

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		dir := rune(line[0])
		amt, err := strconv.Atoi(line[1:])
		if err != nil {
			error := fmt.Errorf("Error parsing 'amt': %w", err)
			fmt.Println(error)
			continue
		}

		movements = append(movements, move {
			direction: dir,
			amount: amt,
		})
	}


	var position Uint100_t = 50
	counter := 0



	for _, line := range movements {
		i := 0
		if line.direction == 'L' {
			for ; i < line.amount; i++ {
				position = position.Add(1)
				if position == 0 {
					counter ++
				}
			}
		} else if line.direction == 'R' {
			for ; i < line.amount; i ++ { 
				position = position.Sub(1)
				if position == 0 {
					counter ++
				}
			}
		} else {
			fmt.Printf("Error, unknown direction: %v\n", line.direction)
		}


	}

	println(counter)
}
