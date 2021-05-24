package builder

// Director ...
type Director struct {
	builder IBuilder
}

// NewDirector ...
func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	if d.builder != nil {
		d.builder.setDoorType()
		d.builder.setWindowType()
		d.builder.setNumfloor()
		return d.builder.getHouse()
	}
	return House{
		windowType: "s",
		doorType:   "s",
	}

}
