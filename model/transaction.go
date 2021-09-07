package model

import "fmt"

type Transaction struct {
	ID                  string
	Amount              int64
	CreatedAt           int64
	VendorTransactionID string
	Description         string
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"(%d) %s.%s : %d (%s)",
		t.CreatedAt,
		t.ID,
		t.VendorTransactionID,
		t.Amount,
		t.Description,
	)
}
