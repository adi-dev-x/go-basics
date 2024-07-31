package main

import (
	"fmt"
)

func gl(h ...int) int {
	v := 0
	for _, i := range h {
		v = v + i
	}
	fmt.Println("h1111")
	return v
}
func gl2(h ...int) int {
	v := 0
	for _, i := range h {
		v = v + i
	}
	fmt.Println("h2222")
	return v
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	p := make(map[int]int)
	for i := 0; i < 5; i++ {
		p[i] = i + 1*2
	}
	fmt.Println(p)
	for _, values := range p {
		fmt.Println(p[values])

	}
	go func() {
		v := gl(4, 4, 4, 4, 4, 4)
		ch1 <- v
		// close(ch1)

	}()

	go func() {
		v := gl2(4, 4, 4, 4, 4, 4)
		ch2 <- v
		// close(ch2)

	}()

	k := <-ch1
	k1 := <-ch2
	fmt.Println("Svdwsvwdsvw", k, "cadd", k1)

}
