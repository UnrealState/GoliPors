package port

import "golipors/internal/transaction/domain"

type TransactionRepository interface {
	RecordTransaction(transaction *domain.Transaction) error
	GetTransactionsByUser(userID uint) ([]*domain.Transaction, error)
}
