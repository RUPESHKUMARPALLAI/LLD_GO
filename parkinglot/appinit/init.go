package appinit

import (
	"fmt"
	"parkinglot/gates"
	"parkinglot/slot"
	"parkinglot/vehicle"
	"time"
)

func Demo() {
	sm := slot.NewSlotManager()

	slot1 := slot.NewSlot("1", slot.FourWheelerSlot)
	slot2 := slot.NewSlot("2", slot.FourWheelerSlot)
	slot3 := slot.NewSlot("3", slot.TwoWheelerSlot)

	sm.AddSlot(slot1)
	sm.AddSlot(slot2)
	sm.AddSlot(slot3)

	fmt.Println("Slots Initialized.")

	entryGate := gates.NewEntryGate("E1", sm)
	exitGate := gates.NewExitGate("X1", sm)

	vehicle1 := vehicle.NewVehicle("KA01AB1234", vehicle.Bike)
	vehicle2 := vehicle.NewVehicle("KA02CD5678", vehicle.Car)

	t1, err := entryGate.CreateTicket(vehicle1)
	if err != nil {
		fmt.Printf("Entry Error: %v\n", err)
		return
	}

	fmt.Printf("Vehicle %s entered. Ticket ID: %s, EntryTime: %v\n", vehicle1.Number, t1.ID, t1.EntryTime)

	t2, err := entryGate.CreateTicket(vehicle2)
	if err != nil {
		fmt.Printf("Entry Error: %v\n", err)
		return
	}

	fmt.Printf("Vehicle %s entered. Ticket ID: %s, EntryTime: %v\n", vehicle2.Number, t2.ID, t2.EntryTime)

	simulatedExitTime1 := t1.EntryTime.Add(2 * time.Hour)
	simulatedExitTime2 := t2.EntryTime.Add(2*time.Hour + 30*time.Minute)

	fmt.Printf("simulatedExitTime1: %v", simulatedExitTime1)
	fmt.Printf("simulatedExitTime2: %v", simulatedExitTime2)

	exitGate.ProcessExit(t1, simulatedExitTime1)
	exitGate.ProcessExit(t2, simulatedExitTime2)

}
