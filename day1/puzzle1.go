package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func move_pointer(oldp int, move int) int {
	return (oldp + move) % 100
}

func main() {
	pointer := 50
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		letter := line[0]
		number, _ := strconv.Atoi(line[1:])
		if letter == 'L' {
			number = -number
		}

		pointer = move_pointer(pointer, number)

		if pointer == 0 {
			sum++
		}
	}
	fmt.Println(sum)
}

