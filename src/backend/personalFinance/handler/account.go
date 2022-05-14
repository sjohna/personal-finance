package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/service"
)

func (handler *AccountHandler) ConfigureRoutes(base *chi.Mux) {
	base.Post("/account", handler.CreateAccount)
	base.Get("/account/{accountID}", handler.GetAccount)
	base.Get("/account", handler.GetAccounts)
	base.Post("/account/{accountID}/debit", handler.CreateDebit)
	base.Post("/account/{accountID}/credit", handler.CreateCredit)
}

type AccountHandler struct {
	Service *service.AccountService
}

func (handler *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateAccount")
	defer logHandlerReturn(log)

	var params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := unmarshalRequestBody(log, r, &params); err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	createdAccount, err := handler.Service.CreateAccount(log, params.Name, params.Description)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, createdAccount)
}

func (handler *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetAccounts")
	defer logHandlerReturn(log)

	accounts, err := handler.Service.GetAccounts(log)
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

	account, err := handler.Service.GetAccount(log, accountID)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, account)
}

func (handler *AccountHandler) CreateDebit(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateDebit")
	defer logHandlerReturn(log)

	accountIDString := chi.URLParam(r, "accountID")
	log.WithField("accountID", accountIDString).Info("Params")

	accountID, err := strconv.ParseInt(accountIDString, 10, 64)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	var params struct {
		Amount     int64     `json:"amount"`
		CurrencyId int64     `json:"currencyId"`
		Time       time.Time `json:"time"`
	}

	if err := unmarshalRequestBody(log, r, &params); err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	debit, err := handler.Service.CreateDebit(log, accountID, params.Amount, params.CurrencyId, params.Time)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, debit)
}

func (handler *AccountHandler) CreateCredit(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateCredit")
	defer logHandlerReturn(log)

	accountIDString := chi.URLParam(r, "accountID")
	log.WithField("accountID", accountIDString).Info("Params")

	accountID, err := strconv.ParseInt(accountIDString, 10, 64)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	var params struct {
		Amount     int64     `json:"amount"`
		CurrencyId int64     `json:"currencyId"`
		Time       time.Time `json:"time"`
	}

	if err := unmarshalRequestBody(log, r, &params); err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	credit, err := handler.Service.CreateCredit(log, accountID, params.Amount, params.CurrencyId, params.Time)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, credit)
}
