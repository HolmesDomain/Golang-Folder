//Andrew Holmes
package main

import ("fmt"
		 "math/rand")

var num1,num2 float64 = 5.5,9.5
		 
func add(x,y float64) float64 {
	return x + y
}
		 
func main() {
	fmt.Println("A number from 1 to 100", rand.Intn(100))
	
	fmt.Println(add(num1,num2))
}