package service

import "github.com/sirupsen/logrus"

func ServiceFunctionLogger(log *logrus.Entry, serviceFunction string) *logrus.Entry {
	log = log.WithField("service-function", serviceFunction)
	log.Info("Called")
	return log
}
