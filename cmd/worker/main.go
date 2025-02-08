package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ffelipelimao/survey/internal/consumer"
	"github.com/ffelipelimao/survey/internal/database"
	"github.com/ffelipelimao/survey/internal/processor"
	"github.com/ffelipelimao/survey/internal/repository"
	usecases "github.com/ffelipelimao/survey/internal/use_cases"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)
		cancel()
	}()

	fmt.Println("Worker Starting...")
	consumer.Start()
}
