package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sjohna/personal-finance/pfdb"
	"github.com/sjohna/personal-finance/service"

	_ "github.com/lib/pq"
)

func main() {
	db, err := pfdb.Connect("localhost", "pf-test")

	if err != nil {
		return
	}

	pfdb.CreateTables(db)

	service := service.PFService{DB: db}

	// init chi

	r := chi.NewRouter()
	r.Post("/createAccount", service.CreateAccount)
	r.Get("/account/{accountID}", service.GetAccount)
	r.Get("/account", service.GetAccounts)

	http.ListenAndServe(":3000", r)

	fmt.Println("We ran...")
}
