
package day06

import (
	"bufio"
	"unicode"
)

type Problem struct {
	values []int;
	operator rune;
}


const DELIM byte = '\n'


func LoadFileIntoMemory(input *bufio.Reader) ([]Problem, error) {


	var problems []Problem

	row := 0
	for {
		line, err := input.ReadString(DELIM)
		if err != nil {
			break
		}

		col := 0
		for _, r := range line {
			if unicode.IsSpace(r) {
				continue
			}

			if len(problems) < col {
				
			}

			if unicode.IsNumber(r) {
				problems[col].values[row] = int(r)
				row ++
			} else {
				problems[col].operator = r
			}
			col ++
		}
	}

	return problems, nil
}
