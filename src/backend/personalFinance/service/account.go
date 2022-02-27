package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/sjohna/personal-finance/pfdb"
)

type AccountService interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccounts(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
}

type createAccountParams struct {
	AccountName string `json:"accountName"`
	AccountDesc string `json:"accountDesc"`
}

func handlerLogger(r *http.Request, handler string) *logrus.Entry {
	log := r.Context().Value("logger").(*logrus.Entry).WithField("handler", handler)
	log.Info("Called")
	return log
}

func respondError(log *logrus.Entry, w http.ResponseWriter, err error, httpResponseCode int) {
	log.WithError(err).WithField("httpResponseCode", httpResponseCode).Error()
	http.Error(w, err.Error(), httpResponseCode)
}

func respondJSON(log *logrus.Entry, w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(value)
	if err != nil {
		respondError(log, w, err, 500)
		return
	}

	written, err := w.Write(jsonResp)
	if err != nil {
		log.Error("Error writing response")
	} else {
		log.WithField("responseBytesWritten", written).Info("Succeeded")
	}
}

func (service *PFService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateAccount")

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		respondError(log, w, err, 500)
		return
	}

	// Unmarshal
	var params createAccountParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		respondError(log, w, err, 500)
		return
	}

	createdAccount, err := pfdb.CreateAccount(log, service.DB, params.AccountName, params.AccountDesc)
	if err != nil {
		respondError(log, w, err, 500)
		return
	}

	respondJSON(log, w, createdAccount)
}

func (service *PFService) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetAccounts")

	accounts, err := pfdb.GetAccounts(log, service.DB)
	if err != nil {
		respondError(log, w, err, 500)
		return
	}

	respondJSON(log, w, accounts)
}

func (service *PFService) GetAccount(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetAccount")

	accountIDString := chi.URLParam(r, "accountID")
	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	account, err := pfdb.GetAccount(log, service.DB, accountID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	respondJSON(log, w, account)
}
