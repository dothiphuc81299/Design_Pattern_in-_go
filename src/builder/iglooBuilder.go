package builder

// gui lai cac trang thai giong House
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "snow window"

}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "snow door"
}

func (b *iglooBuilder) setNumfloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

