package ticket

import (
	"time"
)

type Ticket struct {
	Type		string
	Company		string
	Summary		string
	Priority	int
	Entered		time.Time
}

func CreateTicket(ttype string, company string, summary string, priority int) *Ticket {
	return &Ticket{
		Type: ttype,
		Company: company,
		Summary: summary,
		Priority: priority,
		Entered: time.Now(),
	}
}
