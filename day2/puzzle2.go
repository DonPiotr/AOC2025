package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func haPatternRipetuti(s string) bool {
	var divisori = [15][]int{
		{}, {}, // per 0 e 1
		{1}, // per 2
		{1}, // per 3
		{1,2}, // per 4
		{1}, // per 5
		{1,2,3}, // per 6
		{1}, // per 7
		{1,2,4}, // per 8
		{1,3}, // per 9
		{1,2,5}, // per 10
		{1}, // per 11
		{1,2,3,4,6}, // per 12
		{1}, // per 13
		{1,2,7}, // per 14
	}
	len_s := len(s)
	for _, d := range divisori[len_s] {
		prefix := s[:d]
		s2 := strings.Repeat(prefix, len_s/d)
		if s == s2 {
			return true
		}
	}
	return false
}

func main() {
	sum := 0

	data, _ := io.ReadAll(os.Stdin)
	data_str := string(data)
	data_str = strings.TrimSuffix(data_str, "\n")
	data_arr := strings.Split(data_str, ",")

	for _, r := range data_arr {
		range_start_str, range_stop_str, _ := strings.Cut(r, "-")

		range_start, _ := strconv.Atoi(range_start_str)
		range_stop, _ := strconv.Atoi(range_stop_str)

		for i := range_start; i <= range_stop; i++ {
			i_str := strconv.Itoa(i)
			if !haPatternRipetuti(i_str) {
				continue
			}
			// fmt.Println(i_str)
			sum += i
		}
	}

	fmt.Println(sum)
}

