package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseNumbers() []int {
	nbString := os.Args[1 : len(os.Args)-1]
	var nb []int
	for _, v := range nbString {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("bad argument: %q: %v", v, err)
		}
		nb = append(nb, i)
	}
	return nb
}

func main() {

	nb := parseNumbers()

	mulStr := os.Args[len(os.Args)-1]
	mul, err := strconv.ParseFloat(mulStr, 64)
	if err != nil {
		log.Fatalf("bad argument: %q: %v", mulStr, err)
	}

	log.Println(mul, nb)
	sum := compute(mul, nb)
	fmt.Println("total:", sum)
}
