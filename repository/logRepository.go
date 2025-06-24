package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
)

func GetLogAll() ([]model.SureSureLog, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureLog{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_LOG_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureLog{}, err
	}

	var logs []model.SureSureLog
	err = scan.Rows(&logs, rows)

	defer rows.Close()
	log.Infof("logs: %d", len(logs))
	return logs, nil
}

func GetLogByID(id int) ([]model.SureSureLog, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureLog{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_LOG_GET_BYID, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return []model.SureSureLog{}, err
	}
	var log_ []model.SureSureLog
	err = scan.Rows(&log_, rows)
	defer rows.Close()

	return log_, nil
}

func CreateLog(log_ model.SureSureLog) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	// Build query dynamically
	query := "INSERT INTO SureSureLog ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	if log_.Action != "" {
		query += "Action, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.Action))
		counter++
	}
	if log_.MethodName != "" {
		query += "MethodName, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.MethodName))
		counter++
	}
	if log_.UserID != 0 {
		query += "UserID, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.UserID))
		counter++
	}
	if log_.NameTH != "" {
		query += "NameTH, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.NameTH))
		counter++
	}
	if log_.DataRequest != "" {
		query += "DataRequest, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.DataRequest))
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

func UpdateLog(log_ model.SureSureLog) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}

	// Initialize query parts
	query := "UPDATE SureSureUser SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if log_.Action != "" {
		query += fmt.Sprintf("Action = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.Action))
		counter++
	}

	if log_.MethodName != "" {
		query += fmt.Sprintf("MethodName = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.MethodName))
		counter++
	}

	if log_.UserID != 0 {
		query += fmt.Sprintf("UserID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.UserID))
		counter++
	}

	if log_.NameTH != "" {
		query += fmt.Sprintf("NameTH = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.NameTH))
		counter++
	}

	if log_.DataRequest != "" {
		query += fmt.Sprintf("DataRequest = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.DataRequest))
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = @p" + fmt.Sprintf("%d", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), log_.ID))

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteLog(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_LOG_DELETE, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
