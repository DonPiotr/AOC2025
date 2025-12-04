package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func del_rolls(mm [][]byte) int {
	max_x := len(mm[0])-1
	max_y := len(mm)-1
	turn_sum := 0

	for y, row := range mm {
		for x, loc := range row {
			rolls := 0
			if loc == '.' {
				continue
			}
			for dy := -1; dy <=1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 && dx == 0 {
						continue
					}
					if y + dy < 0 || x + dx < 0 || y + dy > max_y || x + dx > max_x {
						continue
					}
					if mm[y + dy][x + dx] == '@'{
						rolls++
					}
				}
			}
			if rolls < 4 {
				mm[y][x] = '.'
				turn_sum++
			}
		}
	}
	return turn_sum
}

func main() {
	sum := 0
	var mymap [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := []byte(strings.TrimRight(scanner.Text(), "\n"))
		mymap = append(mymap, row)
	}

	for {
		rolls := del_rolls(mymap)
		if rolls == 0 {
			break
		}
		sum += rolls
	}
	fmt.Println(sum)
}

