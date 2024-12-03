package port

import "golipors/internal/chatroom/domain"

type ChatRoomService interface {
	CreateChatRoom(chatroom *domain.Chatroom) error
	GetChatRoomDetails(id uint) (*domain.Chatroom, error)
}
