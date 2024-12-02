package storage

import (
	"golipors/internal/notification/domain"
	"golipors/internal/notification/port"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) port.NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateNotification(notification *domain.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotificationsByUser(userID uint) ([]*domain.Notification, error) {
	var notifications []*domain.Notification
	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) GetNotificationByID(id uint) (*domain.Notification, error) {
	var notification domain.Notification
	err := r.db.First(&notification, id).Error
	return &notification, err
}
