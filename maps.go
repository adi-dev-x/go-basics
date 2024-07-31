package main
import("fmt")

func main(){
	h:=make(map[int] int)
	h[1]=12
	
	k:=make([] int,5)
	k[1]=8
    k[0]=0
	k[2]=9
	for i:=0;i<5;i++{
		h[i]=2*i
		k[i]=3*i
	}
	delete(h,1)

	fmt.Println("mappp",h)
	fmt.Println(k)
	o:=[5]int{1,2,3,4,5}
	p:=[]int{1,2,3,4,5}
	
	fmt.Println(o)
	fmt.Println(p)
}