package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	data := [][]string{}
	for scanner.Scan() {
		s := strings.TrimRight(scanner.Text(), "\n")
		fields := strings.Fields(s)
		data = append(data,fields)
	}

	rows := len(data)
	cols := len(data[0])

	operations := make([][]string, cols)
	for i := 0; i < cols; i++ {
		operations[i] = make([]string, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			operations[j][i] = data[i][j]
		}
	}

	for _, o := range operations {
		fmt.Println(o)
		op := o[len(o)-1]
		o = o[:len(o)-1]
		op_result := 0
		if op == "*" {
			op_result = 1
		}
		for _, a := range o {
			ai, _ := strconv.Atoi(a)
			if op == "+"  {
				op_result += ai
			} else {
				op_result *= ai
			}
		}
		sum += op_result
	}

	fmt.Println(sum)
}
