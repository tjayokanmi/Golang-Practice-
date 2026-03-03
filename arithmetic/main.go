package main

import "fmt"

func main() {
	x, y := 5, 6
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", (x / y))
	fmt.Println("x mod y = ", x%y)

	isbool := true
	hate := false
	fmt.Println(isbool && hate) // True and False should give False
	fmt.Println(isbool || hate) // True or False should give True
	fmt.Println(!hate)          // Not false is true
	fmt.Println(!isbool)        // Not true is false

}

// This is a short exercise from Edureka Golang toturial on youtube "https://youtu.be/Q0sKAMal4WQ?t=770"
