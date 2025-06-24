package repository

import (
	"context"
	"database/sql"
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
	rows, err := conn.QueryContext(ctx, model.SQL_BANK_ACCOUNT_GET_BYID, sql.Named("ID", id))
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
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.UserID))
		counter++
	}

	if account.BankCode != "" {
		query += "BankCode, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.BankCode))
		counter++
	}

	if account.PromptPayType != "" {
		query += "PromptPayType, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.PromptPayType))
		counter++
	}

	if account.AccountNo != "" {
		query += "AccountNo, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.AccountNo))
		counter++
	}

	if account.AccountType != "" {
		query += "AccountType, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.AccountType))
		counter++
	}

	if account.NameTH != "" {
		query += "NameTH, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.NameTH))
		counter++
	}

	if account.NameEN != "" {
		query += "NameEN, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.NameEN))
		counter++
	}

	if account.IsActive {
		query += "IsActive, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.IsActive))
		counter++
	}

	if account.CreatedDate != "" {
		query += "CreatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.CreatedDate))
		counter++
	}

	if account.UpdatedDate != "" {
		query += "UpdatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.UpdatedDate))
		counter++
	}

	query = query[:len(query)-2] + ") "
	values = values[:len(values)-2] + ")"
	finalQuery := query + " " + values

	log.Infof("finalQuery: %v", finalQuery)
	result, err := conn.ExecContext(ctx, finalQuery, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return 0, err
	}

	log.Infof("result: %v", result)
	// Retrieve the last inserted ID
	// lastInsertedID, err := result.LastInsertId()
	// if err != nil {
	// 	log.Errorf("Error retrieving last insert ID: %v", err)
	// 	return 0, err
	// }

	return 0, nil
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
		query += fmt.Sprintf("UserID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.UserID))
		counter++
	}

	if account.BankCode != "" {
		query += fmt.Sprintf("BankCode = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.BankCode))
		counter++
	}

	if account.PromptPayType != "" {
		query += fmt.Sprintf("PromptPayType = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.PromptPayType))
		counter++
	}

	if account.AccountNo != "" {
		query += fmt.Sprintf("AccountNo = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.AccountNo))
		counter++
	}

	if account.AccountType != "" {
		query += fmt.Sprintf("AccountType = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.AccountType))
		counter++
	}

	if account.NameTH != "" {
		query += fmt.Sprintf("NameTH = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.NameTH))
		counter++
	}

	if account.NameEN != "" {
		query += fmt.Sprintf("NameEN = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.NameEN))
		counter++
	}

	if account.IsActive {
		query += fmt.Sprintf("IsActive = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.IsActive))
		counter++
	}

	if account.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.UpdatedDate))
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = @p" + fmt.Sprintf("%d", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), account.ID))

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
	rows, err := conn.QueryContext(ctx, model.SQL_BANK_ACCOUNT_DELETE, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
