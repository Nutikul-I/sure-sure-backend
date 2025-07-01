package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/blockloop/scan"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type UserRepository struct {
	DB *sql.DB
}

func GetOrCreateUser(user model.SureSureUser) (model.SureSureUser, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		log.Errorf("Database ping error: %v", err)
		return model.SureSureUser{}, err
	}

	log.Infof("Checking user with Username: %s and Password: %s", user.Username, user.Password)

	rows, err := conn.QueryContext(ctx, model.SQL_USER_GET_BY_USERNAME,
		user.Username,
		user.Password,
	)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.SureSureUser{}, err
	}
	var result model.SureSureUser
	err = scan.Row(&result, rows)
	defer rows.Close()
	log.Infof("result: %v", result)
	// Step 2: Scan the result
	if result.ID == 0 && user.UserType == "merchant" {
		log.Infof("No existing user found, creating a new user.")
		userID, err := CreateUser(user)
		// if err != nil {
		// 	log.Errorf("CreateUser error: %v", err)
		// 	return model.SureSureUser{}, err
		// }
		log.Infof("User created with ID: %d", userID)

		result, err = GetUserByID(userID)
		if err != nil {
			log.Errorf("GetUserByID error: %v", err)
			return model.SureSureUser{}, err
		}
		log.Infof("New user retrieved: %v", result)
	}

	// Step 3: Generate JWT token
	token, err := generateJWT(result.ID, result.Username)
	if err != nil {
		log.Errorf("JWT generation error: %v", err)
		return model.SureSureUser{}, err
	}
	log.Infof("JWT token generated: %s", token)
	result.Token = token

	// Step 4: Update user with the new token
	err = UpdateUser(result)
	if err != nil {
		log.Errorf("UpdateUser error: %v", err)
		return model.SureSureUser{}, err
	}
	log.Infof("User updated with new token: %v", result)

	return result, nil
}

func GetUserAll() ([]model.SureSureUser, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.SureSureUser{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_USER_GET)

	if err != nil {
		log.Errorf("ERROR: %#v", err)
		return []model.SureSureUser{}, err
	}

	var users []model.SureSureUser
	err = scan.Rows(&users, rows)

	defer rows.Close()
	log.Infof("users: %d", len(users))
	return users, nil
}

func GetUserByID(id string) (model.SureSureUser, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return model.SureSureUser{}, err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_USER_GET_BYID, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return model.SureSureUser{}, err
	}
	var user model.SureSureUser
	err = scan.Row(&user, rows)
	defer rows.Close()
	if user.ID == 0 {
		log.Errorf("Not Found: %v", err)
		return model.SureSureUser{}, err
	}

	return user, nil
}

