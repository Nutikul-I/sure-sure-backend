package model

type SureSureOrderPackage struct {
	ID          int     `json:"id" db:"id"`
	RefNo       string  `json:"ref_no" db:"refno"`
	UserID      int     `json:"user_id" db:"userid"`
	PackageID   int     `json:"package_id" db:"packageid"`
	Price       float64 `json:"price" db:"price"`
	Status      string  `json:"status" db:"status"`
	CreatedDate string  `json:"created_date" db:"createddate"`
	UpdatedDate string  `json:"updated_date" db:"updateddate"`
}
type OrderDetailPost struct {
	ReferenceNo             string `json:"ReferenceNo" `
	OrderNo                 string `json:"OrderNo" `
	MerchantID              int    `json:"MerchantID" `
	ProductDetail           string `json:"ProductDetail" `
	Total                   int    `json:"Total" `
	FeeRate                 int    `json:"FeeRate" `
	CurrencyRate            int    `json:"CurrencyRate" `
	CardType                string `json:"CardType" `
	CardIssuer              string `json:"CardIssuer" `
	CurrencyCode            string `json:"CurrencyCode" `
	CustomerEmail           string `json:"CustomerEmail" `
	Status                  string `json:"Status" `
	StatusName              string `json:"StatusName" `
	PostBackUrl             string `json:"PostBackUrl "`
	PostBackParameters      string `json:"PostBackParameters" `
	PostBackMethod          string `json:"PostBackMethod" `
	PostBackCompleted       bool   `json:"PostBackCompleted" `
	ReturnUrl               string `json:"ReturnUrl" `
	Installment             string `json:"installment" `
	InstallmentMode         string `json:"InstallmentMode" `
	InstallmentInterestRate string `json:"InstallmentInterestRate" `
	InstallmentMonth        string `json:"InstallmentMonth" `
	TotalInterestRate       int    `json:"TotalInterestRate" `
	PaymentPerMonth         int    `json:"PaymentPerMonth" `
	OrderDateTime           string `json:"OrderDateTime" `
}

var SQL_ORDER_PACKAGE_GET = `SELECT 
ID,
COALESCE (RefNo,'') AS RefNo,
COALESCE (UserID,0) AS UserID,
COALESCE (PackageID,0) AS PackageID,
COALESCE (Price,0.00) AS Price,
COALESCE (Status,'') AS Status,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
FROM SureSureOrderPackage`
var SQL_ORDER_PACKAGE_PENDING_GET = `SELECT 
ID,
COALESCE (RefNo,'') AS RefNo,
COALESCE (UserID,0) AS UserID,
COALESCE (PackageID,0) AS PackageID,
COALESCE (Price,0.00) AS Price,
COALESCE (Status,'') AS Status,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
FROM SureSureOrderPackage
WHERE Status = 'PENDING'`
var SQL_ORDER_PACKAGE_GET_BYID = `SELECT 
ID,
COALESCE (RefNo,'') AS RefNo,
COALESCE (UserID,0) AS UserID,
COALESCE (PackageID,0) AS PackageID,
COALESCE (Price,0.00) AS Price,
COALESCE (Status,'') AS Status,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
FROM SureSureOrderPackage WHERE UserID = $1`
var SQL_ORDER_PACKAGE_GET_BYREFNO = `SELECT 
ID,
COALESCE (RefNo,'') AS RefNo,
COALESCE (UserID,0) AS UserID,
COALESCE (PackageID,0) AS PackageID,
COALESCE (Price,0.00) AS Price,
COALESCE (Status,'') AS Status,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
FROM SureSureOrderPackage WHERE RefNo = $1`
var SQL_ORDER_PACKAGE_DELETE = "DELETE FROM SureSureOrderPackage WHERE ID = $1"
