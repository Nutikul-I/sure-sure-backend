package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
)

type PackageRepository struct {
	DB *sql.DB
}

func GetPackageAll() ([]model.SureSurePackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSurePackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_PACKAGE_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSurePackage{}, err
	}

	var packages []model.SureSurePackage
	err = scan.Rows(&packages, rows)

	defer rows.Close()
	log.Infof("packages: %d", len(packages))
	return packages, nil
}

func GetPackageByID(id int) (model.SureSurePackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return model.SureSurePackage{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_PACKAGE_GET_BYID, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.SureSurePackage{}, err
	}
	var pkg model.SureSurePackage
	err = scan.Row(&pkg, rows)
	defer rows.Close()
	if pkg.ID == 0 {
		log.Errorf("Not Found: %v", err)
		return model.SureSurePackage{}, err
	}

	return pkg, nil
}

func CreatePackage(pkg model.SureSurePackage) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return 0, err
	}
	// Build query dynamically
	query := "INSERT INTO SureSurePackage ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	if pkg.PackageName != "" {
		query += "PackageName, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackageName))
		counter++
	}

	if pkg.PackagePrice != 0 {
		query += "PackagePrice, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackagePrice))
		counter++
	}

	if pkg.QuotaLimit != 0 {
		query += "QuotaLimit, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.QuotaLimit))
		counter++
	}

	if pkg.Amount != 0 {
		query += "Amount, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Amount))
		counter++
	}

	if pkg.Ordered != 0 {
		query += "Ordered, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Ordered))
		counter++
	}

	if pkg.Duration != 0 {
		query += "Duration, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Duration))
		counter++
	}

	if pkg.IsActive {
		query += "IsActive, "
		values += fmt.Sprintf("@p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.IsActive))
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
	// Retrieve the last inserted ID
	// lastInsertedID, err := result.LastInsertId()
	// if err != nil {
	// 	log.Errorf("Error retrieving last insert ID: %v", err)
	// 	return 0, err
	// }

	return 0, nil
}

func UpdatePackage(pkg model.SureSurePackage) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// Initialize query parts
	query := "UPDATE SureSurePackage SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if pkg.PackageName != "" {
		query += fmt.Sprintf("PackageName = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackageName))
		counter++
	}

	if pkg.PackagePrice != 0 {
		query += fmt.Sprintf("PackagePrice = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.PackagePrice))
		counter++
	}

	if pkg.QuotaLimit != 0 {
		query += fmt.Sprintf("QuotaLimit = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.QuotaLimit))
		counter++
	}

	if pkg.Amount != 0 {
		query += fmt.Sprintf("Amount = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Amount))
		counter++
	}

	if pkg.Ordered != 0 {
		query += fmt.Sprintf("Ordered = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Ordered))
		counter++
	}

	if pkg.Duration != 0 {
		query += fmt.Sprintf("Duration = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.Duration))
		counter++
	}

	query += fmt.Sprintf("IsActive = @p%d, ", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.IsActive))
	counter++

	if pkg.UpdatedDate != "" {
		query += fmt.Sprintf("UpdatedDate = @p%d, ", counter)
		params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.UpdatedDate))
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = @p" + fmt.Sprintf("%d", counter)
	params = append(params, sql.Named(fmt.Sprintf("p%d", counter), pkg.ID))

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeletePackage(id int) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_PACKAGE_DELETE, sql.Named("ID", id))
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}
