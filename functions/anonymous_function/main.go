package main
import ("fmt")

// Assigning function to a variable
var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main(){
	age := 24
	func(){
		fmt.Println("His age is ", age)
	}()
	fmt.Println("The area is ",area(5,5))

	// Passing params to anonymous function
	func(l int, b int){
		fmt.Println(l*b)
	}(4,6)
}