func CreateUser(user model.SureSureUser) (string, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return "", err
	}
	uid := util.GenerateRandomText(16)

	// Build query dynamically
	query := "INSERT INTO SureSureUser ("
	values := "VALUES ("
	params := []interface{}{}
	counter := 1

	query += "UID, "
	values += fmt.Sprintf("$%d, ", counter)
	params = append(params, uid)
	counter++
	if user.MerchantID != 0 {
		query += "MerchantID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.MerchantID)
		counter++
	}
	if user.PackageID != 0 {
		query += "PackageID, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.PackageID)
		counter++
	}
	if user.UserType != "" {
		query += "UserType, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.UserType)
		counter++
	} else {
		query += "UserType, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, "merchant-register")
		counter++
	}
	if user.Picture != "" {
		query += "Picture, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Picture)
		counter++
	}
	if user.NameTH != "" {
		query += "NameTH, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.NameTH)
		counter++
	}
	if user.NameEN != "" {
		query += "NameEN, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.NameEN)
		counter++
	}
	if user.Phone != "" {
		query += "Phone, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Phone)
		counter++
	}
	if user.Website != "" {
		query += "Website, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Website)
		counter++
	}
	if user.UserRole != "" {
		query += "UserRole, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.UserRole)
		counter++
	} else {
		query += "UserRole, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, "merchant-register")
		counter++
	}
	if user.Address != "" {
		query += "Address, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Address)
		counter++
	}
	if user.Email != "" {
		query += "Email, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Email)
		counter++
	}
	if user.Username != "" {
		query += "Username, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Username)
		counter++
	}
	if user.Password != "" {
		query += "Password, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Password)
		counter++
	}
	if user.IsActive != 0 {
		query += "IsActive, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.IsActive)
		counter++
	}
	if user.StoreName != "" {
		query += "StoreName, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.StoreName)
		counter++
	}
	if user.StoreCategoryType != "" {
		query += "StoreCategoryType, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.StoreCategoryType)
		counter++
	}
	if user.StorePhone != "" {
		query += "StorePhone, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.StorePhone)
		counter++
	}
	if user.StoreEmail != "" {
		query += "StoreEmail, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.StoreEmail)
		counter++
	}
	if user.QuotaLeft > 0 {
		query += "QuotaLeft, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.QuotaLeft)
		counter++
	}
	if user.QuotaALL > 0 {
		query += "QuotaALL, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.QuotaALL)
		counter++
	}
	if user.Step > 0 {
		query += "Step, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.Step)
		counter++
	}
	if user.PackageChangeDate != "" {
		query += "PackageChangeDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.PackageChangeDate)
		counter++
	}
	if user.BillDate != "" {
		query += "BillDate, "
		values += fmt.Sprintf("$%d, ", counter)
		params = append(params, user.BillDate)
		counter++
	}
	query += "AccessToken, "
	values += fmt.Sprintf("$%d, ", counter)
	params = append(params, util.GenerateRandomText(24))
	counter++

	query = query[:len(query)-2] + ") "
	values = values[:len(values)-2] + ")"
	finalQuery := query + values + " RETURNING ID"

	log.Infof("finalQuery: %v", finalQuery)
	result := conn.QueryRowContext(ctx, finalQuery, params...)
	// Retrieve the last inserted ID
	var lastInsertedID int64
	if err := result.Scan(&lastInsertedID); err != nil {
		log.Errorf("Error retrieving last inserted ID: %v", err)
		return "", err
	}

	return uid, nil
}

