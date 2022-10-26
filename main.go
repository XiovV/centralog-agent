package main

import (
	"fmt"
	"github.com/XiovV/centralog-agent/pkg/docker"
	"github.com/XiovV/centralog-agent/repository"
	"github.com/XiovV/centralog-agent/server"
	"github.com/fvbock/endless"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"os"
)

var (
	appEnv = os.Getenv("APP_ENV")
)

func generateMessages(size int) []string {
	messages := []string{}

	for i := 1; i <= size; i++ {
		messages = append(messages, fmt.Sprintf("message%d", i))
	}

	return messages
}

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalln("couldn't init logger: ", err)
	}

	defer logger.Sync()

	repo := repository.New()
	dockerController := docker.New(repo)

	srv := server.Server{
		Logger:     logger,
		Docker:     dockerController,
		Repository: repo,
	}

	logger.Info("running...", zap.String("port", os.Getenv("PORT")))

	endless.ListenAndServe(":"+os.Getenv("PORT"), srv.Serve())
}

func generateApiKey() string {
	apiKeyLength := 39

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, apiKeyLength)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func initLogger() (*zap.Logger, error) {
	if appEnv == server.LOCAL_ENV || appEnv == server.STAGING_ENV {
		logger, err := zap.NewDevelopment()

		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	logger, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	return logger, nil
}