package ticket

import (
	"parkinglot/slot"
	"parkinglot/vehicle"
	"time"
	"github.com/google/uuid"
)

type TicketStatus string

const(
	Active TicketStatus = "Active"
	Closed TicketStatus = "Closed"
)

type Ticket struct {
	ID      string
	Vehicle *vehicle.Vehicle
	Slot *slot.Slot
	EntryTime time.Time
	ExitTime *time.Time
	Status TicketStatus
}

func (t *Ticket) Close(time time.Time) {
	now := time
	t.ExitTime = &now
	t.Status = Closed
}

func NewTicket(v *vehicle.Vehicle, s *slot.Slot) *Ticket {
	return &Ticket{
		ID:        uuid.New().String(),
		Vehicle:   v,
		Slot:      s,
		EntryTime: time.Now(),
		Status:    Active,
	}
}

