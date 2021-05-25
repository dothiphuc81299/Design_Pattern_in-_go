package abstractFactory

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

// GetSportsFactory ..
func GetSportsFactory(brand string) ISportFactory {
	switch brand {
	case "adidas":
		return &Adidas{}
	case "nike":
		return &Nike{}
	}
	return nil
}
