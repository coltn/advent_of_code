package day06

import (
	"bufio"
	"strconv"
	"strings"
)

type Problem struct {
	Values []int;
	Operator string;
}


const DELIM byte = '\n'


func LoadFileIntoMemory(input *bufio.Reader) ([]Problem, error) {


	var problems []Problem

	for {
		line, err := input.ReadString(DELIM)
		if err != nil {
			break
		}

		chars := strings.Split(line, " ")


		col := 0
		for _, r := range chars {

			if r == " " || r == "\n" || r == "" {
				continue
			}

			r = strings.TrimSpace(r)


			if len(problems) < col + 1 {
				problems = append(problems, Problem{})
			}

			n, err := strconv.Atoi(r)
			if err != nil {
				problems[col].Operator = r
				col ++
				continue
			}

			problems[col].Values = append(problems[col].Values, n)
			col ++
		}
	}

	return problems, nil
}

func HandleProblems(problems []Problem) int {
	totals := make([]int, len(problems))

	for i, p := range problems {
		sum := 0
		for _, v range p.Values {
			

		}

		total[i] = sum
	}
}
