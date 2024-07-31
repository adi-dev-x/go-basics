package main

import (
	"fmt"
)

func g(d ...int) int {
	v := 0
	for _, i := range d {
		v = v + i
	}
	return v

}
func main() {
	fmt.Println(g(10, 30, 40))

	kl := [5]int{5, 5, 5, 5, 5}
	p := make(map[int]int)
	for i := 0; i < 5; i++ {
		p[i] = kl[i]
	}
	fmt.Println(p)
	for _, values := range p {
		fmt.Println(values)

	}
}
