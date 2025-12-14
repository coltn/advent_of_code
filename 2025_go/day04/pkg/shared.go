
package day04

import (
	"bufio"
)


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

func ReadRows(scanner *bufio.Scanner) {

}
