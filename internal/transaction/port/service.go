package port

import "golipors/internal/transaction/domain"

type TransactionService interface {
	CreateTransaction(transaction *domain.Transaction) error
	ListTransactionsByUser(userID uint) ([]*domain.Transaction, error)
}
