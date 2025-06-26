package repository

import (
	"context"
	"fmt"

	"github.com/textures1245/payso-check-slip-backend/model"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
)

func GetAllBank() ([]model.SureSureBank, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureBank{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_BANK_ACCOUNT_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureBank{}, err
	}

	var bankAccounts []model.SureSureBank
	err = scan.Rows(&bankAccounts, rows)

	defer rows.Close()
	log.Infof("bankAccounts: %d", len(bankAccounts))
	return bankAccounts, nil
}

func GetBankByID(id int) ([]model.SureSureBank, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureBank{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_BANK_ACCOUNT_GET_BYID, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return []model.SureSureBank{}, err
	}
	var bankAccount []model.SureSureBank
	err = scan.Rows(&bankAccount, rows)
	defer rows.Close()
	return bankAccount, nil
}

func CreateBank(account model.SureSureBank) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	// Build query dynamically
	query := "INSERT INTO SureSureBank ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	if account.UserID != 0 {
		query += "UserID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.UserID)
		counter++
	}

	if account.BankCode != "" {
		query += "BankCode, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.BankCode)
		counter++
	}

	if account.PromptPayType != "" {
		query += "PromptPayType, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.PromptPayType)
		counter++
	}

	if account.AccountNo != "" {
		query += "AccountNo, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.AccountNo)
		counter++
	}

	if account.AccountType != "" {
		query += "AccountType, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.AccountType)
		counter++
	}

	if account.NameTH != "" {
		query += "NameTH, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.NameTH)
		counter++
	}

	if account.NameEN != "" {
		query += "NameEN, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.NameEN)
		counter++
	}

	if account.IsActive {
		query += "IsActive, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.IsActive)
		counter++
	}

	if account.CreatedDate != "" {
		query += "CreatedDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.CreatedDate)
		counter++
	}

	if account.UpdatedDate != "" {
		query += "UpdatedDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, account.UpdatedDate)
		counter++
	}

	query = query[:len(query)-2] + ") "
	values = values[:len(values)-2] + ")"
	finalQuery := query + values + " RETURNING ID"

	log.Infof("finalQuery: %v", finalQuery)
	result := conn.QueryRowContext(ctx, finalQuery, params...)

	// Retrieve the last inserted ID
	var lastInsertedID int64
	if err := result.Scan(&lastInsertedID); err != nil {
		log.Errorf("Error retrieving last inserted ID: %v", err)
		return 0, err
	}

	return int(lastInsertedID), nil
}

func UpdateBank(account model.SureSureBank) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// Initialize query parts
	query := "UPDATE SureSureBank SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if account.UserID != 0 {
		query += fmt.Sprintf("UserID = $%d, ", counter)
		params = append(params, account.UserID)
		counter++
	}

	if account.BankCode != "" {
		query += fmt.Sprintf("BankCode = $%d, ", counter)
		params = append(params, account.BankCode)
		counter++
	}

	if account.PromptPayType != "" {
		query += fmt.Sprintf("PromptPayType = $%d, ", counter)
		params = append(params, account.PromptPayType)
		counter++
	}

	if account.AccountNo != "" {
		query += fmt.Sprintf("AccountNo = $%d, ", counter)
		params = append(params, account.AccountNo)
		counter++
	}

	if account.AccountType != "" {
		query += fmt.Sprintf("AccountType = $%d, ", counter)
		params = append(params, account.AccountType)
		counter++
	}

	if account.NameTH != "" {
		query += fmt.Sprintf("NameTH = $%d, ", counter)
		params = append(params, account.NameTH)
		counter++
	}

	if account.NameEN != "" {
		query += fmt.Sprintf("NameEN = $%d, ", counter)
		params = append(params, account.NameEN)
		counter++
	}

	if account.IsActive {
		query += fmt.Sprintf("IsActive = $%d, ", counter)
		params = append(params, account.IsActive)
		counter++
	}

	if account.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = $%d, ", counter)
		params = append(params, account.UpdatedDate)
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = $" + fmt.Sprintf("%d", counter)
	params = append(params, account.ID)

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteBank(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_BANK_ACCOUNT_DELETE, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
