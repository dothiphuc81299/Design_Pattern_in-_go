package main

import (
	"design-pattern-go-code/src/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Println("normal  house:%s", normalHouse.GetDoorType())

	fmt.Println("normal  house:%s", normalHouse.GetWindowType())
	fmt.Println("normal  house:%d", normalHouse.GetFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Println("igloo  house:%s", iglooHouse.GetDoorType())

	fmt.Println("igloo  house:%s", iglooHouse.GetWindowType())
	fmt.Println("igloo  house:%d", iglooHouse.GetFloor())
}
