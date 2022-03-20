package repo

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DAO interface {
	Logger() *logrus.Entry

	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Unsafe() DAO
}

type DBDAO struct {
	sqlxer *sqlx.DB
	logger *logrus.Entry
}

type TxDAO struct {
	sqlxer *sqlx.Tx
	logger *logrus.Entry
	err    error
}

func NewDBDAO(db *sqlx.DB, logger *logrus.Entry) *DBDAO {
	if db == nil {
		logger.Fatal("db parameter not provided to NewDBDAO!")
	}

	if logger == nil {
		logger.Fatal("logger parameter not provided to NewDBDAO!")
	}

	// TODO: get sequence number here and add field(s) to logger
	//log := logger.WithField("repo-dao")

	return &DBDAO{
		db,
		logger,
	}
}

func (dao *DBDAO) Logger() *logrus.Entry {
	return dao.logger
}

func (dao *DBDAO) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dao.sqlxer.Exec(query, args...)
}

func (dao *DBDAO) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return dao.sqlxer.ExecContext(ctx, query, args...)
}

func (dao *DBDAO) Get(dest interface{}, query string, args ...interface{}) error {
	return dao.sqlxer.Get(dest, query, args...)
}

func (dao *DBDAO) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return dao.sqlxer.GetContext(ctx, dest, query, args...)
}

func (dao *DBDAO) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return dao.sqlxer.NamedExec(query, arg)
}

func (dao *DBDAO) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return dao.NamedExecContext(ctx, query, arg)
}

func (dao *DBDAO) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return dao.sqlxer.NamedQuery(query, arg)
}

func (dao *DBDAO) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return dao.sqlxer.PrepareNamed(query)
}

func (dao *DBDAO) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return dao.sqlxer.PrepareNamedContext(ctx, query)
}

func (dao *DBDAO) Preparex(query string) (*sqlx.Stmt, error) {
	return dao.sqlxer.Preparex(query)
}

func (dao *DBDAO) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return dao.sqlxer.PreparexContext(ctx, query)
}

func (dao *DBDAO) Rebind(query string) string {
	return dao.Rebind(query)
}

func (dao *DBDAO) Select(dest interface{}, query string, args ...interface{}) error {
	return dao.sqlxer.Select(dest, query, args...)
}

func (dao *DBDAO) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return dao.sqlxer.SelectContext(ctx, dest, query, args...)
}

func (dao *DBDAO) Unsafe() DAO {
	return &DBDAO{
		dao.sqlxer.Unsafe(),
		dao.logger.WithField("repo-unsafe", true),
	}
}

func (dao *DBDAO) Close() error {
	dao.logger.Info("Close")
	return nil
}

func (dao *TxDAO) Logger() *logrus.Entry {
	return dao.logger
}

func (dao *TxDAO) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := dao.sqlxer.Exec(query, args...)
	if err != nil {
		dao.err = err
	}

	return result, err
}

func (dao *TxDAO) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := dao.sqlxer.ExecContext(ctx, query, args...)
	if err != nil {
		dao.err = err
	}

	return result, err
}

func (dao *TxDAO) Get(dest interface{}, query string, args ...interface{}) error {
	err := dao.sqlxer.Get(dest, query, args...)
	if err != nil {
		dao.err = err
	}

	return err
}

func (dao *TxDAO) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := dao.sqlxer.GetContext(ctx, dest, query, args...)
	if err != nil {
		dao.err = err
	}

	return err
}

func (dao *TxDAO) NamedExec(query string, arg interface{}) (sql.Result, error) {
	result, err := dao.sqlxer.NamedExec(query, arg)
	if err != nil {
		dao.err = err
	}

	return result, err
}

func (dao *TxDAO) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	result, err := dao.NamedExecContext(ctx, query, arg)
	if err != nil {
		dao.err = err
	}

	return result, err
}

func (dao *TxDAO) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	rows, err := dao.sqlxer.NamedQuery(query, arg)
	if err != nil {
		dao.err = err
	}

	return rows, err
}

func (dao *TxDAO) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	namedStmnt, err := dao.sqlxer.PrepareNamed(query)
	if err != nil {
		dao.err = err
	}

	return namedStmnt, err
}

func (dao *TxDAO) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	namedStmt, err := dao.sqlxer.PrepareNamedContext(ctx, query)
	if err != nil {
		dao.err = err
	}

	return namedStmt, err
}

func (dao *TxDAO) Preparex(query string) (*sqlx.Stmt, error) {
	stmt, err := dao.sqlxer.Preparex(query)
	if err != nil {
		dao.err = err
	}

	return stmt, err
}

func (dao *TxDAO) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	stmt, err := dao.sqlxer.PreparexContext(ctx, query)
	if err != nil {
		dao.err = err
	}

	return stmt, err
}

func (dao *TxDAO) Rebind(query string) string {
	return dao.Rebind(query)
}

func (dao *TxDAO) Select(dest interface{}, query string, args ...interface{}) error {
	err := dao.sqlxer.Select(dest, query, args...)
	if err != nil {
		dao.err = err
	}

	return err
}

func (dao *TxDAO) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := dao.sqlxer.SelectContext(ctx, dest, query, args...)
	if err != nil {
		dao.err = err
	}

	return err
}

func (dao *TxDAO) Unsafe() DAO {
	return &TxDAO{
		dao.sqlxer.Unsafe(),
		dao.logger.WithField("repo-unsafe", true),
		dao.err,
	}
}

func (dao *TxDAO) Close() error {
	if dao.err != nil {
		dao.logger.Warn("Rolling back transaction")
		rollbackErr := dao.sqlxer.Rollback()
		if rollbackErr != nil {
			dao.logger.WithError(rollbackErr).Error("Error rolling back transaction!")
		}
		dao.logger.Info("Rolled back")

		return dao.err
	}

	if dao.err == nil {
		commitErr := dao.sqlxer.Commit()
		if commitErr != nil {
			dao.logger.WithError(commitErr).Error("Error committing transaction!")
			return commitErr
		}
		dao.logger.Info("Committed")
	}

	return nil
}
