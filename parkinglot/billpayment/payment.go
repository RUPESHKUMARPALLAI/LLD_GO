package billpayment

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus string

const (
	Success PaymentStatus = "Success"
	Failure PaymentStatus = "Failure"
	Pending PaymentStatus = "Pending"
)

type Payment struct {
	ID     string
	Bill   *Bill
	Status PaymentStatus
	PaidAt time.Time
}

func NewPayment(bill *Bill) *Payment {
	return &Payment{
		ID:     uuid.New().String(),
		Bill:   bill,
		Status: Pending,
	}
}

func (p *Payment) ProcessPayment(success bool) {
	if success {
		p.Status = Success
		p.PaidAt = time.Now()
	} else {
		p.Status = Failure
		p.PaidAt = time.Now()
	}
}