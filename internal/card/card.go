package card

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	PINLength    = 4
	NumberLength = 16
)

type CreditCard struct {
	Type    string
	Number  string
	Expires string
	CVV     string
	PIN     string
}

func New() *CreditCard {
	return &CreditCard{
		Type:    gofakeit.CreditCardType(),
		Number:  gofakeit.DigitN(NumberLength),
		Expires: gofakeit.CreditCardExp(),
		CVV:     gofakeit.CreditCardCvv(),
		PIN:     gofakeit.DigitN(PINLength),
	}
}

func (c CreditCard) String() string {
	return fmt.Sprintf(
		"----------------------\n"+
			"Type: %s\nNumber: %s\nExpires: %s\nCVV: %s\nPIN: %s\n"+
			"----------------------",
		c.Type, c.Number, c.Expires, c.CVV, c.PIN)
}
