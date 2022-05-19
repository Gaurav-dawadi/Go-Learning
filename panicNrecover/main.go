package main
import "fmt"

func recoverFromPanic(){
	if err := recover(); err != nil{
		fmt.Println("Recovering system from Panic Mode: ", err)
	}
}

func main() {
	var action int
    fmt.Println("Enter 1 or 2")
    fmt.Scanln(&action)

    /*  Use of Switch Case in Golang */   
    switch action {
        case 1:
            fmt.Printf("I am One")
        case 2:
            fmt.Printf ("I am Two")
		default:
			defer recoverFromPanic()
			panic(fmt.Sprintf("I am a  panicked %d",action))
    }
}