package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/service"
)

type CurrencyHandler struct {
	Service *service.CurrencyService
}

func (handler *CurrencyHandler) ConfigureRoutes(base *chi.Mux) {
	base.Post("/currency", handler.CreateCurrency)
	base.Get("/currency/{currencyID}", handler.GetCurrency)
	base.Get("/currency", handler.GetCurrencies)
}

func (handler *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "CreateCurrency")
	defer logHandlerReturn(log)

	var params struct {
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Magnitude    int    `json:"magnitude"`
	}

	if err := unmarshalRequestBody(log, r, &params); err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	createdCurrency, err := handler.Service.CreateCurrency(log, params.Name, params.Abbreviation, params.Magnitude)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, createdCurrency)
}

func (handler *CurrencyHandler) GetCurrencies(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetCurrencies")
	defer logHandlerReturn(log)

	currencies, err := handler.Service.GetCurrencies(log)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, currencies)
}

func (handler *CurrencyHandler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	log := handlerLogger(r, "GetCurrency")
	defer logHandlerReturn(log)

	currencyIDString := chi.URLParam(r, "currencyID")
	log.WithField("currencyID", currencyIDString).Info("Params")

	currencyID, err := strconv.ParseInt(currencyIDString, 10, 64)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	currency, err := handler.Service.GetCurrency(log, currencyID)
	if err != nil {
		respondInternalServerError(log, w, err)
		return
	}

	respondJSON(log, w, currency)
}
