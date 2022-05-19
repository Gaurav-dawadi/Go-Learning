package main
import ("fmt")


/*
Accepts a multiple number of arguments
*/

// Multiple arguments of same kind
func variadicString(s ...string){
	fmt.Println(s)
}

// Multiple arguments where first position is only one type and from there are of int type
func variadicMulti(s string, a ...int){
	fmt.Println(s, a)
}

// Multiple arguments of different type
func variadicMultiArgs(i ...interface{}){
	// fmt.Println(i)
	for _,v := range i{
		fmt.Println(v)
	}
}

func main()  {
	variadicString("Harry", "Son", "Dejan", "Skipp", "Hugo", "Romero")
	variadicMulti("Sergio",1,2,3,4,5)
	variadicMultiArgs("Sergio", 1, 1.6, true, []int{'a','b','c'}, map[string]int{"Harry": 28, "Son":29, "Hugo": 35, "Skipp": 21})
}