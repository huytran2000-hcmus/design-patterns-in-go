package main

import (
	"fmt"

	"github.com/huytran2000-hcmus/design-pattern-in-golang/structural/adapter/daikin"
	"github.com/huytran2000-hcmus/design-pattern-in-golang/structural/adapter/daikin/adapter"
	"github.com/huytran2000-hcmus/design-pattern-in-golang/structural/adapter/panasonic"
)

func main() {
	airCondition := panasonic.NewAirCondition()
	oldRemote := panasonic.NewACRemote(airCondition)
	fmt.Println("--------Old Panasonic remote-------")
	panasonic.TestAC(oldRemote)

	panasonicAdapter := adapter.NewPanasonicAC(airCondition)
	newRemote := daikin.NewACRemote(panasonicAdapter)
	fmt.Println("--------New Daikin remote----------")
	daikin.TestAC(newRemote)
}
