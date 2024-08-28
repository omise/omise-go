package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/examples/src"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, _ := omise.NewClient(
		os.Getenv("PUBLIC_KEY"),
		os.Getenv("SECRET_KEY"),
	)

	// client.SetDebug(true)

	/**
	 * Start: Create charge with QR
	 */
	chargeQR, err := src.CreateChargeQR(client)

	if err != nil {
		log.Println(err)
	}

	src.CustomLog(chargeQR)

	// Calling Sleep method
	time.Sleep(1 * time.Second)

	/**
	 * Start: Create charge with card
	 */
	chargeCard, err := src.CreateChargeCard(client)

	if err != nil {
		log.Println(err)
	}

	src.CustomLog(chargeCard)

	// Calling Sleep method
	time.Sleep(1 * time.Second)

	// /**
	//  * Start: Create recipient
	//  */
	recipient, err := src.CreateRecipient(client)

	if err != nil {
		log.Println(err)
	}

	src.CustomLog(recipient)

	// Calling Sleep method
	time.Sleep(1 * time.Second)

	// /**
	//  * Start: Create charge schedule
	//  */
	chargeSchedule, err := src.CreateChargeSchedule(client)

	if err != nil {
		log.Println(err)
	}

	src.CustomLog(chargeSchedule)

	// Calling Sleep method
	time.Sleep(1 * time.Second)

	// /**
	//  * Start: Create transfer schedule
	//  */
	transferSchedule, err := src.CreateTransferSchedule(client)

	if err != nil {
		log.Println(err)
	}

	src.CustomLog(transferSchedule)
}
