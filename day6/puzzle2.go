package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Operation struct {
	args []int
	op   byte
}

func getCol(d []string, i int) string {
	s := ""
	for _, r := range d {
		s += string(r[i])
	}
	return s
}

func calcOp(o Operation) int {
	result := 0
	if o.op == 42 {
		result = 1
	}
	for _, a := range o.args {
		if o.op == 42 {
			result *= a
		} else {
			result += a
		}
	}
	return result
}

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	data := []string{}
	cols := 0
	for scanner.Scan() {
		s := strings.TrimRight(scanner.Text(), "\n")
		if cols < len(s) {
			cols = len(s)
		}
		data = append(data,s)
	}
	for i, row := range data {
		n := cols - len(row)
		if n > 0 {
			data[i] = row + strings.Repeat(" ", n)
		}
	}

	var o Operation
	for i := cols-1; i >= 0; i-- {
		c := getCol(data,i)
		if strings.TrimSpace(c) == "" {
			// colonna vuota -> è tempo di contare
			sum += calcOp(o)
			o = Operation{}
			continue
		}

		if c[len(c)-1] != ' ' {
			// un numero con un operatore
			o.op = c[len(c)-1]
			c = c[:len(c)-1]
		}
		// un numero senza un operatore
		c = strings.TrimSpace(c)
			n, _ := strconv.Atoi(c)
		o.args = append(o.args, n)
	}
	// ultimo conteggio
	sum += calcOp(o)

	fmt.Println(sum)
}
