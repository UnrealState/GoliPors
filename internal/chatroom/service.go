package chatroom

import (
	"golipors/internal/chatroom/domain"
	"golipors/internal/chatroom/port"
)

type chatRoomService struct {
	repo port.ChatRoomRepository
}

func NewChatRoomService(repo port.ChatRoomRepository) port.ChatRoomService {
	return &chatRoomService{repo: repo}
}

func (s *chatRoomService) CreateChatRoom(chatroom *domain.Chatroom) error {
	return s.repo.CreateChatRoom(chatroom)
}

func (s *chatRoomService) GetChatRoomDetails(id uint) (*domain.Chatroom, error) {
	return s.repo.GetChatRoomByID(id)
}
