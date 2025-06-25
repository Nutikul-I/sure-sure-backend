package model

type SureSureUser struct {
	ID                int    `json:"id"`
	UID               string `json:"uid"`
	MerchantID        int    `json:"merchant_id"`
	PackageID         int    `json:"package_id"`
	Token             string `json:"token"`
	AccessToken       string `json:"access_token"`
	UserType          string `json:"user_type"`
	Picture           string `json:"picture"`
	NameTH            string `json:"name_th"`
	NameEN            string `json:"name_en"`
	Phone             string `json:"phone"`
	Website           string `json:"website"`
	UserRole          string `json:"user_role"`
	Address           string `json:"address"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	IsActive          bool   `json:"is_active"`
	StoreName         string `json:"store_name"`
	StoreCategoryType string `json:"store_category_type"`
	StorePhone        string `json:"store_phone"`
	StoreEmail        string `json:"store_email"`
	QuotaUsage        int    `json:"quota_usage"`
	QuotaLeft         int    `json:"quota_left"`
	QuotaALL          int    `json:"quota_all"`
	Step              int    `json:"step"`
	PackageChangeDate string `json:"package_change_date"`
	BillDate          string `json:"bill_date"`
	CreatedDate       string `json:"created_date"`
	UpdatedDate       string `json:"updated_date"`
}

type MerchantCategory struct {
	CatID            int    `json:"cat_id"`
	ISOCode          string `json:"iso_code"`
	CategoryNameEN   string `json:"category_name_en"`
	CategoryNameTH   string `json:"category_name_th"`
	CategoryDetailEN string `json:"category_detail_en"`
	CategoryDetailTH string `json:"category_detail_th"`
	Logo             string `json:"logo"`
	Priority         int    `json:"priority"`
	Enable           bool   `json:"enable"`
}

var SQL_USER_GET = `SELECT 
ID, 
UID, 
COALESCE(MerchantID, 0) AS MerchantID, 
COALESCE(PackageID, 0) AS PackageID, 
COALESCE(Token, '') AS Token, 
COALESCE(AccessToken, '') AS AccessToken, 
COALESCE(UserType, '') AS UserType, 
COALESCE(Picture, '') AS Picture, 
COALESCE(NameTH, '') AS NameTH, 
COALESCE(NameEN, '') AS NameEN, 
COALESCE(Phone, '') AS Phone, 
COALESCE(Website, '') AS Website, 
COALESCE(UserRole, '') AS UserRole, 
COALESCE(Address, '') AS Address, 
COALESCE(Email, '') AS Email, 
COALESCE(Username, '') AS Username, 
COALESCE(IsActive, 0) AS IsActive, 
COALESCE(StoreName, '') AS StoreName, 
COALESCE(StoreCategoryType, '') AS StoreCategoryType, 
COALESCE(StorePhone, '') AS StorePhone, 
COALESCE(StoreEmail, '') AS StoreEmail, 
COALESCE(QuotaUsage, 0) AS QuotaUsage, 
COALESCE(QuotaLeft, 0) AS QuotaLeft, 
COALESCE(QuotaALL, 0) AS QuotaALL,
COALESCE(Step, 0) AS Step,
PackageChangeDate AS PackageChangeDate, 
BillDate AS BillDate, 
CreatedDate AS CreatedDate, 
UpdatedDate AS UpdatedDate FROM SureSureUser`
var SQL_USER_GET_BYID = `SELECT 
ID, 
UID, 
COALESCE(MerchantID, 0) AS MerchantID, 
COALESCE(PackageID, 0) AS PackageID, 
COALESCE(Token, '') AS Token, 
COALESCE(AccessToken, '') AS AccessToken, 
COALESCE(UserType, '') AS UserType, 
COALESCE(Picture, '') AS Picture, 
COALESCE(NameTH, '') AS NameTH, 
COALESCE(NameEN, '') AS NameEN, 
COALESCE(Phone, '') AS Phone, 
COALESCE(Website, '') AS Website, 
COALESCE(UserRole, '') AS UserRole, 
COALESCE(Address, '') AS Address, 
COALESCE(Email, '') AS Email, 
COALESCE(Username, '') AS Username, 
COALESCE(IsActive, 0) AS IsActive, 
COALESCE(StoreName, '') AS StoreName, 
COALESCE(StoreCategoryType, '') AS StoreCategoryType, 
COALESCE(StorePhone, '') AS StorePhone, 
COALESCE(StoreEmail, '') AS StoreEmail, 
COALESCE(QuotaUsage, 0) AS QuotaUsage, 
COALESCE(QuotaLeft, 0) AS QuotaLeft, 
COALESCE(QuotaALL, 0) AS QuotaALL,
COALESCE(Step, 0) AS Step,
PackageChangeDate AS PackageChangeDate, 
BillDate AS BillDate, 
CreatedDate AS CreatedDate, 
UpdatedDate AS UpdatedDate FROM SureSureUser WHERE UID = @ID`
var SQL_USER_GET_BY_USERNAME = `SELECT 
ID, 
UID, 
COALESCE(MerchantID, 0) AS MerchantID, 
COALESCE(PackageID, 0) AS PackageID, 
COALESCE(Token, '') AS Token, 
COALESCE(AccessToken, '') AS AccessToken, 
COALESCE(UserType, '') AS UserType, 
COALESCE(Picture, '') AS Picture, 
COALESCE(NameTH, '') AS NameTH, 
COALESCE(NameEN, '') AS NameEN, 
COALESCE(Phone, '') AS Phone, 
COALESCE(Website, '') AS Website, 
COALESCE(UserRole, '') AS UserRole, 
COALESCE(Address, '') AS Address, 
COALESCE(Email, '') AS Email, 
COALESCE(Username, '') AS Username, 
COALESCE(IsActive, 0) AS IsActive, 
COALESCE(StoreName, '') AS StoreName, 
COALESCE(StoreCategoryType, '') AS StoreCategoryType, 
COALESCE(StorePhone, '') AS StorePhone, 
COALESCE(StoreEmail, '') AS StoreEmail, 
COALESCE(QuotaUsage, 0) AS QuotaUsage, 
COALESCE(QuotaLeft, 0) AS QuotaLeft, 
COALESCE(QuotaALL, 0) AS QuotaALL,
COALESCE(Step, 0) AS Step,
PackageChangeDate AS PackageChangeDate, 
BillDate AS BillDate, 
CreatedDate AS CreatedDate, 
UpdatedDate AS UpdatedDate FROM SureSureUser WHERE Username = @Username AND Password = @Password`
var SQL_USER_DELETE = "DELETE FROM SureSureUser WHERE UID = @ID"
var SQL_USER_GET_DUPLICATE = "SELECT COUNT(StoreName) FROM SureSureUser ssu WHERE  StoreName = @StoreName AND StorePhone = @StorePhone AND StoreEmail = @StoreEmail AND ID != @ID"

var SQL_CATEGORY_GET = `SELECT 
CatID, 
COALESCE(ISOCode, '') AS ISOCode, 
COALESCE(CategoryNameEN, '') AS CategoryNameEN, 
COALESCE(CategoryNameTH, '') AS CategoryNameTH, 
COALESCE(CategoryDetailEN, '') AS CategoryDetailEN, 
COALESCE(CategoryDetailTH, '') AS CategoryDetailTH, 
COALESCE(Logo, '') AS Logo, 
Priority, 
Enable 
FROM MerchantCategory;`
