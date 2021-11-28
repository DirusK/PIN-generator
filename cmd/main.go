package main

import (
	"flag"
	"fmt"
	"log"
	"pin-generation/pkg/visa"

	"pin-generation/internal/card"
	"pin-generation/internal/config"
	"pin-generation/pkg/ibm"
)

func main() {
	cfgPath := flag.String("c", config.DefaultPath, "configuration file")
	flag.Parse()

	state, err := config.New(*cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	creditCard := card.New()
	fmt.Println("Received a new credit card...")
	fmt.Println(creditCard)

	ibmGenerator := ibm.New(state.PINGenerationKey, state.Decimalization)
	pinOffset, err := ibmGenerator.GenerateOffsetPIN(creditCard.Number, creditCard.PIN)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("IBM-3624: for specified %s user's PIN is generated %s offset PIN \n", creditCard.PIN, pinOffset)

	visaGenerator := visa.New(state.PINGenerationKey, state.PinVerificationKeyIndex)
	PVV, err := visaGenerator.GeneratePVV(creditCard.Number, creditCard.PIN)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("VISA-PVV: for specified %s user's PIN is generated %s VISA PVV \n", creditCard.PIN, PVV)
}
