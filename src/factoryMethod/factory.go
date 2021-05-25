package factoryMethod

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
}

type PaymentType int

const (
	Cash PaymentType = iota
	DebitCard
)

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	fmt.Print("pay", amount)
	return ""
}

func (c *DebitCardPM) Pay(amount float32) string {
	fmt.Print("pay", amount)
	return ""
}

func GetPaymentMethod(t PaymentType) PaymentMethod {
	switch t {
	case Cash:
		return new(CashPM)
	case DebitCard:
		return new(DebitCardPM)
	}
	return nil
}
