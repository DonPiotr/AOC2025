package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"bytes"
	"slices"
	"sort"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	manifold := [][]byte{}

	for scanner.Scan() {
		row := []byte(strings.TrimRight(scanner.Text(), "\n"))
		manifold = append(manifold, row)
	}
	max_x := len(manifold[0])-1
	max_y := len(manifold)-1

	prev_row_beams := []int{}
	{
		s_x := bytes.IndexByte(manifold[0],'S')
		prev_row_beams = append(prev_row_beams,s_x)
	}

	for y := 1; y <= max_y; y++ {
		//fmt.Println("Riga:",y)
		new_beams := []int{}
		for _, x := range prev_row_beams{
			if manifold[y][x] == '.' {
				new_beams = append(new_beams,x)
			} else if manifold[y][x] == '^' {
				sum++
				if x > 0 {
					new_beams = append(new_beams,x-1)
				}
				if x < max_x {
					new_beams = append(new_beams,x+1)
				}
			}
		}
		sort.Ints(new_beams)
		prev_row_beams = slices.Compact(new_beams)
		//fmt.Println(prev_row_beams)
	}

	fmt.Println(sum)
}
