package day03

import (
	"bufio"
	"fmt"
)


func getBankJoltage(line string) (int, error) {
	fmt.Printf("debug: line = %v\n", line)
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

	fmt.Printf("debug: max1 = %v\n", max1)
	fmt.Printf("debug: max2 = %v\n", max2)

	final := max1 * 10
	final += max2
	fmt.Printf("debug: final = %v\n", final)

	return final, nil

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
