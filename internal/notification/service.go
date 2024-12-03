package notification

import (
	"errors"
	"golipors/internal/notification/domain"
	"golipors/internal/notification/port"
)

type notificationService struct {
	repo port.NotificationRepository
}

func NewNotificationService(repo port.NotificationRepository) port.NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) SendNotification(notification *domain.Notification) error {
	if notification.UserID == 0 {
		return errors.New("user ID is required")
	}

	if notification.Message == "" {
		return errors.New("notification message cannot be empty")
	}

	return s.repo.CreateNotification(notification)
}

func (s *notificationService) GetUserNotifications(userID uint) ([]*domain.Notification, error) {
	return s.repo.GetNotificationsByUser(userID)
}

func (s *notificationService) GetNotificationDetails(id uint) (*domain.Notification, error) {
	return s.repo.GetNotificationByID(id)
}
