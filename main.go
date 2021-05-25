package main

import (
	"design-pattern-go-code/src/abstractFactory"
	p "design-pattern-go-code/src/prototype"
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

	// prototype
	file1 := &p.File{Name: "File 1"}
	file2 := &p.File{Name: "File 2"}
	file3 := &p.File{Name: "File 3"}
	folder1 := &p.Folder{
		Childrens: []p.INode{file1},
		Name:      "Folder 1",
	}

	folder2 := &p.Folder{
		Childrens: []p.INode{folder1, file2, file3},
		Name:      "folder 2",
	}

	fmt.Println("Printing for folder 2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("Printing for clone folder 2")
	cloneFolder.Print(" ")

}

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Println("Logo: %s", s.GetLogo())
	fmt.Println("size: %d", s.GetSize())
}

func printShortDetails(s abstractFactory.IShort) {
	fmt.Println("Logo: %s", s.GetLogo())
	fmt.Println("size: %d", s.GetSize())
}
