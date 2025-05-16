package main

import ("fmt")

func changeValue(p *int){ // this will receive add not variable value
	fmt.Println("WE are now in func and will update value of this add: ",p)
	*p = 102
}


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

	jay := 50
	fmt.Println("THIs is prior val", jay, "  on add :",&jay)
	changeValue(&jay)

	fmt.Println("THIs is final val", jay, "  on add :",&jay)


	

}