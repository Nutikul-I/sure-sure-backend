package model

import "time"

type SureSureBank struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	BankCode      string    `json:"bank_code"`
	PromptPayType string    `json:"prompt_pay_type"`
	AccountNo     string    `json:"account_no"`
	AccountType   string    `json:"account_type"`
	NameTH        string    `json:"name_th"`
	NameEN        string    `json:"name_en"`
	IsActive      bool      `json:"is_active"`
	CreatedDate   time.Time `json:"created_date"`
	UpdatedDate   time.Time `json:"updated_date"`
}

var SQL_BANK_ACCOUNT_GET = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (BankCode,'') AS BankCode,
COALESCE (PromptPayType,'') AS PromptPayType,
COALESCE (AccountNo,'') AS AccountNo,
COALESCE (AccountType,'') AS AccountType,
COALESCE (NameTH,'') AS NameTH,
COALESCE (NameEN,'') AS NameEN,
COALESCE (IsActive,0) AS IsActive,
COALESCE (CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS CreatedDate,
COALESCE (UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS UpdatedDate
 FROM SureSureBank`

// ใช้สำหรับดึงข้อมูลบัญชีธนาคารทั้งหมด
var SQL_BANK_ACCOUNT_GET_BYID = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (BankCode,'') AS BankCode,
COALESCE (PromptPayType,'') AS PromptPayType,
COALESCE (AccountNo,'') AS AccountNo,
COALESCE (AccountType,'') AS AccountType,
COALESCE (NameTH,'') AS NameTH,
COALESCE (NameEN,'') AS NameEN,
COALESCE (IsActive,0) AS IsActive,
COALESCE (CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS CreatedDate,
COALESCE (UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS UpdatedDate
 FROM SureSureBank WHERE UserID = $1`

//ดึงข้อมูลตาม UserID

var SQL_BANK_ACCOUNT_DELETE = "DELETE FROM SureSureBank WHERE ID = $1"

//ลบข้อมูลตาม ID
