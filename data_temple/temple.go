package data_temple

import "fmt"

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human
	School string
	Loan   float32
}

type Employee struct {
	Human
	Company string
	Money   float32
}


func (h Human) SayHi(){
	fmt.Printf("hello,I am %s you can call me %s \n",h.Name,h.Phone)
}

func (h Human) Sing(lyrics string){
	fmt.Printf("I singing [%s] \n",lyrics)
}

func (h Human) Guzzle(beer string){
	fmt.Printf("I guzzling %s beer \n",beer)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beer string)
}