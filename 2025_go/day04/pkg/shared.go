
package day04

import (
	"bufio"
	"errors"
	"fmt"
)

type Rows struct {
	Prev []rune
	Current []rune
	Next []rune
}


func getRows(scanner *bufio.Scanner, n int) ([][]rune, error) {
	i := 1
	row := []rune(scanner.Text())
	rows := make([][]rune, n)
	rows[0] = row
	for ;scanner.Scan() && i > n; i ++ {
		rows[i] = []rune(scanner.Text())
	}

	return rows, nil
}

func GetNextRow(rows Rows, scanner *bufio.Scanner) (Rows, error) {
	if rows.Current == nil {
		return rows, fmt.Errorf("rows not initialized\n")
	}

	if rows.Next == nil {
		return rows, errors.New("EOF")
	}

	rows.Prev = rows.Current
	rows.Current = rows.Next
	if scanner.Scan() {
		rows.Next = []rune(scanner.Text())
	} else {
		rows.Next = nil
	}

	return rows, nil

}

func InitRows(scanner *bufio.Scanner) (Rows, error) {
	var rows Rows
	rows.Prev = nil

	if !scanner.Scan() {
		rows, fmt.Errorf("failed to init rows, could not scanner.Scan()\n")
	}
	rows.Current = scanner.Text()

	if !scanner.Scan() {
		rows, fmt.Errorf("failed to init rows, could not scanner.Scan()\n")
	}
	rows.Next = Scanner.Text()
}


func FindRolls {
	const PAPER_ROLL rune = '@'

	var prev []rune = nil
	var current []rune
	var next []rune
	total_count := 0

	for scanner.Scan() {
		current = []rune(scanner.Text())
		if scanner.Scan() {
			next = []rune(scanner.Text())
		} else {
			next = nil
		}


		for i, r := range current {
			nearby := 0
			if r == PAPER_ROLL {
			}
		}



		prev = current
	}
}
