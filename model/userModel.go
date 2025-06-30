package model

type SureSureUser struct {
	ID                int    `json:"id" db:"id"`
	UID               string `json:"uid" db:"uid"`
	MerchantID        int    `json:"merchant_id" db:"merchantid"`
	PackageID         int    `json:"package_id" db:"packageid"`
	Token             string `json:"token" db:"token"`
	AccessToken       string `json:"access_token" db:"accesstoken"`
	UserType          string `json:"user_type" db:"usertype"`
	Picture           string `json:"picture" db:"picture"`
	NameTH            string `json:"name_th" db:"nameth"`
	NameEN            string `json:"name_en" db:"nameen"`
	Phone             string `json:"phone" db:"phone"`
	Website           string `json:"website" db:"website"`
	UserRole          string `json:"user_role" db:"userrole"`
	Address           string `json:"address" db:"address"`
	Email             string `json:"email" db:"email"`
	Username          string `json:"username" db:"username"`
	Password          string `json:"password" db:"password"`
	IsActive          bool   `json:"is_active" db:"isactive"`
	StoreName         string `json:"store_name" db:"storename"`
	StoreCategoryType string `json:"store_category_type" db:"storecategorytype"`
	StorePhone        string `json:"store_phone" db:"storephone"`
	StoreEmail        string `json:"store_email" db:"storeemail"`
	QuotaUsage        int    `json:"quota_usage" db:"quotausage"`
	QuotaLeft         int    `json:"quota_left" db:"quotaleft"`
	QuotaALL          int    `json:"quota_all" db:"quotaall"`
	Step              int    `json:"step" db:"step"`
	PackageChangeDate string `json:"package_change_date" db:"packagechangedate"`
	BillDate          string `json:"bill_date" db:"billdate"`
	CreatedDate       string `json:"created_date" db:"createddate"`
	UpdatedDate       string `json:"updated_date" db:"updateddate"`
}

type MerchantCategory struct {
	CatID            int    `json:"cat_id" db:"catid"`
	ISOCode          string `json:"iso_code" db:"isocode"`
	CategoryNameEN   string `json:"category_name_en" db:"categorynameen"`
	CategoryNameTH   string `json:"category_name_th" db:"categorynameth"`
	CategoryDetailEN string `json:"category_detail_en" db:"categorydetailen"`
	CategoryDetailTH string `json:"category_detail_th" db:"categorydetailth"`
	Logo             string `json:"logo" db:"logo"`
	Priority         int    `json:"priority" db:"priority"`
	Enable           bool   `json:"enable" db:"enable"`
}

var SQL_USER_GET = `SELECT 
id, 
uid, 
COALESCE (merchantid, 0) AS merchantid, 
COALESCE (packageid, 0) AS packageid, 
COALESCE (token, '') AS token, 
COALESCE (accesstoken, '') AS accesstoken, 
COALESCE (usertype, '') AS usertype, 
COALESCE (picture, '') AS picture, 
COALESCE (nameth, '') AS nameth, 
COALESCE (nameen, '') AS nameen, 
COALESCE (phone, '') AS phone, 
COALESCE (website, '') AS website, 
COALESCE (userrole, '') AS userrole, 
COALESCE (address, '') AS address, 
COALESCE (email, '') AS email, 
COALESCE (username, '') AS username, 
COALESCE (isactive, 0) AS isactive, 
COALESCE (storename, '') AS storename, 
COALESCE (storecategorytype, '') AS storecategorytype, 
COALESCE (storephone, '') AS storephone, 
COALESCE (storeemail, '') AS storeemail, 
COALESCE (quotausage, 0) AS quotausage, 
COALESCE (quotaleft, 0) AS quotaleft, 
COALESCE (quotaall, 0) AS quotaall,
COALESCE (step, 0) AS step,
packagechangedate AS packagechangedate, 
billdate AS billdate, 
createddate AS createddate, 
updateddate AS updateddate FROM suresureuser`

var SQL_USER_GET_BYID = `SELECT 
id, 
uid, 
COALESCE (merchantid, 0) AS merchantid, 
COALESCE (packageid, 0) AS packageid, 
COALESCE (token, '') AS token, 
COALESCE (accesstoken, '') AS accesstoken, 
COALESCE (usertype, '') AS usertype, 
COALESCE (picture, '') AS picture, 
COALESCE (nameth, '') AS nameth, 
COALESCE (nameen, '') AS nameen, 
COALESCE (phone, '') AS phone, 
COALESCE (website, '') AS website, 
COALESCE (userrole, '') AS userrole, 
COALESCE (address, '') AS address, 
COALESCE (email, '') AS email, 
COALESCE (username, '') AS username, 
COALESCE (isactive, 0) AS isactive, 
COALESCE (storename, '') AS storename, 
COALESCE (storecategorytype, '') AS storecategorytype, 
COALESCE (storephone, '') AS storephone, 
COALESCE (storeemail, '') AS storeemail, 
COALESCE (quotausage, 0) AS quotausage, 
COALESCE (quotaleft, 0) AS quotaleft, 
COALESCE (quotaall, 0) AS quotaall,
COALESCE (step, 0) AS step,
packagechangedate AS packagechangedate, 
billdate AS billdate, 
createddate AS createddate, 
updateddate AS updateddate FROM suresureuser WHERE id = $1`

var SQL_USER_GET_BY_USERNAME = `SELECT 
id, 
uid, 
COALESCE (merchantid, 0) AS merchantid, 
COALESCE (packageid, 0) AS packageid, 
COALESCE (token, '') AS token, 
COALESCE (accesstoken, '') AS accesstoken, 
COALESCE (usertype, '') AS usertype, 
COALESCE (picture, '') AS picture, 
COALESCE (nameth, '') AS nameth, 
COALESCE (nameen, '') AS nameen, 
COALESCE (phone, '') AS phone, 
COALESCE (website, '') AS website, 
COALESCE (userrole, '') AS userrole, 
COALESCE (address, '') AS address, 
COALESCE (email, '') AS email, 
COALESCE (username, '') AS username, 
COALESCE (isactive, 0) AS isactive, 
COALESCE (storename, '') AS storename, 
COALESCE (storecategorytype, '') AS storecategorytype, 
COALESCE (storephone, '') AS storephone, 
COALESCE (storeemail, '') AS storeemail, 
COALESCE (quotausage, 0) AS quotausage, 
COALESCE (quotaleft, 0) AS quotaleft, 
COALESCE (quotaall, 0) AS quotaall,
COALESCE (step, 0) AS step,
packagechangedate AS packagechangedate, 
billdate AS billdate, 
createddate AS createddate, 
updateddate AS updateddate FROM suresureuser WHERE username = $1 AND password = $2`

var SQL_USER_DELETE = "DELETE FROM suresureuser WHERE uid = $1"

var SQL_USER_GET_DUPLICATE = "SELECT COUNT(storename) FROM suresureuser ssu WHERE storename = $1 AND storephone = $2 AND storeemail = $3 AND id != $4"

var SQL_CATEGORY_GET = `SELECT 
"catid" AS CatID, 
COALESCE ("isocode", '') AS ISOCode, 
COALESCE ("categorynameen", '') AS CategoryNameEN, 
COALESCE ("categorynameth", '') AS CategoryNameTH, 
COALESCE ("categorydetailen", '') AS CategoryDetailEN, 
COALESCE ("categorydetailth", '') AS CategoryDetailTH, 
COALESCE ("logo", '') AS Logo, 
"priority" AS Priority, 
"enable" AS Enable 
FROM "merchantcategory";`
