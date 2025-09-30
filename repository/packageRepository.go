package repository

import (
	"context"
	"database/sql"
	"fmt"

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

	log.Infof("Executing query: %s", model.SQL_PACKAGE_GET)
	rows, err := conn.QueryContext(ctx, model.SQL_PACKAGE_GET)
	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSurePackage{}, err
	}
	defer rows.Close()

	var packages []model.SureSurePackage

	// Manual scanning for debugging
	for rows.Next() {
		var pkg model.SureSurePackage
		err := rows.Scan(
			&pkg.ID,
			&pkg.PackageName,
			&pkg.PackagePrice,
			&pkg.QuotaLimit,
			&pkg.Amount,
			&pkg.Ordered,
			&pkg.Duration,
			&pkg.IsActive,
			&pkg.CreatedDate,
			&pkg.UpdatedDate,
		)
		if err != nil {
			log.Errorf("Error scanning row: %v", err)
			continue
		}
		packages = append(packages, pkg)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("Rows error: %v", err)
		return []model.SureSurePackage{}, err
	}

	log.Infof("Total packages found: %d", len(packages))
	return packages, nil
}

func GetPackageByID(id int) (model.SureSurePackage, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return model.SureSurePackage{}, err
	}
	row := conn.QueryRowContext(ctx, model.SQL_PACKAGE_GET_BYID, id)

	var pkg model.SureSurePackage
	err = row.Scan(
		&pkg.ID,
		&pkg.PackageName,
		&pkg.PackagePrice,
		&pkg.QuotaLimit,
		&pkg.Amount,
		&pkg.Ordered,
		&pkg.Duration,
		&pkg.IsActive,
		&pkg.CreatedDate,
		&pkg.UpdatedDate,
	)
	if err != nil {
		log.Errorf("Error scanning row: %v", err)
		return model.SureSurePackage{}, err
	}

	return pkg, nil
}

func CreatePackage(pkg model.SureSurePackage) (int, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		log.Errorf("Database ping failed: %v", err)
		return 0, err
	}

	log.Infof("Creating package: %+v", pkg)

	// Simplified query with all required fields
	query := `INSERT INTO SureSurePackage 
		(PackageName, PackagePrice, QuotaLimit, Duration, IsActive, Amount, Ordered, CreatedDate, UpdatedDate) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
		RETURNING ID`

	params := []interface{}{
		pkg.PackageName,
		pkg.PackagePrice,
		pkg.QuotaLimit,
		pkg.Duration,
		pkg.IsActive,
		pkg.Amount,
		pkg.Ordered,
	}

	log.Infof("CreatePackage query: %s", query)
	log.Infof("CreatePackage params: %v", params)

	var newID int
	err = conn.QueryRowContext(ctx, query, params...).Scan(&newID)
	if err != nil {
		log.Errorf("Error executing create query: %v", err)
		return 0, err
	}

	log.Infof("Successfully created package with ID: %d", newID)
	return newID, nil
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
		query += fmt.Sprintf("PackageName = $%d, ", counter)
		params = append(params, pkg.PackageName)
		counter++
	}

	if pkg.PackagePrice != 0 {
		query += fmt.Sprintf("PackagePrice = $%d, ", counter)
		params = append(params, pkg.PackagePrice)
		counter++
	}

	if pkg.QuotaLimit != 0 {
		query += fmt.Sprintf("QuotaLimit = $%d, ", counter)
		params = append(params, pkg.QuotaLimit)
		counter++
	}

	if pkg.Amount != 0 {
		query += fmt.Sprintf("Amount = $%d, ", counter)
		params = append(params, pkg.Amount)
		counter++
	}

	if pkg.Ordered != 0 {
		query += fmt.Sprintf("Ordered = $%d, ", counter)
		params = append(params, pkg.Ordered)
		counter++
	}

	if pkg.Duration != 0 {
		query += fmt.Sprintf("Duration = $%d, ", counter)
		params = append(params, pkg.Duration)
		counter++
	}

	// Always update IsActive
	query += fmt.Sprintf("IsActive = $%d, ", counter)
	params = append(params, pkg.IsActive)
	counter++

	// Add UpdatedDate
	query += "UpdatedDate = CURRENT_TIMESTAMP, "

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE ID = $%d", counter)
	params = append(params, pkg.ID)

	log.Infof("UpdatePackage query: %s", query)
	log.Infof("UpdatePackage params: %v", params)

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
	_, err = conn.ExecContext(ctx, model.SQL_PACKAGE_DELETE, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	return nil
}
