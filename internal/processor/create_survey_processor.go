package processor

import (
	"context"
	"encoding/json"

	"github.com/ffelipelimao/survey/internal/entities"
)

type SaveSurveyUseCase interface {
	Create(ctx context.Context, survey *entities.Survey) error
}

type SaveSurveyProcessor struct {
	saveSurveyUseCase SaveSurveyUseCase
}

func NewSaveSurveyProcessor(saveSurveyUseCase SaveSurveyUseCase) *SaveSurveyProcessor {
	return &SaveSurveyProcessor{
		saveSurveyUseCase: saveSurveyUseCase,
	}
}

func (sp *SaveSurveyProcessor) Handle(ctx context.Context, msg []byte) error {
	var survey entities.Survey
	err := json.Unmarshal(msg, &survey)
	if err != nil {
		return err
	}
	return sp.saveSurveyUseCase.Create(ctx, &survey)
}
