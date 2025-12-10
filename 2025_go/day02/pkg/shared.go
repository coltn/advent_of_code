package day02

import (
	"bufio"
	"io"
	"strconv"
	"slices"
)

type RangeInt struct {
	Start int
	End int
}


func FindInvalidIdHalf(ranges []RangeInt) []int {
	var invalid []int
	for _, id_range := range ranges {
		for current := id_range.Start; current <= id_range.End; current ++ {
			elem := []rune(strconv.Itoa(current))
			l := len(elem)
			if (l / 2) % 2 == 0 {
				if slices.Compare(elem[:l / 2], elem[l / 2:]) == 0 {
					invalid = append(invalid, current)
				}
			}
		}
	}

	return invalid
}

func ParseInputToRanges(reader *bufio.Reader) ([]RangeInt, error) {
	var ranges []RangeInt

	for {
		start_data, err := reader.ReadBytes(byte('-'))
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return ranges, err
			}
		}
		start, err := strconv.Atoi(string(start_data[:len(start_data) - 1]))
		if err != nil {
			return ranges, err
		}

		end_data, err := reader.ReadBytes(byte(','))
		if err != nil && err != io.EOF {
			return ranges, err
		}
		end, err := strconv.Atoi(string(end_data[:len(end_data) - 1]))
		if err != nil {
			return ranges, err
		}
		
		ranges = append(ranges, RangeInt {
			Start: start,
			End: end,
		})
	}

	return ranges, nil
}
