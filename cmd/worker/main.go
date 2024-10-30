package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ffelipelimao/survey/internal/consumer"
	"github.com/ffelipelimao/survey/internal/database"
	"github.com/ffelipelimao/survey/internal/processor"
	"github.com/ffelipelimao/survey/internal/repository"
	usecases "github.com/ffelipelimao/survey/internal/use_cases"
)

func main() {
	ctx := context.Background()

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
	saveSurveyUseCase := usecases.NewSaveSurveyUseCase(surveyRepository)
	saveSurveyProcessor := processor.NewSaveSurveyProcessor(saveSurveyUseCase)

	consumer, err := consumer.New(ctx, "tp.create-survey", saveSurveyProcessor.Handle)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()

	fmt.Println("Worker Starting...")
	consumer.Start()
}
