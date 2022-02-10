package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (service *PFService) CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var params createAccountParams
	err = json.Unmarshal(body, &params)
	if err != nil { // TODO: better error handling
		http.Error(w, err.Error(), 500)
		return
	}

	createdAccount, err := pfdb.CreateAccount(service.DB, params.AccountName, params.AccountDesc)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(createdAccount)
	w.Write(jsonResp)
}

func (service *PFService) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := pfdb.GetAccounts(service.DB)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(accounts)
	w.Write(jsonResp)
}

func (service *PFService) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountIDString := chi.URLParam(r, "accountID")
	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	account, err := pfdb.GetAccount(service.DB, accountID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(account)
	w.Write(jsonResp)
}
