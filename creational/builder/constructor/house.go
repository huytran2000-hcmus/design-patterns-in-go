package constructor

import "fmt"

type HouseBuilder struct {
	floors     int
	doorType   DoorType
	windowType WindowType
}

type House struct {
	floors     int
	doorType   DoorType
	windowType WindowType
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{}
}

func (hb *HouseBuilder) GetHouse() *House {
	return &House{
		floors:     hb.floors,
		doorType:   hb.doorType,
		windowType: hb.windowType,
	}
}

func (h *House) String() string {
	return fmt.Sprintf(
		"house has %d floors, with %s door and %s window",
		h.floors,
		h.doorType,
		h.windowType,
	)
}

func (hb *HouseBuilder) setFloors(floors int) {
	hb.floors = floors
}

func (hb *HouseBuilder) setDoorType(doorType DoorType) {
	hb.doorType = doorType
}

func (hb *HouseBuilder) setWindowType(windowType WindowType) {
	hb.windowType = windowType
}
