package main

import (
	"fmt"
	"bufio"
	"os"
)

func arr2num(a []int) int {
	num := 0
	for _, d := range a {
		num = num*10 + d
	}
	return num
}

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var big_num [12]int     // viene inizializzato con tutti i zeri
		var indexes [12]int

		start_search_index := 0

		for digit_index := 0; digit_index <= 11; digit_index++ {
			for i := start_search_index; i <= len(line)-12+digit_index; i++ {
				digit := int(line[i] - '0')
				if digit > big_num[digit_index] {
					big_num[digit_index] = digit
					indexes[digit_index] = i
				}
				if digit == 9 {
					break
				}
			}
			start_search_index = indexes[digit_index]+1
		}

		// fmt.Println(arr2num(big_num[:]))
		sum += arr2num(big_num[:])
	}
	// fmt.Println("==================")
	fmt.Println(sum)
}

