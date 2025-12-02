package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func main() {
	sum := 0

	data, _ := io.ReadAll(os.Stdin)
	data_str := string(data)
	data_str = strings.TrimSuffix(data_str, "\n")
	data_arr := strings.Split(data_str, ",")

	for _, r := range data_arr {
		range_start_str, range_stop_str, e := strings.Cut(r, "-")

		range_start, _ := strconv.Atoi(range_start_str)
		range_stop, _ := strconv.Atoi(range_stop_str)

		for i := range_start; i <= range_stop; i++ {
			i_str := strconv.Itoa(i)
			if len(i_str)%2 != 0 {
				continue
			}
			mid := len(i_str)/2
			str_l := i_str[:mid]
			str_r := i_str[mid:]
			if str_l == str_r {
				fmt.Println(i_str)
				sum += i
			}
		}
	}

	fmt.Println(sum)
}

