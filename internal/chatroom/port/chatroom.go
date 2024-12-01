package port

import "golipors/internal/chatroom/domain"

type ChatRoomRepository interface {
	CreateChatRoom(chatroom *domain.Chatroom) error
	GetChatRoomByID(id uint) (*domain.Chatroom, error)
}
