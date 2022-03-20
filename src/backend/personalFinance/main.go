package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/handler"
	"github.com/sjohna/personal-finance/repo"
	"github.com/sjohna/personal-finance/service"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/lib/pq"
)

func configureLogging() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	fileLogger := &lumberjack.Logger{
		Filename:   "./pf.log",
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     365,
		Compress:   false,
	}

	writer := io.MultiWriter(fileLogger, os.Stdout)

	logrus.SetOutput(writer)

	logrus.WithField("startup", true).Info("Logging configured")
}

func main() {
	configureLogging()
	log := logrus.WithField("startup", true)
	log.Info("Starting up")

	db, err := repo.Connect("localhost", "pf-test")
	if err != nil {
		log.WithError(err).Error()
		return
	}

	err = repo.CreateTables(db, log)
	if err != nil {
		log.WithError(err).Error()
		return
	}

	repo := repo.Repo{DB: db}
	accountService := service.AccountService{Repo: &repo}
	accountHandler := handler.AccountHandler{AccountService: &accountService}

	// init chi

	r := chi.NewRouter()
	r.Use(handler.LogRequestContext)
	accountHandler.ConfigureRoutes(r)

	log.Info("Listening on port 3000")

	http.ListenAndServe(":3000", r)

	log.Info("Done")
}
