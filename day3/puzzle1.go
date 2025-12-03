package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		i_last := len(line)-1
		i_of_biggest := 0
		biggest_not_last := 0
		second_biggest := 0

		for i := range line {
			digit := int(line[i] - '0')
			if i < i_last && digit > biggest_not_last {
				biggest_not_last = digit
				second_biggest = 0
				i_of_biggest = i
			} else if i > i_of_biggest && digit > second_biggest {
				second_biggest = digit
			}
		}

		fmt.Println(biggest_not_last, second_biggest)
		sum += biggest_not_last*10 + second_biggest
	}
	fmt.Println(sum)
}

