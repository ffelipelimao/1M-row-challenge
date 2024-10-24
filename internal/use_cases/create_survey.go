package usecases

import (
	"encoding/json"

	"github.com/ffelipelimao/survey/internal/entities"
)

type SurveyPublisher interface {
	Publish(msg []byte) error
}

type CreateSurveyUseCase struct {
	surveyPublisher SurveyPublisher
}

func NewCreateSurveyUseCase(surveyPublisher SurveyPublisher) *CreateSurveyUseCase {
	return &CreateSurveyUseCase{
		surveyPublisher: surveyPublisher,
	}
}

func (cs *CreateSurveyUseCase) Create(survey *entities.Survey) error {
	surveyMessage, err := json.Marshal(survey)
	if err != nil {
		return err
	}
	return cs.surveyPublisher.Publish(surveyMessage)
}
