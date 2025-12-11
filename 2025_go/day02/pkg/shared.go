package day02

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
)

type RangeInt struct {
	Start int
	End int
}

func isInvalidRepeating(elems []rune) bool {
	l := len(elems)

	/* avoid looping if only two elems */
	if l == 2 {
		return elems[0] == elems[1]
	}

	/* check for k number of patterns */
	for k := 1; k <= l / 2; k++ {
		/* continue if length of elems cannot be divided by k */
		if l % k != 0 {
			continue
		}

		/* the pattern we are matching for is elems before k */
		pattern := elems[:k]
		found := true

		for i := k; i < l; i += k {
			if slices.Compare(elems[i:i+k], pattern) != 0 {
				found = false
				break
			}
		}

		if found {
			return true
		}
	}

	return false
}

/* unused, didn't work--keeping for future ref */
func isInvalidEven(elems []rune) bool {

	if isInvalidOdd(elems) {
		fmt.Printf("debug: found even with all elems matching: %v\n", elems)
		return true
	}
	
	l := len(elems)
	split := l / 2

	if l == 2 {
		return elems[0] == elems[1]
	}

	for {
		found := true
		i := split
		for offset := range (l / split) {
			if slices.Compare(elems[:i], elems[i:]) != 0 {
				found = false
				break
			}
			i = (offset + 1) * split
		}

		if found {
			return true
		}
		

		split /= 2
		if split % 2 != 0 {
			return false
		}
	}
}

/* unused, didn't work--keeping for future ref */
func isInvalidOdd(elems []rune) bool {
	for _, elem := range elems {
		if elem != elems[0] {
			return false
		}
	}
	return true
}

func FindInvalidIdRepeating(ranges []RangeInt) []int {
	var invalid []int
	for _, id_range := range ranges {
		for current := id_range.Start; current <= id_range.End; current ++ {
			elems := []rune(strconv.Itoa(current))
			l := len(elems)
			if l == 1 {
				continue
			}

			if isInvalidRepeating(elems) {
				fmt.Printf("debug: found InvalidRepeating: %v\n", current)
				invalid = append(invalid, current)
			}

			// even number of elems
			/* unused, didn't work
			* if l % 2 == 0 {
			* 	if isInvalidEven(elems) {
			* 		fmt.Printf("debug: found InvalidEven: %v\n", elems)
			* 		invalid = append(invalid, current)
			* 	}
			* // odd number of elems
			* } else {
			* 	if isInvalidOdd(elems) {
			* 		fmt.Printf("debug: found InvalidOdd: %v\n", elems)
			* 		invalid = append(invalid, current)
			* 	}
			* }
			*/
		}
	}

	return invalid
}

func FindInvalidIdHalf(ranges []RangeInt) []int {
	var invalid []int
	for _, id_range := range ranges {
		for current := id_range.Start; current <= id_range.End; current ++ {
			elems := []rune(strconv.Itoa(current))
			l := len(elems)
			if (l / 2) % 2 == 0 {
				if slices.Compare(elems[:l / 2], elems[l / 2:]) == 0 {
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
