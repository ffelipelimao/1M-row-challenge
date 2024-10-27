package repository

import (
	"context"
	"errors"

	"github.com/ffelipelimao/survey/internal/entities"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SurveyPostgresRepository struct {
	db *pgxpool.Pool
}

func NewSurveyPostgresRepository(db *pgxpool.Pool) *SurveyPostgresRepository {
	return &SurveyPostgresRepository{db: db}
}

func (r *SurveyPostgresRepository) Save(ctx context.Context, survey *entities.Survey) error {
	query := `
		INSERT INTO surveys (id, merchant_id, user_id, rating, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err := r.db.Exec(ctx, query, survey.ID, survey.MerchantID, survey.UserID, survey.Rating)
	return err
}

func (r *SurveyPostgresRepository) Count(ctx context.Context, merchantID string) (int64, error) {
	var count int64
	query := `
		SELECT COUNT(*)
		FROM surveys
		WHERE merchant_id = $1
	`
	err := r.db.QueryRow(ctx, query, merchantID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *SurveyPostgresRepository) SaveAvg(ctx context.Context, merchantID string, avg float32) error {
	query := `
		INSERT INTO merchant_avg_ratings (merchant_id, average_rating)
		VALUES ($1, $2)
		ON CONFLICT (merchant_id) DO UPDATE
		SET average_rating = $2
	`
	commandTag, err := r.db.Exec(ctx, query, merchantID, avg)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("nenhuma média foi atualizada")
	}
	return nil
}
