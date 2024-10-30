package usecases

import (
	"context"

	"github.com/ffelipelimao/survey/internal/entities"
)

type GetSurveyAvgUseCase struct {
	surveyRepository SurveyRepository
}

func NewGetSurveyAvgUseCase(surveyRepository SurveyRepository) *GetSurveyAvgUseCase {
	return &GetSurveyAvgUseCase{
		surveyRepository: surveyRepository,
	}
}

func (cs *GetSurveyAvgUseCase) Get(ctx context.Context, merchantID string) (*entities.SurveyAvg, error) {
	surveyAvg, err := cs.surveyRepository.GetAvgRating(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	return surveyAvg, nil
}
