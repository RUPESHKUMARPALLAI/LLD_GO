package billpayment

import (
	"math"
	"parkinglot/ticket"
	"parkinglot/vehicle"
	"time"

	"github.com/google/uuid"
)

type RatePerHour int

const (
	TwoWheelerRate  RatePerHour = 10
	FourWheelerRate RatePerHour = 20
)

type Bill struct {
	ID          string
	Ticket      *ticket.Ticket
	Rate        RatePerHour
	HoursParked int
	Amount      int
	GeneratedAt time.Time
}

func GenerateBill(t *ticket.Ticket) *Bill {
	var rate RatePerHour
	switch t.Vehicle.Type {
	case vehicle.Bike:
		rate = TwoWheelerRate
	case vehicle.Truck, vehicle.Car:
		rate = FourWheelerRate
	}
	hours := HoursParked(t.EntryTime, *t.ExitTime)
	Amount := int(rate) * hours
	return &Bill{
		ID:          uuid.New().String(),
		Ticket:      t,
		Rate:        rate,
		HoursParked: hours,
		Amount:      Amount,
		GeneratedAt: time.Now(),
	}
}

func HoursParked(entry, exit time.Time) int {
	duration := exit.Sub(entry)
	return (int(math.Ceil(duration.Hours())))
}
