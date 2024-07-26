package main

import "fmt"

//本质是一个指针
type AnimanIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	corlor string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")

}

func (this *Cat) GetColor() string {
	return this.corlor
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	corlor string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")

}

func (this *Dog) GetColor() string {
	return this.corlor
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimanIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())

}

func main() {
	/*
		var animal AnimanIF
		animal = &Cat{"red"}
		animal.Sleep()

		animal = &Dog{"yellow"}
		animal.Sleep()
	*/
	cat := Cat{"red"}
	dog := Dog{"yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

}
