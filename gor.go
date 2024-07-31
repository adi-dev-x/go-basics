package main

import (
	"fmt"
	"time"
)

func h(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go h("fir")
	//time.Sleep(100 * time.Millisecond)
	h("sec")
}
