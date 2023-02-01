package main

import (
	"fmt"

	"github.com/huytran2000-hcmus/design-patterns-in-go/creational/builder/constructor"
)

func main() {
	houseBuilder := constructor.NewHouseBuilder()
	director := constructor.NewDirector(houseBuilder)

	director.MakeNormalHouse()
	house := houseBuilder.GetHouse()
	fmt.Println(house)

	manualBuilder := constructor.NewManualBuilder()
	director.ChangeBuilder(manualBuilder)
	director.MakeNormalHouse()
	manual := manualBuilder.GetManual()
	fmt.Println(manual)
}
