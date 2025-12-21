package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
	"sort"
	"maps"
	"slices"
)

type Box struct {
	x, y, z int
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

func sliceShift[T any](s []T) (T, []T){
	f := s[0]
	s = s[1:]
	return f, s
}

func boxesDelete(bs []*Box, el *Box) []*Box {
	new_boxes := []*Box{}
	for _, v := range bs {
		if v != el {
			new_boxes = append(new_boxes,v)
		}
	}
	return new_boxes
}

func countCircuits(b_in_c map[*Box]int) int {
	cs := map[int]int{}
	for _, v := range b_in_c {
		cs[v]++
	}
	fmt.Println(cs)
	return len(slices.Collect(maps.Keys(cs)))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	boxes := []*Box{}
	for scanner.Scan() {
		box_s := strings.TrimRight(scanner.Text(), "\n")
		box_sa := strings.Split(box_s,",")
		box := &Box{
			x: atoi(box_sa[0]),
			y: atoi(box_sa[1]),
			z: atoi(box_sa[2]),
		}
		boxes = append(boxes,box)
	}

	// generiamo coppie
	pairs := []Pair{}
	for i := 0; i < len(boxes)-1; i++ {
		for j := i+1; j < len(boxes); j++ {
			p := Pair{}
			p.b1 = boxes[i]
			p.b2 = boxes[j]
			p.d = dist(*boxes[i],*boxes[j])
			pairs = append(pairs,p)
		}
	}
	sort.Slice(pairs, func(i,j int) bool {
		return pairs[i].d < pairs[j].d
	})

	// generiamo circuiti
	next_new_circuit_id := 1	// 0 - no circuit
	boxes_in_circs := map[*Box]int{}
	for {
		var p Pair
		p, pairs = sliceShift(pairs)
		if boxes_in_circs[p.b1] == 0 && boxes_in_circs[p.b2] == 0 {
			// entrambi non in circuito -> creiamo nuovo
			boxes_in_circs[p.b1] = next_new_circuit_id
			boxes_in_circs[p.b2] = next_new_circuit_id
			boxes = boxesDelete(boxes, p.b1)
			boxes = boxesDelete(boxes, p.b2)
			next_new_circuit_id++
		} else if boxes_in_circs[p.b1] != 0 && boxes_in_circs[p.b2] != 0 {
			// entrambi sono nei circuiti -> se diversi gli colleghiamo
			// non c'è bisognio di ricuperare id di circuito scartato
			if boxes_in_circs[p.b1] != boxes_in_circs[p.b2] {
				var major int
				var minor int
				if boxes_in_circs[p.b1] > boxes_in_circs[p.b2] {
					major = boxes_in_circs[p.b1]
					minor = boxes_in_circs[p.b2]
				} else {
					major = boxes_in_circs[p.b2]
					minor = boxes_in_circs[p.b1]
				}
				for k,_ := range boxes_in_circs {
					if boxes_in_circs[k] == major {
						boxes_in_circs[k] = minor
					}
				}
			}
		} else if boxes_in_circs[p.b1] != 0 {
			// primo è nel circuito -> aggiungere l'altro
			boxes_in_circs[p.b2] = boxes_in_circs[p.b1]
			boxes = boxesDelete(boxes, p.b2)
		} else {
			// secondo  è nel circuito -> aggiungere il primo
			boxes_in_circs[p.b1] = boxes_in_circs[p.b2]
			boxes = boxesDelete(boxes, p.b1)
		}
		if len(boxes) == 0 {
			// ultimo box collegato a un circuito; abbiamo capito che a quel punto c'è solo uno
			_ = countCircuits(boxes_in_circs)
			fmt.Println(p.b1.x * p.b2.x)
			break
		}
	}
}
