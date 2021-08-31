package repository

import "app/model"

type TransactionRepository interface {
	Save(model.Transaction) error
}
