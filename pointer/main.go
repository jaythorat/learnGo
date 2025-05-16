package main

import ("fmt")


func main(){
	x := 10
	y := &x

	fmt.Println("X has value :",x)
	fmt.Println("Y has value :",*y)
	*y = 25

	var p *int

	j := 1
	p = &j

	fmt.Println(x,y,p,*p)
	

}