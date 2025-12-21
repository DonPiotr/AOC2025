package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Red struct {
	pos [2]int
	in byte 	// 0 undefined, 78 - N, 69 - E, 83 - S, 87 - W
	out byte
}

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

func printOneRed(r Red) string {
	s := "[ "
	s += "pos: " + strconv.Itoa(r.pos[0]) + "," + strconv.Itoa(r.pos[1]) + " , "
	s += "in: " + string(r.in) + " , out: " + string(r.out)
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

type Line struct{
	p1, p2	Red
	length	int
}

func lineLength(l Line) int {
	if l.p1.pos[0] == l.p2.pos[0] {
		return AbsInt(l.p1.pos[1]-l.p2.pos[1])
	} else {
		return AbsInt(l.p1.pos[0]-l.p2.pos[0])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reds := []Red{}
	last := &Red{}
	for scanner.Scan() {
		r := Red{pos:ReadTile(scanner.Text())}
		if (*last != Red{}) {
			switch {
			// NW corner == 0,0
			// 0 undefined, 78 - N, 69 - E, 83 - S, 87 - W
			case last.pos[0]==r.pos[0] && last.pos[1]<r.pos[1]:
				last.out = 83
				r.in = 78
			case last.pos[0]==r.pos[0] && last.pos[1]>r.pos[1]:
				last.out = 78
				r.in = 83
			case last.pos[0]<r.pos[0] && last.pos[1]==r.pos[1]:
				last.out = 87
				r.in = 69
			case last.pos[0]>r.pos[0] && last.pos[1]==r.pos[1]:
				last.out = 69
				r.in = 87
			}
		}
		reds = append(reds, r)
		last = &reds[len(reds)-1]
	}
	{
		r := &reds[0]
		switch {
		// NW corner == 0,0
		// 0 undefined, 78 - N, 69 - E, 83 - S, 87 - W
		case last.pos[0]==r.pos[0] && last.pos[1]<r.pos[1]:
			last.out = 83
			r.in = 78
		case last.pos[0]==r.pos[0] && last.pos[1]>r.pos[1]:
			last.out = 78
			r.in = 83
		case last.pos[0]<r.pos[0] && last.pos[1]==r.pos[1]:
			last.out = 87
			r.in = 69
		case last.pos[0]>r.pos[0] && last.pos[1]==r.pos[1]:
			last.out = 69
			r.in = 87
		}
	}
	fmt.Println(printReds(reds))

	// in the end we need test last red with 3 consecutives
	reds = append(reds,reds[0])
	reds = append(reds,reds[1])
	reds = append(reds,reds[2])

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
