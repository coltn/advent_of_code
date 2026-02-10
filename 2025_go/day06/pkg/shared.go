package day06

import (
	"bufio"
	"fmt"
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
	total := 0

	for _, p := range problems {
		sumProblem := 0
		for i, v := range p.Values {
			if i == 0 {
				sumProblem = v
				continue
			}

			sum, err := handleOperation(sumProblem, v, p.Operator)
			if err != nil {
				fmt.Printf("Error: %v", err)
				continue
			}
			sumProblem = sum
			

			

		}

		total += sumProblem

	}

	return total
}

func handleOperation(sum int, value int, operator string) (int, error) {
	switch operator {
	case "*":
		return sum * value, nil
	case "+":
		return sum + value, nil
	default:
		return 0, fmt.Errorf("unkown operator: %v\n", operator)
	}
}

