package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")

}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")

}

type Superman struct {
	Human
	level int
}

func (this *Superman) Eat() {

	fmt.Println("Superman.Eat()...")

}

func (this *Superman) Fly() {
	fmt.Println("Superman.Fly()...")

}

func main3() {
	hh := Human{"zhang3", "male"}
	hh.Eat()
	hh.Walk()

	s := Superman{Human{"li4", "female"}, 88}
	s.Walk()
	s.Eat()
	s.Fly()

}
