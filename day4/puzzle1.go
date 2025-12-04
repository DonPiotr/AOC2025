package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	sum := 0
	var mymap []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.TrimRight(scanner.Text(), "\n")
		mymap = append(mymap, row)
	}
	max_x := len(mymap[0])-1
	max_y := len(mymap)-1

	for y, row := range mymap {
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
					if mymap[y + dy][x + dx] == '@'{
						rolls++
					}
				}
			}
			if rolls < 4 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

