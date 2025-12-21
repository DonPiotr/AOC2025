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

	min_x := 10000
	max_x := 0
	min_y := 10000
	max_y := 0
	for _, r := range reds {
		if r[0] < min_x {
			min_x = r[0]
		} else if r[0] > max_x {
			max_x = r[0]
		}
		if r[1] < min_y {
			min_y = r[1]
		} else if r[1] > max_y {
			max_y = r[1]
		}
	}

	fmt.Println("Min x", min_x)
	fmt.Println("Max x", max_x)
	fmt.Println("Min y", min_y)
	fmt.Println("Max y", max_y)
	fmt.Println("Delta x", max_x-min_x)
	fmt.Println("Delta y", max_y-min_y)


}
