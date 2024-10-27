package handlers

import (
	"net/http"

	"github.com/ffelipelimao/survey/internal/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PublisherSurveyUseCase interface {
	Create(survey *entities.Survey) error
}

type PublisherSurveyHandler struct {
	publisherSurveyUseCase PublisherSurveyUseCase
}

func NewPublisherSurveyHandler(publisherSurveyUseCase PublisherSurveyUseCase) *PublisherSurveyHandler {
	return &PublisherSurveyHandler{
		publisherSurveyUseCase: publisherSurveyUseCase,
	}
}

func (cs *PublisherSurveyHandler) Handle(c echo.Context) error {
	s := new(entities.Survey)
	if err := c.Bind(s); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	s.ID = uuid.NewString()

	err := cs.publisherSurveyUseCase.Create(s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}
