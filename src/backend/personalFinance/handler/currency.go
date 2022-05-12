package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/service"
)

type CurrencyHandler struct {
	Service *service.CurrencyService
}

func (handler *CurrencyHandler) ConfigureRoutes(base *chi.Mux) {
	base.Post("/currency", handler.CreateCurrency)

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
