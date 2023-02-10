package table

import (
	"time"
)

type Bill struct {
	ID     int
	PaidAt *time.Time
}
