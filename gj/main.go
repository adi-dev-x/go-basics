package main

import (
	"fmt"
)

type sh interface {
	area() float64
}

type cir struct {
	r float64
}
type rec struct {
	r float64
	l float64
}

func (r cir) area() float64 {
	return 3.14 * r.r

}
func (r rec) area() float64 {
	return r.l * r.r

}

func main() {

	p := rec{r: 12, l: 12}
	pc := cir{r: 121}
	fmt.Println(pc.area())
	fmt.Println(p.area())

}
