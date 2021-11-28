package visa

import (
	"pin-generation/pkg/cipher"
	"unicode"
)

const PVVLength = 4

type Visa struct {
	PINGenerationKey        string
	PINVerificationKeyIndex string
}

func New(PINGenerationKey string, PINVerificationKeyIndex string) *Visa {
	return &Visa{
		PINGenerationKey:        PINGenerationKey,
		PINVerificationKeyIndex: PINVerificationKeyIndex,
	}
}

func (v Visa) GeneratePVV(cardNumber string, userPIN string) (string, error) {
	var TSP string
	TSP += cardNumber[len(cardNumber)-12:len(cardNumber)-1] + v.PINVerificationKeyIndex + userPIN

	ciphertext, err := cipher.TripleDesEncrypt(TSP, v.PINGenerationKey)
	if err != nil {
		return "", err
	}

	var PVV string
	for _, symbol := range ciphertext {
		if unicode.IsNumber(symbol) {
			PVV += string(symbol)
		}
	}

	for _, symbol := range ciphertext {
		if unicode.IsLetter(symbol) {
			PVV += string(symbol)
		}
	}

	return PVV[:PVVLength], nil
}
