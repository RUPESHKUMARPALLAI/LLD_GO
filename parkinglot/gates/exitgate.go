package gates

import (
	"fmt"
	"parkinglot/billpayment"
	"parkinglot/slot"
	"parkinglot/ticket"
	"time"
)

type ExitGate struct {
	ID          string
	SlotManager *slot.SlotManager
}

func NewExitGate(id string, sm *slot.SlotManager) *ExitGate {
	return &ExitGate{
		ID:          id,
		SlotManager: sm,
	}
}

func (eg *ExitGate) ProcessExit(t *ticket.Ticket, time time.Time) (*billpayment.Bill, *billpayment.Payment, error) {
	t.Close(time)
	eg.SlotManager.ReleaseSlot(t.Slot.ID)
	bill := billpayment.GenerateBill(t)
	fmt.Printf("Bill Generated at ExitGate %s: BillID=%s, Amount=%d\n", eg.ID, bill.ID, bill.Amount)
	payment := billpayment.NewPayment(bill)
	payment.ProcessPayment(true)

	fmt.Printf("Payment Status: %s\n", payment.Status)
	if payment.Status != billpayment.Success {
		return bill, payment, fmt.Errorf("payment failed")
	}

	return bill, payment, nil
}