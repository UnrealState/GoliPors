package port

import "golipors/internal/message/domain"

type MessageRepository interface {
	CreateMessage(message *domain.Message) error
	GetMessagesByChatroom(chatroomID uint) ([]*domain.Message, error)
}
