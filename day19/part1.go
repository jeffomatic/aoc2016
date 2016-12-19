package main

import "fmt"

const numElves = 3004953

type elf struct {
	id          int
	numPresents int

	left  *elf
	right *elf
}

func main() {
	elves := make([]elf, numElves, numElves)

	for i := 0; i < numElves; i++ {
		elves[i].id = i + 1
		elves[i].numPresents = 1

		if i != 0 {
			elves[i].right = &elves[i-1]
		}
		if i != numElves-1 {
			elves[i].left = &elves[i+1]
		}
	}

	// Connect first and last elves
	elves[0].right = &elves[numElves-1]
	elves[numElves-1].left = &elves[0]

	var cur *elf
	for cur = &elves[0]; cur != cur.left; cur = cur.left {
		cur.numPresents += cur.left.numPresents
		cur.left.numPresents = 0

		if cur.left.numPresents == 0 {
			cur.left = cur.left.left
		}
	}

	fmt.Println(cur.id)
}
