package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"bytes"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	manifold := [][]byte{}

	for scanner.Scan() {
		row := []byte(strings.TrimRight(scanner.Text(), "\n"))
		manifold = append(manifold, row)
	}
	max_x := len(manifold[0])-1
	max_y := len(manifold)-1

	prev_row_beams := make(map[int]int)
	prev_row_beams[bytes.IndexByte(manifold[0],'S')] = 1

	for y := 1; y <= max_y; y++ {
		new_beams := make(map[int]int)
		for x,_ := range prev_row_beams{
			if manifold[y][x] == '.' {
				new_beams[x] += prev_row_beams[x]
			} else if manifold[y][x] == '^' {
				if x > 0 {
					new_beams[x-1] += prev_row_beams[x]
				}
				if x < max_x {
					new_beams[x+1] += prev_row_beams[x]
				}
			}
		}
		prev_row_beams = new_beams
		//fmt.Println(prev_row_beams)
	}
	//fmt.Println(prev_row_beams)

	sum := 0
	for _, v := range prev_row_beams {
		sum += v
	}
	fmt.Println(sum)
}
