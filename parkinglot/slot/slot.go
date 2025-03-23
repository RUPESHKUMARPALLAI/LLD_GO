package slot

import "parkinglot/vehicle"

type SlotType string

const (
	TwoWheelerSlot  SlotType = "TwoWheelerSlot"
	FourWheelerSlot SlotType = "FourWheelerSlot"
)

type Slot struct {
	ID         string
	Type       SlotType
	IsOccupied bool
	Vehicle    *vehicle.Vehicle
}


func (s *Slot) AddVehicle(vehicle *vehicle.Vehicle) {
	s.Vehicle = vehicle
	s.IsOccupied = true
}

func (s *Slot) RemoveVehicle() {
	s.Vehicle = nil
	s.IsOccupied = false
}

func NewSlot(ID string, Type SlotType) *Slot {
	return &Slot{
		ID: ID,
		Type: Type,
		IsOccupied: false,
		Vehicle: nil,
	}
}

