package main

import (
	"fmt"

	"github.com/huytran2000-hcmus/design-patterns-in-go/creational/abstractfactory/sportswear"
)

func main() {
	adidasFactory, _ := sportswear.GetSportswearFactory(sportswear.AdidasType)
	nikeFactory, _ := sportswear.GetSportswearFactory(sportswear.NikeType)

	adidasShirt := adidasFactory.MakeShirt()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	nikeShoe := adidasFactory.MakeShoe()

	adidasShirt.SetColor(sportswear.Blue)
	adidasShirt.SetSize(5)
	nikeShirt.SetColor(sportswear.Red)
	nikeShirt.SetSize(6)

	adidasShoe.SetColor(sportswear.Blue)
	adidasShoe.SetSize(43)
	nikeShoe.SetColor(sportswear.Red)
	nikeShoe.SetSize(44)

	fmt.Println("Adidas")
	fmt.Println("====================")
	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
	fmt.Println()

	fmt.Println("Nike")
	fmt.Println("====================")
	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)
	fmt.Println()
}

func printShirtDetails(s sportswear.Shirt) {
	fmt.Println("Shirt:")
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShoeDetails(s sportswear.Shoe) {
	fmt.Println("Shoe:")
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
