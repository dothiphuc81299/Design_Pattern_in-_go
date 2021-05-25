package main

import (
	"design-pattern-go-code/src/abstractFactory"
	"fmt"
)

func main() {
	// normalBuilder := builder.GetBuilder("normal")
	// iglooBuilder := builder.GetBuilder("igloo")

	// director := builder.NewDirector(normalBuilder)
	// normalHouse := director.BuildHouse()

	// fmt.Println("normal  house:%s", normalHouse.GetDoorType())

	// fmt.Println("normal  house:%s", normalHouse.GetWindowType())
	// fmt.Println("normal  house:%d", normalHouse.GetFloor())

	// director.SetBuilder(iglooBuilder)
	// iglooHouse := director.BuildHouse()

	// fmt.Println("igloo  house:%s", iglooHouse.GetDoorType())

	// fmt.Println("igloo  house:%s", iglooHouse.GetWindowType())
	// fmt.Println("igloo  house:%d", iglooHouse.GetFloor())

	//abstract factory
	adidasFactory := abstractFactory.GetSportsFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	printShoeDetails(adidasShoe)
	adidasShort := adidasFactory.MakeShort()
	printShortDetails(adidasShort)

}

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Println("Logo: %s", s.GetLogo())
	fmt.Println("size: %d", s.GetSize())
}

func printShortDetails(s abstractFactory.IShort) {
	fmt.Println("Logo: %s", s.GetLogo())
	fmt.Println("size: %d", s.GetSize())
}
