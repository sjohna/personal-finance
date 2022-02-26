package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/pfdb"
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

	db, err := pfdb.Connect("localhost", "pf-test")
	if err != nil {
		log.WithError(err).Error()
		return
	}

	err = pfdb.CreateTables(db)
	if err != nil {
		log.WithError(err).Error()
		return
	}

	service := service.PFService{DB: db}

	// init chi

	r := chi.NewRouter()
	r.Post("/createAccount", service.CreateAccount)
	r.Get("/account/{accountID}", service.GetAccount)
	r.Get("/account", service.GetAccounts)

	log.Info("Listening on port 3000")

	http.ListenAndServe(":3000", r)

	logrus.Info("Done")
}
