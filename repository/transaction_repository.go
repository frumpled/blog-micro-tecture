package repository

import "app/repository/model"

type TransactionRepository interface {
	Save(model.Transaction) error
}
