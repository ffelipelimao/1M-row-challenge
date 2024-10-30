package handlers

import (
	"context"
	"net/http"

	"github.com/ffelipelimao/survey/internal/entities"
	"github.com/labstack/echo/v4"
)

type ListSurveyUseCase interface {
	List(ctx context.Context, merchantID string) ([]*entities.Survey, error)
}

type ListSurveyHandler struct {
	ListSurveyUseCase ListSurveyUseCase
}

func NewListSurveyHandler(ListSurveyUseCase ListSurveyUseCase) *ListSurveyHandler {
	return &ListSurveyHandler{
		ListSurveyUseCase: ListSurveyUseCase,
	}
}

func (cs *ListSurveyHandler) Handle(c echo.Context) error {
	merchantID := c.Param("merchant_id")

	surveys, err := cs.ListSurveyUseCase.List(context.Background(), merchantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, surveys)
}
