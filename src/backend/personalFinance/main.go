package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/account"
	"github.com/sjohna/personal-finance/pfdb"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/lib/pq"
)

var requestCounter int64 = 0

func getNextRequestId() int64 {
	return atomic.AddInt64(&requestCounter, 1)
}

func LogRequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logrus.WithFields(logrus.Fields{
			"request": getNextRequestId(),
		})

		log.WithFields(logrus.Fields{
			"route":         r.URL.Path,
			"method":        r.Method,
			"contentLength": r.ContentLength,
		}).Info()

		ctx := context.WithValue(r.Context(), "logger", log)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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

	accountRepo := account.AccountRepo{DB: db}
	accountHandler := account.AccountHandler{AccountRepo: &accountRepo}

	// init chi

	r := chi.NewRouter()
	r.Use(LogRequestContext)
	accountHandler.ConfigureRoutes(r)

	log.Info("Listening on port 3000")

	http.ListenAndServe(":3000", r)

	logrus.Info("Done")
}