func UpdateUser(user model.SureSureUser) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	// CHECK DUPLICATE
	if user.StoreName != "" && user.StorePhone != "" && user.StoreEmail != "" {
		rows, err := conn.QueryContext(ctx, model.SQL_USER_GET_DUPLICATE, user.StoreName, user.StorePhone, user.StoreEmail, user.ID)

		if err != nil {
			log.Errorf("Error executing query: %v", err)
			return err
		}
		var duplicate int
		err = scan.Row(&duplicate, rows)
		defer rows.Close()
		if duplicate != 0 {
			return fmt.Errorf("duplicate store")
		}
	}

	// Initialize query parts
	query := "UPDATE SureSureUser SET "
	params := []interface{}{}
	counter := 1

	// Dynamically add fields and values
	if user.PackageID != 0 {
		query += fmt.Sprintf("PackageID = $%d, ", counter)
		params = append(params, user.PackageID)
		counter++
	}
	if user.Picture != "" {
		query += fmt.Sprintf("Picture = $%d, ", counter)
		params = append(params, user.Picture)
		counter++
	}
	if user.NameTH != "" {
		query += fmt.Sprintf("NameTH = $%d, ", counter)
		params = append(params, user.NameTH)
		counter++
	}
	if user.NameEN != "" {
		query += fmt.Sprintf("NameEN = $%d, ", counter)
		params = append(params, user.NameEN)
		counter++
	}
	if user.Phone != "" {
		query += fmt.Sprintf("Phone = $%d, ", counter)
		params = append(params, user.Phone)
		counter++
	}
	if user.Website != "" {
		query += fmt.Sprintf("Website = $%d, ", counter)
		params = append(params, user.Website)
		counter++
	}
	if user.Address != "" {
		query += fmt.Sprintf("Address = $%d, ", counter)
		params = append(params, user.Address)
		counter++
	}
	if user.Email != "" {
		query += fmt.Sprintf("Email = $%d, ", counter)
		params = append(params, user.Email)
		counter++
	}
	if user.Username != "" {
		query += fmt.Sprintf("Username = $%d, ", counter)
		params = append(params, user.Username)
		counter++
	}
	if user.Password != "" {
		query += fmt.Sprintf("Password = $%d, ", counter)
		params = append(params, user.Password)
		counter++
	}
	if user.ID != 0 {
		query += fmt.Sprintf("IsActive = $%d, ", counter)
		// ส่ง int ตรงๆ ไม่ต้องแปลง
		params = append(params, user.IsActive)
		counter++
	}
	if user.StoreName != "" {
		query += fmt.Sprintf("StoreName = $%d, ", counter)
		params = append(params, user.StoreName)
		counter++
	}
	if user.StoreCategoryType != "" {
		query += fmt.Sprintf("StoreCategoryType = $%d, ", counter)
		params = append(params, user.StoreCategoryType)
		counter++
	}
	if user.StorePhone != "" {
		query += fmt.Sprintf("StorePhone = $%d, ", counter)
		params = append(params, user.StorePhone)
		counter++
	}
	if user.StoreEmail != "" {
		query += fmt.Sprintf("StoreEmail = $%d, ", counter)
		params = append(params, user.StoreEmail)
		counter++
	}
	if user.QuotaLeft > 0 {
		query += fmt.Sprintf("QuotaLeft = $%d, ", counter)
		params = append(params, user.QuotaLeft)
		counter++
	}
	if user.QuotaALL > 0 {
		query += fmt.Sprintf("QuotaALL = $%d, ", counter)
		params = append(params, user.QuotaALL)
		counter++
	}
	if user.Step > 0 {
		query += fmt.Sprintf("Step = $%d, ", counter)
		params = append(params, user.Step)
		counter++
	}
	if user.PackageChangeDate != "" {
		query += fmt.Sprintf("PackageChangeDate = $%d, ", counter)
		params = append(params, user.PackageChangeDate)
		counter++
	}
	if user.BillDate != "" {
		query += fmt.Sprintf("BillDate = $%d, ", counter)
		params = append(params, user.BillDate)
		counter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + " WHERE ID = $" + fmt.Sprintf("%d", counter)
	params = append(params, user.ID)

	// Execute query
	_, err = conn.ExecContext(ctx, query, params...)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return err
	}
	rows, err := conn.QueryContext(ctx, model.SQL_USER_DELETE, id)
	if err != nil {
		log.Errorf("Error executing query: %v", err)
		return err
	}
	defer rows.Close()
	return nil
}

func GetCategoryAll() ([]model.MerchantCategory, error) {
	conn := ConnectDB()
	ctx := context.Background()
	err := conn.PingContext(ctx)
	if err != nil {
		return []model.MerchantCategory{}, err
	}
	log.Info("GetCategoryAll")
	log.Infof("SQL Query: %s", model.SQL_CATEGORY_GET)

	rows, err := conn.QueryContext(ctx, model.SQL_CATEGORY_GET)
	if err != nil {
		log.Errorf("Query ERROR: %#v", err)
		return []model.MerchantCategory{}, err
	}
	defer rows.Close()

	var category []model.MerchantCategory
	err = scan.Rows(&category, rows)
	if err != nil {
		log.Errorf("Scan ERROR: %#v", err)
		return []model.MerchantCategory{}, err
	}

	log.Infof("category count: %d", len(category))
	return category, nil
}

func generateJWT(userID int, username string) (string, error) {
	const secretKey = "jdnfksdmfksd"

	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Infof("tokenString: %v", err)
		return "", err
	}
	log.Infof("tokenString: %s", tokenString)
	return tokenString, nil
}
