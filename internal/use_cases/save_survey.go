package usecases

import (
	"context"

	"github.com/ffelipelimao/survey/internal/entities"
)

type SurveyRepository interface {
	Save(ctx context.Context, survey *entities.Survey) error
	Count(ctx context.Context, surveyID string) (int64, error)
	SaveAvg(ctx context.Context, surveyID string, avg float32) error
}

type SaveSurveyUseCase struct {
	surveyRepository SurveyRepository
}

func NewSaveSurveyUseCase(surveyRepository SurveyRepository) *SaveSurveyUseCase {
	return &SaveSurveyUseCase{
		surveyRepository: surveyRepository,
	}
}

func (cs *SaveSurveyUseCase) Create(ctx context.Context, survey *entities.Survey) error {
	if err := cs.surveyRepository.Save(ctx, survey); err != nil {
		return err
	}

	totalSurveys, err := cs.surveyRepository.Count(ctx, survey.ID)
	if err != nil {
		return err
	}

	if totalSurveys == 0 {
		if err := cs.surveyRepository.SaveAvg(ctx, survey.ID, survey.Rating); err != nil {
			return err
		}
	}

	average := (survey.Rating + (average * float32(totalSurveys-1))) / float32(totalSurveys)

	if err := cs.surveyRepository.SaveAvg(ctx, survey.ID, average); err != nil {
		return err
	}

	return nil
}
