package main
import "fmt"

//var d[6] int,g[4] string
var(
	 d[6] int
	g[6] string

)
func main(){
	d[5]=3
	g[5]="wdwf"
	fmt.Println(d,"sdd",g)
	maps:= map[int]string{
		1:"h",
		2:"f",
	}
	fmt.Println(maps)
    for id,pet :=range maps{
		fmt.Println(id,pet)
	}
	x := make([]float64, 5, 10)
	fmt.Println(x)


}