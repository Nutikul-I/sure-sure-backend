package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
)

type OrderPackageRepository struct {
	DB *sql.DB
}

func GetOrderPackageAll() ([]model.SureSureOrderPackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureOrderPackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureOrderPackage{}, err
	}

	var orderPackages []model.SureSureOrderPackage
	err = scan.Rows(&orderPackages, rows)

	defer rows.Close()
	log.Infof("orderPackages: %d", len(orderPackages))
	return orderPackages, nil
}
func GetOrderPackagPending() ([]model.SureSureOrderPackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureOrderPackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_PENDING_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureOrderPackage{}, err
	}

	var orderPackages []model.SureSureOrderPackage
	err = scan.Rows(&orderPackages, rows)

	defer rows.Close()
	log.Infof("orderPackages: %d", len(orderPackages))
	return orderPackages, nil
}

func GetOrderPackageByID(id int) ([]model.SureSureOrderPackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureOrderPackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_GET_BYID, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return []model.SureSureOrderPackage{}, err
	}
	var pkg []model.SureSureOrderPackage
	err = scan.Rows(&pkg, rows)
	defer rows.Close()

	return pkg, nil
}

func GetOrderPackageByRefNo(RefNo string) (model.SureSureOrderPackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return model.SureSureOrderPackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_GET_BYREFNO, sql.Named("RefNo", RefNo))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.SureSureOrderPackage{}, err
	}
	var pkg model.SureSureOrderPackage
	err = scan.Row(&pkg, rows)
	defer rows.Close()
	return pkg, nil
}

func CreateOrderPackage(pkg model.SureSureOrderPackage) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	log.Info("CreateOrderPackage")
	// Build query dynamically
	query := "INSERT INTO SureSureOrderPackage ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	if pkg.RefNo != "" {
		query += "RefNo, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.RefNo))
		counter++
	}

	if pkg.UserID != 0 {
		query += "UserID, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.UserID))
		counter++
	}

	if pkg.PackageID != 0 {
		query += "PackageID, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackageID))
		counter++
	}

	if pkg.Price != 0 {
		query += "Price, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Price))
		counter++
	}

	if pkg.Status != "" {
		query += "Status, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Status))
		counter++
	}

	if pkg.CreatedDate != "" {
		query += "CreatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.CreatedDate))
		counter++
	}

	if pkg.UpdatedDate != "" {
		query += "UpdatedDate, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.UpdatedDate))
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
	if pkg.Status == "SUCCESS" {
		_, err = conn.ExecContext(ctx, model.SQL_PACKAGE_FROM_ORDER_UPDATE, sql.Named("RefNo", pkg.RefNo))
		if err != nil {
			log.Errorf("Error executing query: %v", err)
			return 0, err
		}
		_, err = conn.ExecContext(ctx, model.SQL_USER_FROM_ORDER_UPDATE, sql.Named("RefNo", pkg.RefNo))
		if err != nil {
			log.Errorf("Error executing query: %v", err)
			return 0, err
		}

	}
	// Retrieve the last inserted ID
	// lastInsertedID, err := result.LastInsertId()
	// if err != nil {
	// 	log.Errorf("Error retrieving last insert ID: %v", err)
	// 	return 0, err
	// }

	return 0, nil
}

func UpdateOrderPackage(pkg model.SureSureOrderPackage) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// Initialize query parts
	query := "UPDATE SureSureOrderPackage SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values

	if pkg.UserID != 0 {
		query += fmt.Sprintf("UserID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.UserID))
		counter++
	}

	if pkg.PackageID != 0 {
		query += fmt.Sprintf("PackageID = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackageID))
		counter++
	}

	if pkg.Price != 0 {
		query += fmt.Sprintf("Price = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Price))
		counter++
	}

	if pkg.Status != "" {
		query += fmt.Sprintf("Status = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Status))
		counter++
	}

	if pkg.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.UpdatedDate))
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE RefNo = @p" + fmt.Sprintf("%d", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.RefNo))

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	if pkg.Status == "SUCCESS" {
		_, err = conn.ExecContext(ctx, model.SQL_PACKAGE_FROM_ORDER_UPDATE, sql.Named("RefNo", pkg.RefNo))
		if err != nil {
			log.Errorf("Error executing query: %v", err)
			return err
		}
		_, err = conn.ExecContext(ctx, model.SQL_USER_FROM_ORDER_UPDATE, sql.Named("RefNo", pkg.RefNo))
		if err != nil {
			log.Errorf("Error executing query: %v", err)
			return err
		}
	}

	return nil
}

func DeleteOrderPackage(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_DELETE, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
