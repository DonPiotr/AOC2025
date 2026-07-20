package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Red struct {
	x int
	y int
}

type Border struct {
	r1, r2 Red
	length int
}

func ReadTile(s string) Red {
	a := strings.Split(strings.TrimRight(s, "\n"),",")
	nred := Red{}
	nred.x, _ := strconv.Atoi(a[0])
	nred.y, _ := strconv.Atoi(a[1])
	return nred
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

func printOneRed(r Red) string {
	s := "[ "
	s += "x: " + strconv.Itoa(r.x) + ", y: " + strconv.Itoa(r.y)
	s += " ]\n"
	return s
}

func printReds(a []Red) string {
	s := "{\n"
	for _, r := range a {
		s += "\t" + printOneRed(r)
	}
	s += "}"
	return s
}

func lineLength(l Line) int {
	if l.r1.x == l.r2.x {
		return AbsInt(l.r1.y-l.r2.y)
	} else {
		return AbsInt(l.r1.x-l.r2.x)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reds := []Red{}
	last := &Red{}
	for scanner.Scan() {
		r := Red{pos:ReadTile(scanner.Text())}
	}
	fmt.Println(printReds(reds))

	max_area := 0
	for ri := 0; ri < len(reds)-3; ri++ {
		for rj := 0; rj < len(reds)-3; rj++ {
			rt_i := reds[ri]
			rt_j := reds[rj]
			a := AreaRect(rt_i.pos,rt_j.pos)
			//fmt.Println(rt_i, rt_j, a)
			if a > max_area {
				max_area = a
			}
		}
	}

	fmt.Println(max_area)
}
