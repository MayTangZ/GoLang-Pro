package main

import "fmt"

type Hero struct {
	Name string
	Ad   int
	//Level
}

/*
func (this Hero) show() {
	fmt.Println("hero = ", this)
}

func (this Hero) GetName() string {

	return this.Name

}

func (this Hero) SetName(newName string) {
	this.Name = newName
}

*/

func (this *Hero) show() {
	fmt.Println("hero = ", this)
}

func (this *Hero) GetName() string {

	return this.Name

}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main2() {
	hero := Hero{Name: "zhang3", Ad: 100}
	hero.show()
	hero.SetName("Li4")
	hero.show()
}
