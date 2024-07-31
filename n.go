package main
import "fmt"


func main(){
	maps:=map[int]string{
		1:"po",
		2:"lo",
		3:"kl",
	}

	pr,ok:=maps[8]
	i:=2
	delete(maps,i)
	fmt.Println(ok,pr,maps)
}