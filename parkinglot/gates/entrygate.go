package gates

import (
	"fmt"
	"parkinglot/slot"
	"parkinglot/ticket"
	"parkinglot/vehicle"
)

type EntryGate struct {
	ID          string
	SlotManager *slot.SlotManager
}

func (eg *EntryGate) CreateTicket(v *vehicle.Vehicle) (*ticket.Ticket,error) {
	slot, err := eg.SlotManager.AssignSlot(v)
	if err!=nil {
		return nil, err
	}
	t := ticket.NewTicket(v, slot)
	fmt.Printf("Ticket id: %s entry time: %s\n", t.ID, t.EntryTime)
	return t, nil
}

func NewEntryGate(id string, sm *slot.SlotManager) *EntryGate {
	return &EntryGate{
		ID:          id,
		SlotManager: sm,
	}
}

