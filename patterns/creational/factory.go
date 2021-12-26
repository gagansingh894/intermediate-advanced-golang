package creational

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}

}

type CashPM struct {}
type DebitCardPM struct {}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

func FactoryExample() {
	// Cash example
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		fmt.Println("A payment method of type Cash must exist")
	} else {
		msg := payment.Pay(5.0)
		fmt.Println(msg)
	}

	// Debit card example
	payment, err = GetPaymentMethod(DebitCard)
	if err != nil {
		fmt.Println("A payment method of type DebitCard must exist")
	} else {
		msg := payment.Pay(5.0)
		fmt.Println(msg)
	}

	// Unknown type
	payment, err = GetPaymentMethod(20)
	if err != nil {
		fmt.Println(err)
	} else {
		msg := payment.Pay(5.0)
		fmt.Println(msg)
	}


}
