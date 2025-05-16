package main

import ("fmt")

func changeValue(p *int){ // this will receive add not variable value
	fmt.Println("WE are now in func and will update value of this add: ",p)
	*p = 102
}


type User struct{
	Name string
	age int
}

func changeName(p *User){
	fmt.Println(p)
	(*p).Name = "JAy"
}


func main(){
	
	user := User{Name: "JAY22",age:10}
	fmt.Println(user)
	changeName(&user)
	fmt.Println(user)



	

}