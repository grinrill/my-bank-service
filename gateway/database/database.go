package database

import (
	"database/sql"

	"github.com/grinrill/my-bank-service/domain/bankRepo"
	"github.com/grinrill/my-bank-service/domain/bankRules"
)

// NewBankRepo Create new database
func NewBankRepo() (database BankRepo, cancel bankRepo.CancelFunc, err error) {
	db, cancelDB, err := newDB()
	if err!=nil {
		return
	}

	database = BankRepo{
		db: db,
		accountStmt: db.Prepare("SELECT * FROM accounts WHERE accountID=(?)"),
		accountCurrencyStmt: db.Prepare("SELECT currency FROM accounts WHERE accountID=(?)"),
		accountBalanceStmt: db.Prepare("SELECT balance FROM accounts WHERE accountID=(?)"),
		incrementBalanceStmt: db.Prepare("UPDATE accounts SET balance=balance+(?) WHERE accountID=(?)"),
	}

	cancel = func() error {
		for c := range []bankRepo.CancelFunc{
			database.accountStmt.Close,
			database.accountCurrencyStmt.Close,
			database.accountBalanceStmt.Close,
			database.incrementBalanceStmt.Close,
			cancelDB,
		} {
			if err := c; err!=nil {
				return err
			}
		}
	}

	return
}

// BankRepo Structure for working with bank accounts database
type BankRepo struct {
	db                   *sql.DB
	accountStmt          *sql.Stmt
	accountCurrencyStmt  *sql.Stmt
	accountBalanceStmt   *sql.Stmt
	incrementBalanceStmt *sql.Stmt
}

// GetAccountCurrency Return account currency
func (r BankRepo) GetAccountCurrency(accountID int) (cur bankRules.Currency, err error) {
	err = r.accountCurrencyStmt.QueryRow(accountID).Scan(&cur)
	return
}

// GetAccountBalance Return account balance
func (r BankRepo) GetAccountBalance(accountID int) (balance float64, err error) {
	err = r.accountBalanceStmt.QueryRow(accountID).Scan(&balance)
	return
}

// IncrementAccountBalance Inrement account balance
func (r BankRepo) IncrementAccountBalance(accountID int, sum float64) (err error) {
	err = r.incrementBalanceStmt.Exec(sum, accountID)
	return
}

// DecrementAccountBalance Decrement account balance
func (r BankRepo) DecrementAccountBalance(accountID int, sum float64) (err error) {
	return r.IncrementAccountBalance(accountID, -sum)
}

