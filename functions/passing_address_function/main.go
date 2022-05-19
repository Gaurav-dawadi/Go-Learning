package main
import "fmt"


func userInfo(age *int64,name *string){
	fmt.Printf("%v is %v years old\n",*name,*age)
}

func marksInfo(name *string, marks *int64) string{
	result := fmt.Sprintf("%v scored %v in Math",*name,*marks)
	return result
}

func main(){
	var age int64 = 28
	var name string = "Harry" 
	var marks int64 = 98
	
	userInfo(&age,&name)
	result := marksInfo(&name,&marks)
	fmt.Println(result)
}