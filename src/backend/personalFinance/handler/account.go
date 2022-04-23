package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/service"
)

func (handler *AccountHandler) ConfigureRoutes(base *chi.Mux) {
	base.Post("/account", handler.CreateAccount)
	base.Get("/account/{accountID}", handler.GetAccount)
	base.Get("/account", handler.GetAccounts)
}

type AccountHandler struct {
	AccountService *service.AccountService
}

type createAccountParams struct {
	AccountName string `json:"accountName"`
	AccountDesc string `json:"accountDesc"`
}

func (handler *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateAccount")
	defer logHandlerReturn(log)

	var params createAccountParams
	if err := unmarshalRequestBody(log, r, &params); err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	createdAccount, err := handler.AccountService.CreateAccount(log, params.AccountName, params.AccountDesc)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, createdAccount)
}

func (handler *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetAccounts")
	defer logHandlerReturn(log)

	accounts, err := handler.AccountService.GetAccounts(log)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, accounts)
}

func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetAccount")
	defer logHandlerReturn(log)

	accountIDString := chi.URLParam(r, "accountID")
	log.WithField("accountID", accountIDString).Info("Params")

	accountID, err := strconv.ParseInt(accountIDString, 10, 64)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	account, err := handler.AccountService.GetAccount(log, accountID)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, account)
}
