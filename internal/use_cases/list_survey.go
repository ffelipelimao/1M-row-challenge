package usecases

import (
	"context"

	"github.com/ffelipelimao/survey/internal/entities"
)

type ListSurveysUseCase struct {
	surveyRepository SurveyRepository
}

func NewListSurveysUseCase(surveyRepository SurveyRepository) *ListSurveysUseCase {
	return &ListSurveysUseCase{
		surveyRepository: surveyRepository,
	}
}

func (cs *ListSurveysUseCase) List(ctx context.Context, merchantID string) ([]*entities.Survey, error) {
	surveys, err := cs.surveyRepository.ListSurveys(ctx, merchantID)
	if err != nil {
		return nil, err
	}

	if surveys == nil {
		return []*entities.Survey{
			{},
		}, nil
	}

	return surveys, nil
}
