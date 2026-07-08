package service

import (
	"context"
	"task-tracker/notification-service/internal/domain"
	"task-tracker/notification-service/internal/port/out"
)

type NotificationService struct {
	userContactRepo  out.UserContactRepository
	notificationRepo out.NotificationRepository
	provider         out.NotificationProvider
}

func NewNotificationService(userContactRepo out.UserContactRepository, notificationRepo out.NotificationRepository, provider out.NotificationProvider) *NotificationService {
	return &NotificationService{
		userContactRepo:  userContactRepo,
		notificationRepo: notificationRepo,
		provider:         provider,
	}
}

func (n NotificationService) Send(ctx context.Context, notification domain.Notification) error {
	err := n.provider.Send(ctx, notification)
	if err != nil {
		return err
	}

	_, err = n.notificationRepo.Save(ctx, notification)
	if err != nil {
		return err
	}

	return nil
}

func (n NotificationService) GetStats(ctx context.Context, filter domain.StatisticsFilter) ([]domain.Statistics, error) {
	response, err := n.notificationRepo.GetStatistics(ctx, filter)
	if err != nil {
		return nil, err
	}

	return response, err
}
