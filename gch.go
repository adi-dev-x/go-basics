package main

import (
	"fmt"
)

func gl(g ...int) int {
	v := 0
	for _, i := range g {
		v = v + i

	}
	return v
}

func main() {
	kc := make(chan int)
	kc1 := make(chan int)
	go func() {
		va := <-kc
		kc1 <- va

	}()
	go func() {
		da := gl(123, 2, 4)
		kc <- da

	}()
	res := <-kc1
	fmt.Println(res)

}
