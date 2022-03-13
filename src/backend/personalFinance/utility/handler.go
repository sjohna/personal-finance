package utility

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HandlerLogger(r *http.Request, handler string) *logrus.Entry {
	log := r.Context().Value("logger").(*logrus.Entry).WithField("handler", handler)
	log.Info("Called")
	return log
}

func RespondError(log *logrus.Entry, w http.ResponseWriter, err error, httpResponseCode int) {
	log.WithError(err).WithField("httpResponseCode", httpResponseCode).Error()
	http.Error(w, err.Error(), httpResponseCode)
}

func RespondJSON(log *logrus.Entry, w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(value)
	if err != nil {
		RespondError(log, w, err, 500)
		return
	}

	written, err := w.Write(jsonResp)
	if err != nil {
		log.Error("Error writing response")
	} else {
		log.WithField("responseBytesWritten", written).Info("Succeeded")
	}
}
