package usecases

import (
	"context"

	"github.com/ffelipelimao/survey/internal/entities"
)

type SurveyRepository interface {
	Save(ctx context.Context, survey *entities.Survey) error
	Count(ctx context.Context, surveyID string) (int64, error)
	SaveAvg(ctx context.Context, surveyID string, avg float32) error
	GetAvgRating(ctx context.Context, merchantID string) (*entities.SurveyAvg, error)
	ListSurveys(ctx context.Context, merchantID string) ([]*entities.Survey, error)
}
