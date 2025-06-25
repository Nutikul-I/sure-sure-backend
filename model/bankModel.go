package model

type SureSureBank struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	BankCode      string `json:"bank_code"`
	PromptPayType string `json:"prompt_pay_type"`
	AccountNo     string `json:"account_no"`
	AccountType   string `json:"account_type"`
	NameTH        string `json:"name_th"`
	NameEN        string `json:"name_en"`
	IsActive      bool   `json:"is_active"`
	CreatedDate   string `json:"created_date"`
	UpdatedDate   string `json:"updated_date"`
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
COALESCE (CreatedDate,'') AS CreatedDate,
COALESCE (UpdatedDate,'') AS UpdatedDate
 FROM SureSureBank`
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
COALESCE (CreatedDate,'') AS CreatedDate,
COALESCE (UpdatedDate,'') AS UpdatedDate
 FROM SureSureBank WHERE UserID = @ID`
var SQL_BANK_ACCOUNT_DELETE = "DELETE FROM SureSureBank WHERE ID = @ID"
