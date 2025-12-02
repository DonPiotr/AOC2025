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
		fmt.Println("* Riga:", line)
		letter := line[0]
		number, _ := strconv.Atoi(line[1:])

		hm100 := number / 100
		if hm100 > 0 {
			sum += hm100
			fmt.Println("Aggiunto: ", hm100)
			number -= hm100 * 100
		}

		if letter == 'L' {
			number = -number
		}

		if pointer != 0 && (pointer + number >= 100 || pointer + number <= 0) {
			fmt.Println("Aggiunto: 1")
			sum++
		}

		pointer = move_pointer(pointer, number)
		if pointer < 0 {
			pointer += 100
		}
		fmt.Println("Pointer: ", pointer)

	}
	fmt.Println(sum)
}

