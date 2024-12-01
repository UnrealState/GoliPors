package domain

import (
	chatRoomDomain "golipors/internal/chatRoom/domain"
	userDomain "golipors/internal/user/domain"
	"time"
)

type Message struct {
	ID         uint `gorm:"primaryKey"`
	ChatroomID uint
	Chatroom   *chatRoomDomain.Chatroom `gorm:"foreignKey:ChatroomID"`
	UserID     uint
	User       *userDomain.User `gorm:"foreignKey:UserID"`
	Content    string           `gorm:"type:text;not null"`
	CreatedAt  time.Time        `gorm:"not null"`
}
