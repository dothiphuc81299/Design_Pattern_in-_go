package builder

// IBuilder
type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumfloor()
	getHouse() House
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloor":
		return &iglooBuilder{}
	}
	return nil
}
