package handlers

import (
	"context"
	"net/http"

	"github.com/ffelipelimao/survey/internal/entities"
	"github.com/labstack/echo/v4"
)

type GetSurveyAvgUseCase interface {
	Get(ctx context.Context, merchantID string) (*entities.SurveyAvg, error)
}

type GetSurveyAvgHandler struct {
	GetSurveyAvgUseCase GetSurveyAvgUseCase
}

func NewGetSurveyAvgHandler(GetSurveyAvgUseCase GetSurveyAvgUseCase) *GetSurveyAvgHandler {
	return &GetSurveyAvgHandler{
		GetSurveyAvgUseCase: GetSurveyAvgUseCase,
	}
}

func (cs *GetSurveyAvgHandler) Handle(c echo.Context) error {
	merchantID := c.Param("merchant_id")

	surveyAvg, err := cs.GetSurveyAvgUseCase.Get(context.Background(), merchantID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if surveyAvg == nil {
		return c.JSON(http.StatusNotFound, "not found")
	}

	return c.JSON(http.StatusOK, surveyAvg)
}
