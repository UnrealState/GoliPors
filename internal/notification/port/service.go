package port

import "golipors/internal/notification/domain"

type NotificationService interface {
	SendNotification(notification *domain.Notification) error
	GetUserNotifications(userID uint) ([]*domain.Notification, error)
	GetNotificationDetails(id uint) (*domain.Notification, error)
}
