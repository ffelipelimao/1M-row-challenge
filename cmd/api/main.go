package main

import (
	"log"

	"github.com/ffelipelimao/survey/internal/handlers"
	"github.com/ffelipelimao/survey/internal/publisher"
	usecases "github.com/ffelipelimao/survey/internal/use_cases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	publisherSurvey, err := publisher.New("tp.create-survey")
	if err != nil {
		log.Fatal(err)
	}

	defer publisherSurvey.Stop()

	createSurveyUseCase := usecases.NewCreateSurveyUseCase(publisherSurvey)
	createSurveyHandler := handlers.NewCreateSurveyHandler(createSurveyUseCase)

	// Routes
	e.POST("/survey", createSurveyHandler.Handle)

	e.Start(":8080")
}
