package handlers

import (
	"net/http"

	"github.com/ffelipelimao/survey/internal/entities"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateSurveyUseCase interface {
	Create(survey *entities.Survey) error
}

type CreateHandlerSurvey struct {
	createSurveyUseCase CreateSurveyUseCase
}

func NewCreateSurveyHandler(createSurveyUseCase CreateSurveyUseCase) *CreateHandlerSurvey {
	return &CreateHandlerSurvey{
		createSurveyUseCase: createSurveyUseCase,
	}
}

func (cs *CreateHandlerSurvey) Handle(c echo.Context) error {
	s := new(entities.Survey)
	if err := c.Bind(s); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	s.ID = uuid.NewString()

	err := cs.createSurveyUseCase.Create(s)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}
