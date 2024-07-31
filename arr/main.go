package main

import (
	"fmt"
)

func rev(s string) string {

	var k = []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		k[i], k[j] = k[j], k[i]

	}
	return string(k)

}

func main() {
	k := make([]int, 5, 10)

	fmt.Println(rev("afqjhgfds"))
}
