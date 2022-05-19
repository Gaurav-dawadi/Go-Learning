package main
import "fmt"


func area_square(l int64, b int64) int64{
	area := l*b
	return area
}

func main(){
	var (
		length int64 = 10
		width int64 = 5
	)
	result := area_square(length, width)
	fmt.Println("The area of square is: ", result)
}