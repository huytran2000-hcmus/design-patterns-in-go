package constructor

type Builder interface {
	setFloors(floors int)
	setDoorType(doorType DoorType)
	setWindowType(windowType WindowType)
}

type DoorType string

const (
	woodenDoor DoorType = "wooden"
	ironDoor   DoorType = "iron"
)

type WindowType string

const (
	woodenWindow WindowType = "wooden"
	ironWindow   WindowType = "iron"
)

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) MakeNormalHouse() {
	d.builder.setFloors(2)
	d.builder.setDoorType(woodenDoor)
	d.builder.setWindowType(woodenWindow)
}

func (d *Director) ChangeBuilder(builder Builder) {
	d.builder = builder
}
