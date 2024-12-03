package port

import "golipors/internal/notification/domain"

type NotificationRepository interface {
	CreateNotification(notification *domain.Notification) error
	GetNotificationsByUser(userID uint) ([]*domain.Notification, error)
	GetNotificationByID(id uint) (*domain.Notification, error)
}
