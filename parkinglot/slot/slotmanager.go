package slot

import (
	"errors"
	"parkinglot/vehicle"
)

type SlotManager struct {
	Slots []*Slot
}

func NewSlotManager() *SlotManager {
	return &SlotManager{
		Slots: []*Slot{},
	}
}

func (sm *SlotManager) AddSlot(slot *Slot) {
	sm.Slots = append(sm.Slots, slot)
}

func (sm *SlotManager) RemoveSlot(slotID string) {
	for i, s := range sm.Slots {
		if s.ID == slotID {
			sm.Slots = append(sm.Slots[:i], sm.Slots[i+1:]...)
			return
		}
	}
}

func (sm *SlotManager) AssignSlot(v *vehicle.Vehicle) (*Slot, error) {
	var requiredType SlotType
	switch v.Type {
	case vehicle.Bike:
		requiredType = TwoWheelerSlot
	case vehicle.Car, vehicle.Truck:
		requiredType = FourWheelerSlot	
	}

	for _, s := range sm.Slots {
		if !s.IsOccupied && s.Type == requiredType {
			s.AddVehicle(v)
			return s, nil
		}
	}
	return nil, errors.New("no slots are available")
}

func (sm *SlotManager) ReleaseSlot(slotID string) {
	for _,s := range sm.Slots {
		if s.ID == slotID {
			s.RemoveVehicle()
			return
		}
	}
}