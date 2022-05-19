package main
import ("fmt")

/*
UDF allows us to define our own function types
*/

// Using UDF for HOF's squareSum function from file HOF/main.go
type First func(int) int
type Second func(int) First

func squareSum(x int) Second{
	return func(y int) First{
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main(){
	fmt.Println("The result of squareSum is ", squareSum(5)(4)(3))
}
