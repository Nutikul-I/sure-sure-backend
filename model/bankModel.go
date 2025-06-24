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
ISNULL (UserID,0) AS UserID,
ISNULL (BankCode,'') AS BankCode,
ISNULL (PromptPayType,'') AS PromptPayType,
ISNULL (AccountNo,'') AS AccountNo,
ISNULL (AccountType,'') AS AccountType,
ISNULL (NameTH,'') AS NameTH,
ISNULL (NameEN,'') AS NameEN,
ISNULL (IsActive,0) AS IsActive,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
 FROM SureSureBank`
var SQL_BANK_ACCOUNT_GET_BYID = `SELECT 
ID,
ISNULL (UserID,0) AS UserID,
ISNULL (BankCode,'') AS BankCode,
ISNULL (PromptPayType,'') AS PromptPayType,
ISNULL (AccountNo,'') AS AccountNo,
ISNULL (AccountType,'') AS AccountType,
ISNULL (NameTH,'') AS NameTH,
ISNULL (NameEN,'') AS NameEN,
ISNULL (IsActive,0) AS IsActive,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
 FROM SureSureBank WHERE UserID = @ID`
var SQL_BANK_ACCOUNT_DELETE = "DELETE FROM SureSureBank WHERE ID = @ID"
