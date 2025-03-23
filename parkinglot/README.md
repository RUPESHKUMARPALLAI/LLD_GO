+-----------------------------+
|         Vehicle             |
+-----------------------------+
| - number: string            |
| - type: VehicleType (Enum)  |  ← car, bike, truck
+-----------------------------+

+-----------------------------+
|           Slot              |
+-----------------------------+
| - id: string                |
| - type: SlotType (Enum)     |  ← small (bike), large (car/truck)
| - isOccupied: bool          |
| - vehicle: Vehicle          |  <<HAS-A Vehicle>>
+-----------------------------+

+-----------------------------+
|        SlotManager          |  <<Singleton>>
+-----------------------------+
| - slots: []Slot             |
+-----------------------------+
| + addSlot(slot Slot)        |
| + removeSlot(id string)     |
| + assignSlot(vehicle Vehicle): Slot |
| + releaseSlot(id string)    |
+-----------------------------+

+-----------------------------+
|           Ticket            |
+-----------------------------+
| - id: string                |
| - vehicle: Vehicle          |  <<HAS-A Vehicle>>
| - slot: Slot                |  <<HAS-A Slot>>
| - entryTime: time.Time      |
| - exitTime: time.Time       |
| - status: TicketStatus (Enum) ← active, closed
+-----------------------------+

+-----------------------------+
|            Bill             |
+-----------------------------+
| - billId: string            |
| - ticket: Ticket            |  <<HAS-A Ticket>>
| - ratePerHour: float        |  ← Enum-driven by vehicle type
| - hoursParked: int          |
| - amount: float             |
+-----------------------------+

+-----------------------------+
|          Payment            |
+-----------------------------+
| - paymentId: string         |
| - bill: Bill                |  <<HAS-A Bill>>
| - status: PaymentStatus (Enum) ← success, fail
+-----------------------------+

+-----------------------------+            +------------------------------+
|         EntryGate           |            |         ExitGate             |
+-----------------------------+            +------------------------------+
| - id: string                |            | - id: string                 |
+-----------------------------+            +------------------------------+
| + createTicket(vehicle Vehicle): Ticket |  + processExit(ticket Ticket): Payment |
|   → Uses SlotManager                    |     → Uses SlotManager
|   → Assigns Slot                        |     → Releases Slot
|   → Creates Ticket                      |     → Closes Ticket
|                                         |     → Creates Bill
|                                         |     → Processes Payment
+-----------------------------+            +------------------------------+
