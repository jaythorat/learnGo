package main

import ("fmt")







func main() {
	
	var age int = 3
	const jay = 3.12 // cannot update later

	fmt.Println(age,jay)

	age = 10
	jay = 1 // throws error
	fmt.Println(age, jay)


}