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
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_GET_BYID, id)
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
	rows, err := conn.QueryContext(ctx, model.SQL_ORDER_PACKAGE_GET_BYREFNO, RefNo)
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

	// Simplified query with all required fields - ใช้ PostgreSQL syntax
	query := `INSERT INTO SureSureOrderPackage 
		(RefNo, UserID, PackageID, Price, Status, CreatedDate, UpdatedDate) 
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
		RETURNING ID`

	params := []interface{}{
		pkg.RefNo,
		pkg.UserID,
		pkg.PackageID,
		pkg.Price,
		pkg.Status,
	}

	log.Infof("CreateOrderPackage query: %s", query)
	log.Infof("CreateOrderPackage params: %v", params)

	var newID int
	err = conn.QueryRowContext(ctx, query, params...).Scan(&newID)
	if err != nil {
		log.Errorf("Error executing create query: %v", err)
		return 0, err
	}

	log.Infof("Successfully created order package with ID: %d", newID)

	if pkg.Status == "SUCCESS" {
		_, err = conn.ExecContext(ctx, model.SQL_PACKAGE_FROM_ORDER_UPDATE, pkg.RefNo)
		if err != nil {
			log.Errorf("Error executing package update query: %v", err)
			return 0, err
		}
		_, err = conn.ExecContext(ctx, model.SQL_USER_FROM_ORDER_UPDATE, pkg.RefNo)
		if err != nil {
			log.Errorf("Error executing user update query: %v", err)
			return 0, err
		}
	}

	return newID, nil
}

func UpdateOrderPackage(pkg model.SureSureOrderPackage) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}

	// Initialize query parts - ใช้ PostgreSQL syntax
	query := "UPDATE SureSureOrderPackage SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if pkg.UserID != 0 {
		query += fmt.Sprintf("UserID = $%d, ", counter)
		params = append(params, pkg.UserID)
		counter++
	}

	if pkg.PackageID != 0 {
		query += fmt.Sprintf("PackageID = $%d, ", counter)
		params = append(params, pkg.PackageID)
		counter++
	}

	if pkg.Price != 0 {
		query += fmt.Sprintf("Price = $%d, ", counter)
		params = append(params, pkg.Price)
		counter++
	}

	if pkg.Status != "" {
		query += fmt.Sprintf("Status = $%d, ", counter)
		params = append(params, pkg.Status)
		counter++
	}

	// Always update UpdatedDate to current timestamp
	query += "UpdatedDate = CURRENT_TIMESTAMP, "

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE RefNo = $%d", counter)
	params = append(params, pkg.RefNo)

	log.Infof("UpdateOrderPackage query: %s", query)
	log.Infof("UpdateOrderPackage params: %v", params)

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing update query: %v", err)
		return err
	}

	if pkg.Status == "SUCCESS" {
		_, err = conn.ExecContext(ctx, model.SQL_PACKAGE_FROM_ORDER_UPDATE, pkg.RefNo)
		if err != nil {
			log.Errorf("Error executing package update query: %v", err)
			return err
		}
		_, err = conn.ExecContext(ctx, model.SQL_USER_FROM_ORDER_UPDATE, pkg.RefNo)
		if err != nil {
			log.Errorf("Error executing user update query: %v", err)
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
	_, err = conn.ExecContext(ctx, model.SQL_ORDER_PACKAGE_DELETE, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	return nil
}
