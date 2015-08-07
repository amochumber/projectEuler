package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := 0
	sum := 0

	for (a + b) < 4000000 {
		c = a + b
		a = b
		b = c
		if c%2 == 0 {
			sum += c
		}
	}
	fmt.Println(sum)
}
