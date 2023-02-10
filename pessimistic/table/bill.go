package table

import (
	"time"

	_ "github.com/lib/pq"
)

type Bill struct {
	ID     int
	PaidAt *time.Time
}
