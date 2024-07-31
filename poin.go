package main
import("fmt")


func re(x *int) x int {
    *x=100
	return 
}
func main(){
	g:=10
	fmt.Println(re(&g))
}