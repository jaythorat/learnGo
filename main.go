package main

import ("fmt"
		"time")


func tp(ch chan string){
	time.Sleep(1 * time.Second)	
	fmt.Println("HI",<- ch)
	ch <- "Gora"
}




func main() {
	ch := make(chan string)
	// go sayHello(ch)
	// msg := <- ch

	go tp(ch)
	ch <- "Jayyyyyyy"
	fmt.Println(" Jay")
	fmt.Print(<-ch)


}