package email

import (
	"context"
	"task-tracker/notification-service/internal/domain"

	"github.com/resend/resend-go/v3"
)

type NotificationProvider struct {
	client      *resend.Client
	senderEmail string
}

func NewNotificationProvider(token, senderEmail string) *NotificationProvider {
	return &NotificationProvider{
		client:      resend.NewClient(token),
		senderEmail: senderEmail,
	}
}

func (n NotificationProvider) Send(context context.Context, notification domain.Notification) error {
	email := ""
	request := resend.SendEmailRequest{
		From:        n.senderEmail,
		To:          []string{email},
		Text:        notification.NotificationBody.Text,
		Attachments: photoUrlsToAttachments(notification.NotificationBody.PhotoUrls),
	}

	_, err := n.client.Emails.Send(&request)
	if err != nil {
		return err
	}

	return nil
}

func photoUrlsToAttachments(Urls []string) []*resend.Attachment {
	attachments := make([]*resend.Attachment, 0, len(Urls))

	for _, url := range Urls {
		attachment := resend.Attachment{
			Path: url,
		}
		attachments = append(attachments, &attachment)
	}

	return attachments
}
