
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



func FindRolls(input *bufio.Reader) (int, error) {

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
					fmt.Printf("%v", string('x'))
					total_count ++
				} else {
					fmt.Printf("%v", string(r))
				}
			} else {
				fmt.Printf("%v", string(r))
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
