package model

type SureSureBank struct {
	ID            int    `json:"id" db:"id"`
	UserID        int    `json:"user_id" db:"userid"`
	BankCode      string `json:"bank_code" db:"bankcode"`
	PromptPayType string `json:"prompt_pay_type" db:"promptpaytype"`
	AccountNo     string `json:"account_no" db:"accountno"`
	AccountType   string `json:"account_type" db:"accounttype"`
	NameTH        string `json:"name_th" db:"nameth"`
	NameEN        string `json:"name_en" db:"nameen"`
	IsActive      bool   `json:"is_active" db:"isactive"`
	CreatedDate   string `json:"created_date" db:"createddate"`
	UpdatedDate   string `json:"updated_date" db:"updateddate"`
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
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
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
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
 FROM SureSureBank WHERE UserID = $1`

//ดึงข้อมูลตาม UserID

var SQL_BANK_ACCOUNT_DELETE = "DELETE FROM SureSureBank WHERE ID = $1"

//ลบข้อมูลตาม ID
