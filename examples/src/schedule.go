package src

import (
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/omise/omise-go/schedule"
)

func CreateChargeSchedule(client *omise.Client) (*omise.Schedule, error) {
	customer, err := CreateCustomer(client)

	if err != nil {
		return nil, err
	}

	schedule, createSchedule := &omise.Schedule{}, &operations.CreateChargeSchedule{
		Every:       2,
		Period:      schedule.PeriodDay,
		StartDate:   "2024-09-01",
		EndDate:     "2024-12-01",
		Customer:    customer.ID,
		Amount:      100000,
		Description: "Membership fee",
	}

	if err := client.Do(schedule, createSchedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

func CreateTransferSchedule(client *omise.Client) (*omise.Schedule, error) {
	recipient, err := CreateRecipient(client)

	if err != nil {
		return nil, err
	}

	schedule, createSchedule := &omise.Schedule{}, &operations.CreateTransferSchedule{
		Every:     2,
		Period:    schedule.PeriodDay,
		StartDate: "2024-09-01",
		EndDate:   "2024-12-01",
		Amount:    100000,
		Recipient: recipient.ID,
	}

	if err := client.Do(schedule, createSchedule); err != nil {
		return nil, err
	}

	return schedule, nil
}
