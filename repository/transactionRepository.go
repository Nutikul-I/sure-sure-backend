package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
)

type SureSureTransactionRepository struct {
	DB *sql.DB
}

func GetTransactionAll() ([]model.SureSureTransaction, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureTransaction{}, err
	}
	rows, err := conn.Query(model.SQL_GET_ALL_TRANSACTIONS)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureTransaction{}, err
	}

	var transactions []model.SureSureTransaction
	err = scan.Rows(&transactions, rows)

	defer rows.Close()
	log.Infof("transactions: %d", len(transactions))
	return transactions, nil
}

func GetTransactionByID(id int) ([]model.SureSureTransaction, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureTransaction{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_GET_TRANSACTION_BY_ID, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return []model.SureSureTransaction{}, err
	}
	var t []model.SureSureTransaction
	err = scan.Rows(&t, rows)
	defer rows.Close()
	return t, nil
}

func CreateTransaction(t model.SureSureTransaction) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	// Build query dynamically
	query := "INSERT INTO SureSureTransaction ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	if t.UserID != 0 {
		query += "UserID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.UserID)
		counter++
	}
	if t.QrCode != "" {
		query += "QrCode, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.QrCode)
		counter++
	}
	if t.RefNo != "" {
		query += "RefNo, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.RefNo)
		counter++
	}
	if t.LineUserID != "" {
		query += "LineUserID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.LineUserID)
		counter++
	}
	if t.LineGroupID != "" {
		query += "LineGroupID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.LineGroupID)
		counter++
	}
	if t.Amount != 0 {
		query += "Amount, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.Amount)
		counter++
	}
	if t.CSTID != "" {
		query += "CSTID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.CSTID)
		counter++
	}
	if t.RQUID != "" {
		query += "RQUID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.RQUID)
		counter++
	}
	if t.TXID != "" {
		query += "TXID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.TXID)
		counter++
	}
	if t.SenderBankCode != "" {
		query += "SenderBankCode, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.SenderBankCode)
		counter++
	}
	if t.SenderAccountNo != "" {
		query += "SenderAccountNo, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.SenderAccountNo)
		counter++
	}
	if t.SenderName != "" {
		query += "SenderName, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.SenderName)
		counter++
	}
	if t.SenderName2 != "" {
		query += "SenderName2, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.SenderName2)
		counter++
	}
	if t.ReceiveBankCode != "" {
		query += "ReceiveBankCode, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.ReceiveBankCode)
		counter++
	}
	if t.ReceiveAccountNo != "" {
		query += "ReceiveAccountNo, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.ReceiveAccountNo)
		counter++
	}

	if t.ProxyAccountNo != "" {
		query += "ProxyAccountNo, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.ProxyAccountNo)
		counter++
	}
	if t.Ref1 != "" {
		query += "Ref1, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.Ref1)
		counter++
	}
	if t.Ref2 != "" {
		query += "Ref2, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.Ref2)
		counter++
	}
	if t.ReceiveName != "" {
		query += "ReceiveName, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.ReceiveName)
		counter++
	}
	if t.ReceiveName2 != "" {
		query += "ReceiveName2, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.ReceiveName2)
		counter++
	}
	if t.Message != "" {
		query += "Message, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.Message)
		counter++
	}
	if t.StatusCode != "" {
		query += "StatusCode, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.StatusCode)
		counter++
	}
	if t.Status != "" {
		query += "Status, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.Status)
		counter++
	}
	if t.TransDate != "" {
		query += "TransDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.TransDate)
		counter++
	}
	if t.TransTime != "" {
		query += "TransTime, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.TransTime)
		counter++
	}
	if t.CreatedDate != "" {
		query += "CreatedDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.CreatedDate)
		counter++
	}
	if t.UpdatedDate != "" {
		query += "UpdatedDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, t.UpdatedDate)
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

func UpdateTransaction(t model.SureSureTransaction) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}

	// Initialize query parts
	query := "UPDATE SureSureTransaction SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if t.UserID != 0 {
		query += fmt.Sprintf("UserID = $%d, ", counter)
		params = append(params, t.UserID)
		counter++
	}
	if t.QrCode != "" {
		query += fmt.Sprintf("QrCode = $%d, ", counter)
		params = append(params, t.QrCode)
		counter++
	}
	if t.RefNo != "" {
		query += fmt.Sprintf("RefNo = $%d, ", counter)
		params = append(params, t.RefNo)
		counter++
	}
	if t.LineUserID != "" {
		query += fmt.Sprintf("LineUserID = $%d, ", counter)
		params = append(params, t.LineUserID)
		counter++
	}
	if t.LineGroupID != "" {
		query += fmt.Sprintf("LineGroupID = $%d, ", counter)
		params = append(params, t.LineGroupID)
		counter++
	}
	if t.Amount != 0 {
		query += fmt.Sprintf("Amount = $%d, ", counter)
		params = append(params, t.Amount)
		counter++
	}
	if t.CSTID != "" {
		query += fmt.Sprintf("CSTID = $%d, ", counter)
		params = append(params, t.CSTID)
		counter++
	}
	if t.RQUID != "" {
		query += fmt.Sprintf("RQUID = $%d, ", counter)
		params = append(params, t.RQUID)
		counter++
	}
	if t.TXID != "" {
		query += fmt.Sprintf("TXID = $%d, ", counter)
		params = append(params, t.TXID)
		counter++
	}
	if t.SenderBankCode != "" {
		query += fmt.Sprintf("SenderBankCode = $%d, ", counter)
		params = append(params, t.SenderBankCode)
		counter++
	}
	if t.SenderAccountNo != "" {
		query += fmt.Sprintf("SenderAccountNo = $%d, ", counter)
		params = append(params, t.SenderAccountNo)
		counter++
	}
	if t.SenderName != "" {
		query += fmt.Sprintf("SenderName = $%d, ", counter)
		params = append(params, t.SenderName)
		counter++
	}
	if t.SenderName2 != "" {
		query += fmt.Sprintf("SenderName2 = $%d, ", counter)
		params = append(params, t.SenderName2)
		counter++
	}
	if t.ReceiveBankCode != "" {
		query += fmt.Sprintf("ReceiveBankCode = $%d, ", counter)
		params = append(params, t.ReceiveBankCode)
		counter++
	}
	if t.ReceiveAccountNo != "" {
		query += fmt.Sprintf("ReceiveAccountNo = $%d, ", counter)
		params = append(params, t.ReceiveAccountNo)
		counter++
	}
	if t.ProxyAccountNo != "" {
		query += fmt.Sprintf("ProxyAccountNo = $%d, ", counter)
		params = append(params, t.ProxyAccountNo)
		counter++
	}
	if t.Ref1 != "" {
		query += fmt.Sprintf("Ref1 = $%d, ", counter)
		params = append(params, t.Ref1)
		counter++
	}
	if t.Ref2 != "" {
		query += fmt.Sprintf("Ref2 = $%d, ", counter)
		params = append(params, t.Ref2)
		counter++
	}
	if t.ReceiveName != "" {
		query += fmt.Sprintf("ReceiveName = $%d, ", counter)
		params = append(params, t.ReceiveName)
		counter++
	}
	if t.ReceiveName2 != "" {
		query += fmt.Sprintf("ReceiveName2 = $%d, ", counter)
		params = append(params, t.ReceiveName2)
		counter++
	}
	if t.Message != "" {
		query += fmt.Sprintf("Message = $%d, ", counter)
		params = append(params, t.Message)
		counter++
	}
	if t.StatusCode != "" {
		query += fmt.Sprintf("StatusCode = $%d, ", counter)
		params = append(params, t.StatusCode)
		counter++
	}
	if t.Status != "" {
		query += fmt.Sprintf("Status = $%d, ", counter)
		params = append(params, t.Status)
		counter++
	}
	if t.TransDate != "" {
		query += fmt.Sprintf("TransDate = $%d, ", counter)
		params = append(params, t.TransDate)
		counter++
	}
	if t.TransTime != "" {
		query += fmt.Sprintf("TransTime = $%d, ", counter)
		params = append(params, t.TransTime)
		counter++
	}
	if t.CreatedDate != "" {
		query += fmt.Sprintf("CreatedDate = $%d, ", counter)
		params = append(params, t.CreatedDate)
		counter++
	}
	if t.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = $%d, ", counter)
		params = append(params, t.UpdatedDate)
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = $" + fmt.Sprintf("%d", counter)
	params = append(params, t.ID)

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteTransaction(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_DELETE_TRANSACTION, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
