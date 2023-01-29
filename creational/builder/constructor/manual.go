package constructor

import "fmt"

type ManualBuilder struct {
	floors     int
	doorType   DoorType
	windowType WindowType
}

func NewManualBuilder() *ManualBuilder {
	return &ManualBuilder{}
}

func (mb *ManualBuilder) GetManual() *Manual {
	return &Manual{
		floors:     mb.floors,
		doorType:   mb.doorType,
		windowType: mb.windowType,
	}
}

type Manual struct {
	floors     int
	doorType   DoorType
	windowType WindowType
}

func (m *Manual) String() string {
	return fmt.Sprintf("manual of a house which has %d floors, with %s door and %s window", m.floors, m.doorType, m.windowType)
}

func (mb *ManualBuilder) setFloors(floors int) {
	mb.floors = floors
}

func (mb *ManualBuilder) setDoorType(doorType DoorType) {
	mb.doorType = doorType
}

func (mb *ManualBuilder) setWindowType(windowType WindowType) {
	mb.windowType = windowType
}
