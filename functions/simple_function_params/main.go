package main
import "fmt"


func area_square(l int64, b int64) int64{
	area := l*b
	return area
}

func calculation(l int64, b int64)(area int64, parameter int64){
	parameter = 2*(l+b)
	area = l*b
	return
}

func main(){
	var (
		length int64 = 10
		width int64 = 5
	)
	result_square := area_square(length, width)
	result_cal_area,result_cal_param := calculation(length, width)
	fmt.Println("The area of square is: ", result_square)
	fmt.Printf("The area is %v and parameter is %v\n", result_cal_area, result_cal_param)
}