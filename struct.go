package main
import("fmt")
type st struct{
	name string

}
func (ps st) hname(){
	fmt.Println(ps.name)
}
func main(){
	v:=st{name:"jjj"}
	v.hname()

}