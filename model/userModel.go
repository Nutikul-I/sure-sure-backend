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
ISNULL(MerchantID, 0) AS MerchantID, 
ISNULL(PackageID, 0) AS PackageID, 
ISNULL(Token, '') AS Token, 
ISNULL(AccessToken, '') AS AccessToken, 
ISNULL(UserType, '') AS UserType, 
ISNULL(Picture, '') AS Picture, 
ISNULL(NameTH, '') AS NameTH, 
ISNULL(NameEN, '') AS NameEN, 
ISNULL(Phone, '') AS Phone, 
ISNULL(Website, '') AS Website, 
ISNULL(UserRole, '') AS UserRole, 
ISNULL(Address, '') AS Address, 
ISNULL(Email, '') AS Email, 
ISNULL(Username, '') AS Username, 
ISNULL(IsActive, 0) AS IsActive, 
ISNULL(StoreName, '') AS StoreName, 
ISNULL(StoreCategoryType, '') AS StoreCategoryType, 
ISNULL(StorePhone, '') AS StorePhone, 
ISNULL(StoreEmail, '') AS StoreEmail, 
ISNULL(QuotaUsage, 0) AS QuotaUsage, 
ISNULL(QuotaLeft, 0) AS QuotaLeft, 
ISNULL(QuotaALL, 0) AS QuotaALL,
ISNULL(Step, 0) AS Step,
ISNULL(PackageChangeDate, '') AS PackageChangeDate, 
ISNULL(BillDate, '') AS BillDate, 
ISNULL(CreatedDate, '') AS CreatedDate, 
ISNULL(UpdatedDate, '') AS UpdatedDate FROM SureSureUser`
var SQL_USER_GET_BYID = `SELECT 
ID, 
UID, 
ISNULL(MerchantID, 0) AS MerchantID, 
ISNULL(PackageID, 0) AS PackageID, 
ISNULL(Token, '') AS Token, 
ISNULL(AccessToken, '') AS AccessToken, 
ISNULL(UserType, '') AS UserType, 
ISNULL(Picture, '') AS Picture, 
ISNULL(NameTH, '') AS NameTH, 
ISNULL(NameEN, '') AS NameEN, 
ISNULL(Phone, '') AS Phone, 
ISNULL(Website, '') AS Website, 
ISNULL(UserRole, '') AS UserRole, 
ISNULL(Address, '') AS Address, 
ISNULL(Email, '') AS Email, 
ISNULL(Username, '') AS Username, 
ISNULL(IsActive, 0) AS IsActive, 
ISNULL(StoreName, '') AS StoreName, 
ISNULL(StoreCategoryType, '') AS StoreCategoryType, 
ISNULL(StorePhone, '') AS StorePhone, 
ISNULL(StoreEmail, '') AS StoreEmail, 
ISNULL(QuotaUsage, 0) AS QuotaUsage, 
ISNULL(QuotaLeft, 0) AS QuotaLeft, 
ISNULL(QuotaALL, 0) AS QuotaALL,
ISNULL(Step, 0) AS Step,
ISNULL(PackageChangeDate, '') AS PackageChangeDate, 
ISNULL(BillDate, '') AS BillDate, 
ISNULL(CreatedDate, '') AS CreatedDate, 
ISNULL(UpdatedDate, '') AS UpdatedDate FROM SureSureUser WHERE UID = @ID`
var SQL_USER_GET_BY_USERNAME = `SELECT 
ID, 
UID, 
ISNULL(MerchantID, 0) AS MerchantID, 
ISNULL(PackageID, 0) AS PackageID, 
ISNULL(Token, '') AS Token, 
ISNULL(AccessToken, '') AS AccessToken, 
ISNULL(UserType, '') AS UserType, 
ISNULL(Picture, '') AS Picture, 
ISNULL(NameTH, '') AS NameTH, 
ISNULL(NameEN, '') AS NameEN, 
ISNULL(Phone, '') AS Phone, 
ISNULL(Website, '') AS Website, 
ISNULL(UserRole, '') AS UserRole, 
ISNULL(Address, '') AS Address, 
ISNULL(Email, '') AS Email, 
ISNULL(Username, '') AS Username, 
ISNULL(IsActive, 0) AS IsActive, 
ISNULL(StoreName, '') AS StoreName, 
ISNULL(StoreCategoryType, '') AS StoreCategoryType, 
ISNULL(StorePhone, '') AS StorePhone, 
ISNULL(StoreEmail, '') AS StoreEmail, 
ISNULL(QuotaUsage, 0) AS QuotaUsage, 
ISNULL(QuotaLeft, 0) AS QuotaLeft, 
ISNULL(QuotaALL, 0) AS QuotaALL,
ISNULL(Step, 0) AS Step,
ISNULL(PackageChangeDate, '') AS PackageChangeDate, 
ISNULL(BillDate, '') AS BillDate, 
ISNULL(CreatedDate, '') AS CreatedDate, 
ISNULL(UpdatedDate, '') AS UpdatedDate FROM SureSureUser WHERE Username = @Username AND Password = @Password`
var SQL_USER_DELETE = "DELETE FROM SureSureUser WHERE UID = @ID"
var SQL_USER_GET_DUPLICATE = "SELECT COUNT(StoreName) FROM SureSureUser ssu WHERE  StoreName = @StoreName AND StorePhone = @StorePhone AND StoreEmail = @StoreEmail AND ID != @ID"

var SQL_CATEGORY_GET = `SELECT 
CatID, 
ISNULL(ISOCode, '') AS ISOCode, 
ISNULL(CategoryNameEN, '') AS CategoryNameEN, 
ISNULL(CategoryNameTH, '') AS CategoryNameTH, 
ISNULL(CategoryDetailEN, '') AS CategoryDetailEN, 
ISNULL(CategoryDetailTH, '') AS CategoryDetailTH, 
ISNULL(Logo, '') AS Logo, 
Priority, 
Enable 
FROM MerchantCategory;`
