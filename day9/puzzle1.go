package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func ReadTile(s string) [2]int {
	a := strings.Split(strings.TrimRight(s, "\n"),",")
	x, _ := strconv.Atoi(a[0])
	y, _ := strconv.Atoi(a[1])
	return [2]int{x,y}
}

func AbsInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func AreaRect(p1,p2 [2]int) int {
	return (1+AbsInt(p1[0]-p2[0]))*(1+AbsInt(p1[1]-p2[1]))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reds := [][2]int{}
	for scanner.Scan() {
		reds = append(reds, ReadTile(scanner.Text()))
	}

	max_area := 0
	for ri := 0; ri < len(reds)-1; ri++ {
		for rj := ri + 1; rj < len(reds); rj++ {
			rt_i := reds[ri]
			rt_j := reds[rj]
			a := AreaRect(rt_i,rt_j)
			//fmt.Println(rt_i, rt_j, a)
			if a > max_area {
				max_area = a
			}
		}
	}

	fmt.Println(max_area)
}
