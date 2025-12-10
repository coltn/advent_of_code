package day02

import (
	"bufio"
	"strconv"
	"io"
)

type Range_t struct {
	Start int
	End int
}

func ParseInputToRanges(reader *bufio.Reader) ([]Range_t, error) {
	var ranges []Range_t

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
		
		ranges = append(ranges, Range_t{
			Start: start,
			End: end,
		})
	}

	return ranges, nil
}
