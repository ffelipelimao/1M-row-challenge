package main

import (
	"context"
	"log"

	"github.com/ffelipelimao/survey/internal/database"
	"github.com/ffelipelimao/survey/internal/handlers"
	"github.com/ffelipelimao/survey/internal/publisher"
	"github.com/ffelipelimao/survey/internal/repository"
	usecases "github.com/ffelipelimao/survey/internal/use_cases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	publisherSurvey, err := publisher.New("tp.create-survey")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	surveyRepository := repository.NewSurveyPostgresRepository(db)

	defer publisherSurvey.Stop()

	publisherSurveyUseCase := usecases.NewPublisherSurveyUseCase(publisherSurvey)
	publisherSurveyHandler := handlers.NewPublisherSurveyHandler(publisherSurveyUseCase)

	listSurveyUseCase := usecases.NewListSurveysUseCase(surveyRepository)
	listSurveyHandler := handlers.NewListSurveyHandler(listSurveyUseCase)

	getSurveyAvgUseCase := usecases.NewGetSurveyAvgUseCase(surveyRepository)
	getSurveyAvgHandler := handlers.NewGetSurveyAvgHandler(getSurveyAvgUseCase)

	// Routes
	e.POST("/survey", publisherSurveyHandler.Handle)
	e.GET("/survey/:merchant_id", listSurveyHandler.Handle)
	e.GET("/survey/:merchant_id/avg", getSurveyAvgHandler.Handle)

	e.Start(":8080")
}
