package main
import "fmt"

type Gadgets uint8
const(
	camera Gadgets = iota
	memory
	bluetooth
	media
	messaging
)
type mobile struct{
	brand string
	model string
}
type smartphone struct{
	gadgets Gadgets
}
type android struct{
	smartphone
	mobile
}

func (s *smartphone) launch(){
	fmt.Println("New Smartphone launched")
	fmt.Println("The gadgets are: ", s.gadgets)
}
func (a *android) samsung(){
	fmt.Printf("New model %s by %s\n",a.model,a.brand)
}


func main(){
	new_phone := &android {}
	new_phone.model = "Galaxy S14"
	new_phone.brand = "Samsung"
	new_phone.gadgets = camera+memory+bluetooth+media+messaging
	new_phone.launch()
	new_phone.samsung()
}