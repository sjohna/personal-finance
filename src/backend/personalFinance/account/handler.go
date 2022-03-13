package account

import (
	"encoding/json"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utility.RespondError(log, w, err, 500)
		return
	}

	// Unmarshal
	var params createAccountParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		utility.RespondError(log, w, err, 500)
		return
	}

	createdAccount, err := handler.AccountRepo.CreateAccount(log, params.AccountName, params.AccountDesc)
	if err != nil {
		utility.RespondError(log, w, err, 500)
		return
	}

	utility.RespondJSON(log, w, createdAccount)
}

func (handler *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	log := utility.HandlerLogger(r, "GetAccounts")

	accounts, err := handler.AccountRepo.GetAccounts(log)
	if err != nil {
		utility.RespondError(log, w, err, 500)
		return
	}

	utility.RespondJSON(log, w, accounts)
}

func (handler *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	log := utility.HandlerLogger(r, "GetAccount")

	accountIDString := chi.URLParam(r, "accountID")
	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	account, err := handler.AccountRepo.GetAccount(log, accountID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utility.RespondJSON(log, w, account)
}
