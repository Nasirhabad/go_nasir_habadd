package main

import "fmt"

func main() {

	// time comlexity -> TO measure the effectveness of the algorithm
	//how to meatuse  time comlexty
	//1.look at  basic aperation (the operation that ocuurs most)
	//2.how long the operation excuted ?

	var n int //declear "n"variable

	fmt.Print("input number:")
	fmt.Scan(&n)

	// genarate square numbers up to "n"
	//o(n)
	for i := 1; i <= n; i++ {
		var squared int = i * i
		fmt.Println(squared)
	}

}
