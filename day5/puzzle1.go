package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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

func main() {
	sum := 0
	var fresh []Range
	var available []int
	var firstPart = true

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := strings.TrimRight(scanner.Text(), "\n")
		if row == "" {
			firstPart = false
		} else if firstPart {
			var myrange Range
			parts := strings.Split(row,"-")
			myrange.first, _ = strconv.Atoi(parts[0])
			myrange.last, _ = strconv.Atoi(parts[1])
			fresh = append(fresh, myrange)
		} else {
			id, _ := strconv.Atoi(row)
			available = append(available, id)
		}
	}
	for _, id := range available {
		for _, r :=range fresh {
			if inRange(r, id) {
				sum++
				break
			}
		}
	}
	fmt.Println(sum)
}

