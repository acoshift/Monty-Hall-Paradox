package main

import (
	"crypto/rand"
	"fmt"
)

func random() int {
	var b [1]byte
	rand.Read(b[:])
	return int(b[0])%3 + 1
}

func rem(p []int, x int) []int {
	r := []int{}
	for _, y := range p {
		if x == y {
			continue
		}
		r = append(r, y)
	}
	return r
}

func sel(p []int, x int, swap bool) int {
	if !swap {
		return x
	}
	p = rem(p, x)
	return p[0]
}

func main() {
	all := []int{1, 2, 3}
	var correct, wrong int
	for i := 0; i < 10000; i++ {
		var x int
		k := random()
		// fmt.Printf("choose one from %v: ", all)
		// fmt.Scanln(&x)
		x = random()

		s := rem(rem(all, k), x)[0]
		// fmt.Printf("%d is wrong\n", s)

		// fmt.Printf("choose one from %v: ", rem(all, s))
		// fmt.Scanln(&x)
		x = sel(rem(all, s), x, false)

		if x == k {
			// fmt.Println("correct!!!")
			correct++
		} else {
			// fmt.Println("wrong!!!")
			wrong++
		}
		// fmt.Println()
	}
	fmt.Printf("\n\ncorrect: %d\nwrong: %d\n", correct, wrong)
}
