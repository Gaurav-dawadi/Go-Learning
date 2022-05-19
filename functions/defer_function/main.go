package main
import "fmt"

/*
defer statements delay the execution of the function or method or an anonymous
 method until the nearby functions returns.
*/

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func third() {
	fmt.Println("Close")
}

// FILO or LIFO
func main() {
	fmt.Println("--------------------------------")
	fmt.Println("Check this Output without defer")
	fmt.Println("--------------------------------")
	third()
	second()
	first()

	fmt.Println("--------------------------------")
	fmt.Println("Check this Output with one defer")
	fmt.Println("--------------------------------")
	defer third()
	second()
	first()

	fmt.Println("--------------------------------")
	fmt.Println("Check this Output with all defer")
	fmt.Println("--------------------------------")
	defer third()
	defer second()
	defer first()
}