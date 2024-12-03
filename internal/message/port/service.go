package port

import "golipors/internal/message/domain"

type MessageService interface {
	SendMessage(message *domain.Message) error
	GetChatroomMessages(chatroomID uint) ([]*domain.Message, error)
}
