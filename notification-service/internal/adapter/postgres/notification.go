package postgres

import (
	"context"
	"task-tracker/notification-service/internal/db"
	"task-tracker/notification-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type NotificationRepository struct {
	queries *db.Queries
}

func NewNotificationRepository(queries *db.Queries) *NotificationRepository {
	return &NotificationRepository{queries: queries}
}

func (n NotificationRepository) Save(ctx context.Context, notification domain.Notification) (uuid.UUID, error) {
	params := db.SaveNotificationParams{
		UserID:   notification.UserID,
		Type:     notification.Type,
		Provider: notification.Provider,
	}

	row, err := n.queries.SaveNotification(ctx, params)
	if err != nil {
		return uuid.UUID{}, err
	}

	return row.ID, nil
}

func (n NotificationRepository) GetStatistics(ctx context.Context, filter domain.StatisticsFilter) ([]domain.Statistics, error) {
	params := db.GetStatisticsParams{
		Provider: pgtype.Text{},
		Date:     pgtype.Timestamptz{},
	}

	if filter.Provider != nil {
		params.Provider = pgtype.Text{
			String: *filter.Provider,
			Valid:  true,
		}
	}

	if filter.Date != nil {
		params.Date = pgtype.Timestamptz{
			Time:  *filter.Date,
			Valid: true,
		}
	}

	row, err := n.queries.GetStatistics(ctx, params)
	if err != nil {
		return nil, err
	}

	stats := make([]domain.Statistics, 0, len(row))
	for idx := 0; idx < len(row); idx++ {
		stats = append(stats, domain.Statistics{
			Type:       row[idx].Type,
			Provider:   row[idx].Provider,
			Date:       row[idx].Date.Time,
			UniqueHits: int(row[idx].UniqueHits),
			TotalHits:  int(row[idx].TotalHits),
		})
	}

	return stats, nil
}
