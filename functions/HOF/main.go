package main
import ("fmt")

// Higher Order Function that returns a function
func doSum() func(x int, y int) int{
	return func(x int, y int) int{
		return x+y
	}
}

// Higher Order Function that takes variable as argument and returns function
func findDiff(x int) func(y int) int{
	return func(y int) int{
		return x-y
	}
}

// Higher Order Function that takes variable as argument and returns function
func sum(x int, y int) int{
	return x+y
}
func partialSum(x int) func(y int) int{
	return func(y int) int{
		return sum(x,y)
	}
}

// Higher Order Function that takes function as an argument and returns function
func cal(a int,b int, performCalculation func(c int,d int) int) int{
	return performCalculation(a,b)
}

//Higher Order Function that returns functions from other Functions
func squareSum(x int) func(y int) func(z int) int{
	return func(y int) func(z int) int{
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}


func main(){
	result_sum := doSum()
	fmt.Println("The result of sum is ", result_sum(10,20))

	result_partial_sum := partialSum(10)
	fmt.Println("The result of partial sum is ", result_partial_sum(5))

	// result_diff := findDiff(15)
	// fmt.Println("The result of difference is ", result_diff(5))
	fmt.Println("The result of difference is ", findDiff(15)(5))  // Another way of calling function other than of above way

	result_cal := cal(10,2, func(x int, y int) int {return x+y})
	fmt.Println("The result of cal is ", result_cal)        // Note: we are not calling result_cal variable

	fmt.Println("The result of squareSum is ", squareSum(5)(5)(5))
}