package table

type Bill struct {
	ID int
}

type Payment struct {
	ID     int
	BillID int
}
