package models

import "time"

type Message struct {
	ID         uint `gorm:"primaryKey"`
	ChatroomID uint
	Chatroom   *Chatroom `gorm:"foreignKey:ChatroomID"`
	UserID     uint
	User       *User     `gorm:"foreignKey:UserID"`
	Content    string    `gorm:"type:text;not null"`
	CreatedAt  time.Time `gorm:"not null"`
}
