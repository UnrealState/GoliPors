package storage

import (
	"golipors/internal/message/domain"
	"golipors/internal/message/port"
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) port.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message *domain.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) GetMessagesByChatroom(chatroomID uint) ([]*domain.Message, error) {
	var messages []*domain.Message
	err := r.db.Where("chatroom_id = ?", chatroomID).Order("created_at").Find(&messages).Error
	return messages, err
}
