package main
import ("fmt")

type geo interface{
	area()
}
type rec struct{
	w,h float64
}
type cir struct{
	r float64
}
func (r rec) area(){
	fmt.Println(r.w*r.h)
}
func (r cir) area(){
	fmt.Println(r.r*r.r)
}
func me(g geo){
	g.area()
}
func main(){
	f:=rec{h:4,w:9}
	g:=cir{r:13}
	me(f)
	me(g)
}

