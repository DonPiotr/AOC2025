package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Range struct {
	first int
	last int
}

func inRange(r Range, id int) bool {
	if id < r.first || id > r.last {
		return false
	}
	return true
}

func ranges_joinable(r1 Range, r2 Range) bool {
	if inRange(r1, r2.first) || inRange(r1, r2.last) {
		return true
	} else if r1.last + 1 == r2.first || r2.last + 1 == r1.first {
		return true
	}
	return false
}

func join_ranges(r1 Range, r2 Range) Range {
	limits := []int{r1.first,r2.first,r1.last,r2.last}
	sort.Ints(limits)
	return Range{first: limits[0], last: limits[3]}
}

func ints_in_range(r Range) int {
	return r.last-r.first+1
}

func main() {
	sum := 0
	var fresh []Range
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.TrimRight(scanner.Text(), "\n")
		if row == "" {
			break
		}
		var myrange Range
		parts := strings.Split(row,"-")
		myrange.first, _ = strconv.Atoi(parts[0])
		myrange.last, _ = strconv.Atoi(parts[1])
		fresh = append(fresh, myrange)
	}
	//fmt.Println(fresh)
	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i].first < fresh[j].first
	})
	//fmt.Println(fresh)

	for {
		joined := false
		for i_r1 := 0; i_r1 < len(fresh); i_r1++ {
			for i_r2:= i_r1+1; i_r2 < len(fresh); i_r2++ {
				if ranges_joinable(fresh[i_r1],fresh[i_r2]) {
					r := join_ranges(fresh[i_r1],fresh[i_r2])
					fresh = append(fresh[:i_r2], fresh[i_r2+1:]...)
					fresh = append(fresh[:i_r1], fresh[i_r1+1:]...)
					fresh = append(fresh, r)
					joined = true
					break
				}
			}
			if joined {
				break
			}
		}
		if !joined {
			break // stop for{} only if it didn't joined anything
		}
	}
	//fmt.Println(fresh)
	for _, r := range fresh {
		sum += ints_in_range(r)
	}
	fmt.Println(sum)
}

