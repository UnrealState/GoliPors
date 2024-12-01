package domain

import (
	chatRoomDomain "golipors/internal/chatroom/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Message struct {
	ID         uint
	ChatroomID uint
	Chatroom   *chatRoomDomain.Chatroom
	UserID     uint
	User       *userDomain.User
	Content    string
	CreatedAt  time.Time
}
