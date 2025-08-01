package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
)

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
)

const (
	GetBothAccountsForTransactionSQL = `SELECT * FROM accounts WHERE id IN (?, ?) ORDER BY id FOR UPDATE`
	GetAllTransactionsSQL            = `SELECT * FROM transactions ORDER BY id LIMIT ? OFFSET ?`
	GetTransactionsByAccountSQL      = `SELECT * FROM transactions WHERE from_account_id = ? OR to_account_id = ? ORDER BY id LIMIT ? OFFSET ?`
	UpdateSenderAccountBalanceSQL    = `UPDATE accounts SET balance = balance - :amount WHERE id = :from_account_id`
	UpdateReceiverAccountBalanceSQL  = `UPDATE accounts SET balance = balance + :amount WHERE id = :to_account_id`
	InsertTransactionRecordSQL       = `INSERT INTO transactions (from_account_id, to_account_id, amount, type, description, created_at) VALUES (:from_account_id, :to_account_id, :amount, :type, :description, :created_at)`
	DeleteTransactionSQL             = `DELETE FROM transactions WHERE id = ?`
)

type Transaction struct {
	ID          int             `db:"id" json:"id"`
	Sender      *int            `db:"from_account_id" json:"from_account_id"`
	Receiver    *int            `db:"to_account_id" json:"to_account_id"`
	Amount      *float64        `db:"amount" json:"amount"`
	Type        TransactionType `db:"type" json:"type"`
	Description string          `db:"description" json:"description"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
}

func (t *Transaction) Validate() error {
	if t.Sender == nil {
		return errors.New("sender id is required")
	}

	if t.Receiver == nil {
		return errors.New("receiver id is required")
	}

	if t.Amount == nil {
		return errors.New("amount is required")
	}

	if t.Type != Deposit && t.Type != Withdrawal && t.Type != Transfer {
		return errors.New("transaction type is required and has to be: deposit, withdrawal or transfer")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

func GetAllTransactions(db *sqlx.DB, transactions *[]Transaction, limit int, offset int) error {
	err := db.Select(transactions, GetAllTransactionsSQL, limit, offset)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetTransactionsByAccount(db *sqlx.DB, transactions *[]Transaction, accountId int, limit int, offset int) error {
	err := db.Select(transactions, GetTransactionsByAccountSQL, accountId, accountId, limit, offset)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func CreateTransaction(db *sqlx.DB, t *Transaction) (statusCode int, err error) {
	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		return http.StatusInternalServerError, errors.New("server error")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				statusCode = http.StatusInternalServerError
			} else {
				statusCode = http.StatusOK
			}
		}
	}()

	// Lock both accounts in one query to avoid deadlocks
	var accounts []Account
	err = tx.Select(&accounts, GetBothAccountsForTransactionSQL, *t.Sender, *t.Receiver)
	if err != nil {
		logger.Error("[Create Transaction DB]: Failed to lock accounts", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	if len(accounts) != 2 {
		logger.Error("[Create Transaction DB]: One or both accounts not found", accounts)
		return http.StatusNotFound, errors.New("sender or receiver account not found")
	}

	// Identify sender and receiver accounts from result
	var senderAccount, receiverAccount *Account
	for i := range accounts {
		switch accounts[i].ID {
		case *t.Sender:
			senderAccount = &accounts[i]
		case *t.Receiver:
			receiverAccount = &accounts[i]
		}
	}

	if senderAccount == nil || receiverAccount == nil {
		logger.Error("[Create Transaction DB]: Failed to assign sender/receiver", accounts)
		return http.StatusInternalServerError, errors.New("failed to process accounts")
	}

	// Validate sender has enough balance
	if *senderAccount.Balance < *t.Amount {
		logger.Error("[Create Transaction DB]: Insufficient balance", senderAccount)
		return http.StatusBadRequest, errors.New("insufficient balance in sender account")
	}

	// Update sender balance
	_, err = tx.NamedExec(UpdateSenderAccountBalanceSQL, &t)
	if err != nil {
		logger.Error("[Create Transaction DB]: Couldn't update sender balance", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	// Update receiver balance
	_, err = tx.NamedExec(UpdateReceiverAccountBalanceSQL, &t)
	if err != nil {
		logger.Error("[Create Transaction DB]: Couldn't update receiver balance", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	// Insert transaction record
	_, err = tx.NamedExec(InsertTransactionRecordSQL, map[string]interface{}{
		"from_account_id": *t.Sender,
		"to_account_id":   *t.Receiver,
		"amount":          *t.Amount,
		"type":            t.Type,
		"description":     t.Description,
		"created_at":      time.Now(),
	})
	if err != nil {
		logger.Error("[Create Transaction DB]: Couldn't insert transaction record", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	return
}

func DeleteTransaction(db *sqlx.DB, id int) (int, error) {
	res, err := db.Exec(DeleteTransactionSQL, id)
	if err != nil {
		logger.Error("[Delete Transaction DB]: could not delete transaction", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("[Delete Transaction DB]: Could not verify transaction deletion", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	if rowsAffected == 0 {
		logger.Error("[Delete Transaction DB]: Could not find transaction with provided id", nil)
		return http.StatusNotFound, errors.New("could not find transaction with provided id")
	}

	return http.StatusOK, nil
}
