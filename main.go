package main

import "fmt"

type User struct {
	Name string
	Age int
}



func (u User) Greet() string {
	return "Hello " + u.Name
}


func main() {
	u := User{"Jay",25}
	a := u.Greet()

    fmt.Println("Hello, Go!",a)
}