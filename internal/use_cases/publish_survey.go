package usecases

import (
	"encoding/json"

	"github.com/ffelipelimao/survey/internal/entities"
)

type SurveyPublisher interface {
	Publish(msg []byte) error
}

type PublisherSurveyUseCase struct {
	surveyPublisher SurveyPublisher
}

func NewPublisherSurveyUseCase(surveyPublisher SurveyPublisher) *PublisherSurveyUseCase {
	return &PublisherSurveyUseCase{
		surveyPublisher: surveyPublisher,
	}
}

func (cs *PublisherSurveyUseCase) Create(survey *entities.Survey) error {
	surveyMessage, err := json.Marshal(survey)
	if err != nil {
		return err
	}
	return cs.surveyPublisher.Publish(surveyMessage)
}
