
package main

import "fmt"

func main() {
	//fmt.Println(fact(7))
	
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n 
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

}
