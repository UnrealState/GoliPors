package storage

import (
	"golipors/internal/chatroom/domain"
	"golipors/internal/chatroom/port"
	"gorm.io/gorm"
)

type chatRoomRepository struct {
	db *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) port.ChatRoomRepository {
	return &chatRoomRepository{db: db}
}

func (r *chatRoomRepository) CreateChatRoom(chatroom *domain.Chatroom) error {
	return r.db.Create(chatroom).Error
}

func (r *chatRoomRepository) GetChatRoomByID(id uint) (*domain.Chatroom, error) {
	var chatroom domain.Chatroom
	err := r.db.Preload("Messages").First(&chatroom, id).Error
	return &chatroom, err
}
