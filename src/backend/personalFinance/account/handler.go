package account

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/utility"
)

func (handler *AccountHandler) ConfigureRoutes(base *chi.Mux) {
	base.Post("/account", handler.CreateAccount)
	base.Get("/account/{accountID}", handler.GetAccount)
	base.Get("/account", handler.GetAccounts)
}

type AccountHandler struct {
	AccountRepo *AccountRepo
}

type createAccountParams struct {
	AccountName string `json:"accountName"`
	AccountDesc string `json:"accountDesc"`
}

func (handler *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	log := utility.HandlerLogger(r, "CreateAccount")

	var params createAccountParams
	if err := utility.UnmarshalRequestBody(log, r, &params); err != nil {
		utility.RespondInternalServerError(log, w, err)
		return
	}

	createdAccount, err := handler.AccountRepo.CreateAccount(log, params.AccountName, params.AccountDesc)
	if err != nil {
		utility.RespondInternalServerError(log, w, err)
		return
	}

	utility.RespondJSON(log, w, createdAccount)
}

func (handler *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := utility.HandlerLogger(r, "GetAccounts")

	accounts, err := handler.AccountRepo.GetAccounts(log)
	if err != nil {
		utility.RespondInternalServerError(log, w, err)
		return
	}

	utility.RespondJSON(log, w, accounts)
}

func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	log := utility.HandlerLogger(r, "GetAccount")

	accountIDString := chi.URLParam(r, "accountID")
	log.WithField("accountID", accountIDString).Info("Params")

	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		utility.RespondInternalServerError(log, w, err)
		return
	}

	account, err := handler.AccountRepo.GetAccount(log, accountID)
	if err != nil {
		utility.RespondInternalServerError(log, w, err)
		return
	}

	utility.RespondJSON(log, w, account)
}
