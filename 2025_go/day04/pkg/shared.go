
package day04

import (
	"bufio"
	"bytes"
	"fmt"
)

type Rows struct {
	Prev []rune
	Current []rune
	Next []rune
}

const TARGET_COUNT int = 4
const PAPER_ROLL rune = '@'
const DELIM byte = '\n'



func FindRollsInFile(input *bufio.Reader) (int, error) {

	current, err := input.ReadBytes(DELIM)
	if err != nil {
		return 0, err
	}
	next, err := input.ReadBytes(DELIM)
	if err != nil {
		return 0, err
	}



	rows := Rows {
		Prev: nil,
		Current: bytes.Runes(current),
		Next: bytes.Runes(next),
	}

	total_count := 0

	for {
		for i, r := range rows.Current {
			if r == PAPER_ROLL {
				if checkRows(rows, i) {
					total_count ++
				}
			}

		}

		if rows.Next == nil {
			break
		}

		rows.Prev = rows.Current
		rows.Current = rows.Next
		next_bytes, err := input.ReadBytes(DELIM)
		if err != nil {
			rows.Next = nil
		} else {
			rows.Next = bytes.Runes(next_bytes)
		}
	}

	return total_count, nil
}

func checkRows(rows Rows, col_position int) bool {
	roll_count := 0
	row_end := len(rows.Current)
	
	if col_position != 0 {
		roll_count += checkPosition(rows.Prev, col_position - 1)
		roll_count += checkPosition(rows.Current, col_position - 1)
		roll_count += checkPosition(rows.Next, col_position - 1)
	} 

	if col_position != row_end - 1 {
		roll_count += checkPosition(rows.Prev, col_position + 1)
		roll_count += checkPosition(rows.Current, col_position + 1)
		roll_count += checkPosition(rows.Next, col_position + 1)
	}

	roll_count += checkPosition(rows.Prev, col_position)
	roll_count += checkPosition(rows.Next, col_position)

	return roll_count < TARGET_COUNT
}

func checkPosition(row []rune, position int) int {
	if row == nil {
		return 0
	}

	if row[position] == PAPER_ROLL {
		return 1
	}

	return 0
}


/* 
 * below used for part02
 * tried using a different method for finding nearby paper rolls.
 * would have liked to find a way to do this without loading the entire file into memory.
 * but i just wanted to get the problem done lol
 */

func LoadFileIntoMemory(input *bufio.Scanner) ([][]rune, error) {
	var grid [][]rune
	for input.Scan() {
		grid = append(grid, []rune(input.Text()))
	}

	return grid, input.Err()
}

func FindRollsInMemory(grid [][]rune) int {
	n_rows := len(grid)
	total := 0
	moveable := -1
	for moveable != 0 {
		moveable = 0
		for r_idx, row := range grid {
			// debug: print current row
			fmt.Printf("%v\n", string(row))
			for c_idx, r := range row {
				nearby_rolls := 0
				n_col := len(row)
				if r == PAPER_ROLL {
					for r_offs := -1; r_offs < 2; r_offs ++ {
						r_io := r_idx + r_offs
						if r_io < 0 || r_io > n_rows - 1{
							continue
						}
						for c_offs := -1; c_offs < 2; c_offs ++ {
							if r_offs == 0 && c_offs == 0 {
								continue
							}

							c_io := c_idx + c_offs

							if c_io < 0 || c_io > n_col - 1 {
								continue
							}

							if grid[r_io][c_io] == PAPER_ROLL {
								nearby_rolls ++
							}
						}
					}

					if nearby_rolls < TARGET_COUNT { 
						moveable ++
						r = 'x'
					}
				}


				/* 
 				 * mutating grid in place asynchronously
				 * technically wrong--AoC says to mutate the grid in one synchronous step
 				 * but i can't be bothered right now, and this gets the result we need.
				 * removals are monotonic and accelerate convergence
 				 */
				grid[r_idx][c_idx] = r
			}
		}
		total += moveable

		// debug: print new line to seperate each grid as it changes
		fmt.Printf("\n")

	}

	return total
}
