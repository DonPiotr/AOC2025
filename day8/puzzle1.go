package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
	"sort"
)

type Box struct {
	x, y, z, c int
}

type Pair struct {
	b1, b2 *Box
	d int
}

func atoi(s string) int{
	i, _ := strconv.Atoi(s)
	return i
}

func dist(b1,b2 Box) int {
	dx := b1.x - b2.x
	dy := b1.y - b2.y
	dz := b1.z - b2.z
	return int(math.Sqrt(float64(dx*dx+dy*dy+dz*dz)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	boxes := []Box{}
	for scanner.Scan() {
		box_s := strings.TrimRight(scanner.Text(), "\n")
		box_sa := strings.Split(box_s,",")
		box := Box{
			x: atoi(box_sa[0]),
			y: atoi(box_sa[1]),
			z: atoi(box_sa[2]),
			c: -1,
		}
		boxes = append(boxes,box)
	}
	n := 10
	if len(boxes) > 20 {
		n = 1000
	}

	// generiamo coppie
	pairs := []Pair{}
	for i := 0; i < len(boxes)-1; i++ {
		for j := i+1; j < len(boxes); j++ {
			p := Pair{}
			p.b1 = &boxes[i]
			p.b2 = &boxes[j]
			p.d = dist(boxes[i],boxes[j])
			pairs = append(pairs,p)
		}
	}
	sort.Slice(pairs, func(i,j int) bool {
		return pairs[i].d < pairs[j].d
	})

	// generiamo circuiti
	next_new_circuit_id := 0
	for pi := 0; pi < n; pi++ {
		p := pairs[pi]
		if p.b1.c == -1 && p.b2.c == -1 {
			// entrambi non in circuito -> creiamo nuovo
			p.b1.c = next_new_circuit_id
			p.b2.c = next_new_circuit_id
			next_new_circuit_id++
		} else if p.b1.c != -1 && p.b2.c != -1 {
			// entrambi sono nei circuiti -> se diversi gli coleghiamo
			if p.b1.c != p.b2.c {
				var major int
				var minor int
				if p.b1.c > p.b2.c {
					major = p.b1.c
					minor = p.b2.c
				} else {
					major = p.b2.c
					minor = p.b1.c
				}
				for bi := range boxes {
					if boxes[bi].c == major {
						boxes[bi].c = minor
					}
				}
			}
		} else if p.b1.c != -1 {
			// primo è nel circuito -> aggiungere l'altro
			p.b2.c = p.b1.c
		} else {
			// secondo  è nel circuito -> aggiungere il primo
			p.b1.c = p.b2.c
		}
	}
	sort.Slice(boxes, func(i,j int) bool {
		return boxes[i].c > boxes[j].c
	})

	// contiamo circuiti per numero di boxes
	cont := map[int]int{}
	for bi := range boxes {
		if boxes[bi].c != -1 {
			cont[boxes[bi].c]++
		}
	}
	cont2 := [][]int{}
	for id, num := range cont {
		cont2 = append(cont2, []int{id,num})
	}
	sort.Slice(cont2, func(i,j int) bool {
		return cont2[i][1] > cont2[j][1]
	})
	result := cont2[0][1]*cont2[1][1]*cont2[2][1]
	fmt.Println(result)

}
