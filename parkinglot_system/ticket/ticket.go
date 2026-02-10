package ticket

import "time"

type Ticket struct {
	VehidleId int
	EntryTime time.Time
	ExitTime  time.Time
	Amount    int
}
