package service

import "github.com/jmoiron/sqlx"

type PFService struct {
	DB *sqlx.DB
}
