package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "wooden window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "wooden door"
}

func (b *normalBuilder) setNumfloor() {
	b.floor = 2
}

func (h *normalBuilder) getHouse() House {
	return House{
		doorType:   h.doorType,
		windowType: h.windowType,
		floor:      h.floor,
	}
}
