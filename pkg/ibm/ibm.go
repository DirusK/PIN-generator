package ibm

import (
	"pin-generation/pkg/cipher"
	"strconv"
)

type Ibm struct {
	PINGenerationKey string
	Decimalization   map[string]string
}

func New(PINGenerationKey string, decimalization map[string]string) *Ibm {
	return &Ibm{
		PINGenerationKey: PINGenerationKey,
		Decimalization:   decimalization,
	}
}

// GenerateOffsetPIN returns pin offset of specified params.
func (ibm Ibm) GenerateOffsetPIN(cardNumber string, userPIN string) (string, error) {
	ciphertext, err := cipher.TripleDesEncrypt(cardNumber[len(cardNumber)-16:], ibm.PINGenerationKey)
	if err != nil {
		return "", err
	}

	var replaced string
	for _, encrypted := range ciphertext {
		replaced += ibm.Decimalization[string(encrypted)]
	}

	var (
		i = 1
		j = 0
	)

	var offsetPIN string
	for i < len(userPIN)+1 {
		naturalPin, err := strconv.ParseInt(replaced[j:i], 10, 32)
		if err != nil {
			return "", err
		}

		customerPin, err := strconv.ParseInt(userPIN[j:i], 10, 32)
		if err != nil {
			return "", err
		}

		sub := (customerPin - naturalPin + 10) % 10
		offsetPIN += strconv.Itoa(int(sub))

		i++
		j++
	}

	return offsetPIN, nil
}
