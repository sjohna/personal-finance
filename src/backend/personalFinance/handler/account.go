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
	log := HandlerLogger(r, "CreateAccount")

	var params createAccountParams
	if err := UnmarshalRequestBody(log, r, &params); err != nil {
		RespondInternalServerError(log, w, err)
		return
	}

	createdAccount, err := handler.AccountService.CreateAccount(log, params.AccountName, params.AccountDesc)
	if err != nil {
		RespondInternalServerError(log, w, err)
		return
	}

	RespondJSON(log, w, createdAccount)
}

func (handler *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := HandlerLogger(r, "GetAccounts")

	accounts, err := handler.AccountService.GetAccounts(log)
	if err != nil {
		RespondInternalServerError(log, w, err)
		return
	}

	RespondJSON(log, w, accounts)
}

func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	log := HandlerLogger(r, "GetAccount")

	accountIDString := chi.URLParam(r, "accountID")
	log.WithField("accountID", accountIDString).Info("Params")

	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		RespondInternalServerError(log, w, err)
		return
	}

	account, err := handler.AccountService.GetAccount(log, accountID)
	if err != nil {
		RespondInternalServerError(log, w, err)
		return
	}

	RespondJSON(log, w, account)
}
