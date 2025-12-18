
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

		// debug
		fmt.Printf("%v", string(rows.Current))

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

func checkRows(rows Rows, position int) bool {
	roll_count := 0
	end := len(rows.Current)
	
	if position != 0 {
		roll_count += checkPosition(rows.Prev, position - 1)
		roll_count += checkPosition(rows.Current, position - 1)
		roll_count += checkPosition(rows.Next, position - 1)
	} else if position != end {
		roll_count += checkPosition(rows.Prev, position + 1)
		roll_count += checkPosition(rows.Current, position + 1)
		roll_count += checkPosition(rows.Next, position + 1)
	}

	roll_count += checkPosition(rows.Prev, position)
	roll_count += checkPosition(rows.Next, position)

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
