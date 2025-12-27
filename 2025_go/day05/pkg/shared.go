
package day05

import (
	"bufio"
	"io"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	Low int
	High int
}


type readerState int
const (
	READING_RANGES = iota
	READING_VALUES
)

const DELIM byte = '\n'

func LoadFileIntoMemory(input *bufio.Reader) ([]Range, []int, error) {
	var state readerState = READING_RANGES

	ranges := make([]Range, 0)
	values := make([]int, 0)

	for {
		data, err := input.ReadBytes(DELIM)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, nil, err
			}
		}

		if len(data) == 1 {
			state = READING_VALUES
			continue
		}

		ranges, values, err = handleData(data, ranges, values, state)
		if err != nil {
			return ranges, values, err
		}
	}

	return ranges, values, nil
}

func handleData(d []byte, r []Range, v []int, s readerState) ([]Range, []int, error) {


	switch s {
		case READING_RANGES:
			
			strs  := strings.Split(string(d), "-")
		
			if len(strs) != 2 {
				err := errors.New("Invalid range in ranges\n")
				return r, v, err
			}

			val_0, err := strconv.Atoi(strs[0])
			if err != nil {
				return r, v, fmt.Errorf("Could not Atoi val_0 when spliting range values: %v\n", err)
			}
			val_1, err := strconv.Atoi(strings.Trim(strs[1], string('\n')))
			if err != nil {
				return r, v, fmt.Errorf("Could not Atoi val_1 when spliting range values: %v\n", err)
			}

			r = append(r, Range{
				Low: val_0,
				High: val_1,
			})

			
			
		case READING_VALUES:
			val, err := strconv.Atoi(strings.Trim(string(d), string('\n')))
			if err != nil {
				return r, v, err
			}
			v = append(v, val)
	}


	return r, v, nil
}

func FindValuesInRanges(ranges []Range, values []int) int {
	matches := 0

	for _, v := range values {
		for _, r := range ranges {
			if v >= r.Low && v <= r.High {
				matches ++
				break
			}
		}
	}

	return matches
}


/*
 * for my second attempt:
 * I am going to sort `ranges` in asc, get unique id's from r.High - r.Low
 * then track r.High before inc r to next range, check if r.Low is < than the last
 * r.High, if so use r.High - last_r.High
 */

func CountValidIDsInRanges(ranges []Range) int {
	count := 0
	last_high := 0
	
	for	_, r := range ranges {
		count +=

	}


	return count
}

/* 
 * the function below does work on a smaller data set,
 * but once you use the full input file, n^2 items is too much to store, and it takes too long.
 * program runs until my laptop runs out of memory (32GiB)
 */
func oldCountValidIDsInRanges(ranges []Range) int {
	unique := make(map[int]struct{}, 0)
	
	for	_, r := range ranges {
		for id := r.Low; id <= r.High; id ++ {
			unique[id] = struct{}{}
		}

	}

	count := 0
	for range unique {
		count ++
	}

	return count
}
