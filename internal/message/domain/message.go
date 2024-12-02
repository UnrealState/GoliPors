package domain

import (
	chatRoomDomain "golipors/internal/chatroom/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type MessageID uint

type Message struct {
	ID         MessageID
	ChatroomID chatRoomDomain.ChatroomID
	Chatroom   *chatRoomDomain.Chatroom
	UserID     userDomain.UserID
	User       *userDomain.User
	Content    string
	CreatedAt  time.Time
}
