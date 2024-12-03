package message

import (
	"golipors/internal/message/domain"
	"golipors/internal/message/port"
)

type messageService struct {
	repo port.MessageRepository
}

func NewMessageService(repo port.MessageRepository) port.MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) SendMessage(message *domain.Message) error {
	return s.repo.CreateMessage(message)
}

func (s *messageService) GetChatroomMessages(chatroomID uint) ([]*domain.Message, error) {
	return s.repo.GetMessagesByChatroom(chatroomID)
}
