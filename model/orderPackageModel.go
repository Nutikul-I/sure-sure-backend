package model

type SureSureOrderPackage struct {
	ID          int     `json:"id"`
	RefNo       string  `json:"ref_no"`
	UserID      int     `json:"user_id"`
	PackageID   int     `json:"package_id"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
	CreatedDate string  `json:"created_date"`
	UpdatedDate string  `json:"updated_date"`
}
type OrderDetailPost struct {
	ReferenceNo             string `json:"ReferenceNo"`
	OrderNo                 string `json:"OrderNo"`
	MerchantID              int    `json:"MerchantID"`
	ProductDetail           string `json:"ProductDetail"`
	Total                   int    `json:"Total"`
	FeeRate                 int    `json:"FeeRate"`
	CurrencyRate            int    `json:"CurrencyRate"`
	CardType                string `json:"CardType"`
	CardIssuer              string `json:"CardIssuer"`
	CurrencyCode            string `json:"CurrencyCode"`
	CustomerEmail           string `json:"CustomerEmail"`
	Status                  string `json:"Status"`
	StatusName              string `json:"StatusName"`
	PostBackUrl             string `json:"PostBackUrl"`
	PostBackParameters      string `json:"PostBackParameters"`
	PostBackMethod          string `json:"PostBackMethod"`
	PostBackCompleted       bool   `json:"PostBackCompleted"`
	ReturnUrl               string `json:"ReturnUrl"`
	Installment             string `json:"installment"`
	InstallmentMode         string `json:"InstallmentMode"`
	InstallmentInterestRate string `json:"InstallmentInterestRate"`
	InstallmentMonth        string `json:"InstallmentMonth"`
	TotalInterestRate       int    `json:"TotalInterestRate"`
	PaymentPerMonth         int    `json:"PaymentPerMonth"`
	OrderDateTime           string `json:"OrderDateTime"`
}

var SQL_ORDER_PACKAGE_GET = `SELECT 
ID,
ISNULL (RefNo,'') AS RefNo,
ISNULL (UserID,0) AS UserID,
ISNULL (PackageID,0) AS PackageID,
ISNULL (Price,0.00) AS Price,
ISNULL (Status,'') AS Status,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
FROM SureSureOrderPackage`
var SQL_ORDER_PACKAGE_PENDING_GET = `SELECT 
ID,
ISNULL (RefNo,'') AS RefNo,
ISNULL (UserID,0) AS UserID,
ISNULL (PackageID,0) AS PackageID,
ISNULL (Price,0.00) AS Price,
ISNULL (Status,'') AS Status,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
FROM SureSureOrderPackage
WHERE Status = 'PENDING'`
var SQL_ORDER_PACKAGE_GET_BYID = `SELECT 
ID,
ISNULL (RefNo,'') AS RefNo,
ISNULL (UserID,0) AS UserID,
ISNULL (PackageID,0) AS PackageID,
ISNULL (Price,0.00) AS Price,
ISNULL (Status,'') AS Status,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
FROM SureSureOrderPackage WHERE UserID = @ID`
var SQL_ORDER_PACKAGE_GET_BYREFNO = `SELECT 
ID,
ISNULL (RefNo,'') AS RefNo,
ISNULL (UserID,0) AS UserID,
ISNULL (PackageID,0) AS PackageID,
ISNULL (Price,0.00) AS Price,
ISNULL (Status,'') AS Status,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
FROM SureSureOrderPackage WHERE RefNo = @RefNo`
var SQL_ORDER_PACKAGE_DELETE = "DELETE FROM SureSureOrderPackage WHERE ID = @ID"
