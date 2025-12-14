package day03

import (
	"bufio"
	"fmt"
	"math"
)

func getBankJoltageBeforeY(line string, y int) (int, string, error) {

	l := len(line)
	max := 0
	searching := line[:l - y]
	var leftover string

	for i, r := range searching {
		n := int(r - '0')
		if n < 0 || n > 9 {
			err := fmt.Errorf("Found invalid rune in line: %v\n", r)
			return 0, "", err
		}

		if n > max {
			leftover = line[i + 1:]
			max = n
		}
	}


	return max, leftover, nil
}


func getBankJoltage(line string) (int, error) {
	max1 := 0
	max2 := 0
	j := 0


	for i, r := range line[:len(line) - 1] {
		n := int(r - '0')
		if n < 0 || n > 9 {
			err := fmt.Errorf("Found invalid rune in line: %v\n", r)
			return 0, err
		}

		if n > max1 {
			j = i + 1
			max1 = n
		}
	}

	for _, r := range line[j:]{
		n := int(r - '0')
		if n < 0 || n > 9 {
			err := fmt.Errorf("Found invalid rune in line: %v\n", r)
			return 0, err
		}
		if n > max2 {
			max2 = n
		}
	}


	final := max1 * 10
	final += max2

	return final, nil

}

func FindAmountOfJoltages(scanner *bufio.Scanner, amount int) ([]int, error) {
	var joltages []int
	for scanner.Scan() {


		bank := scanner.Text()
		bank_total := 0
		var joltage int
		for i := amount - 1; i >= 0; i-- {
			var err error
			joltage, bank, err = getBankJoltageBeforeY(bank, i)


			if err != nil {
				err = fmt.Errorf("Could not getBankJoltage on bank: %v\n got error: %v\n", bank, err)
			}

			if i > 0 {
				bank_total += joltage * int(math.Pow10(i))
			} else {
				bank_total += joltage
			}
		}


		joltages = append(joltages, bank_total)
	}

	return joltages, nil
}

func FindJoltages(scanner *bufio.Scanner) ([]int, error) {
	var joltages []int
	for scanner.Scan() {
		line := scanner.Text()
		joltage, err := getBankJoltage(line)
		if err != nil {
			err = fmt.Errorf("Could not getBankJoltage on line: %v\n got error: %v\n", line, err)
		}
		joltages = append(joltages, joltage)
	}

	return joltages, nil
}
