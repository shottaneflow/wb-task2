package main

import "fmt"

type Human struct{
	name string
	age int
}
 
func (h *Human) getAge() int{
	return h.age
}
type Action struct{
	Human
}

func main(){
	h:=Human{name: "Danil",age: 22}
	a:=Action{Human: h}
	fmt.Println(a.getAge())
}

