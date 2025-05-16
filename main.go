package main

import "fmt"

type User struct {
	Name string
	Age int
}


func add(a,b int) int {
	return a+b
}

func divide (a,b int) (int,error){
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a/b,nil
}


func (u User) Greet() string {
	return "Hello " + u.Name
}


func main() {
	u := User{"Jay",25}
	a := u.Greet()

	add := add(10,10)

	d,error := divide(10,0)
	if error != nil {
		fmt.Println("Errror :",error)
	}

	fmt.Println(d)


    fmt.Println("Hello, Go!",a,add)
}