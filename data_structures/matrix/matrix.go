package main

import "log"

func initMatrix() {
	var c [][]int
	c = [][]int{
		{2, 3, 5},
		{4, 5, 5},
	}
	log.Print(c[0][0])
}
