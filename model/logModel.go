package model

type SureSureLog struct {
	ID          int    `json:"id" db:"id"`
	Action      string `json:"action" db:"action"`
	MethodName  string `json:"method_name" db:"methodname"`
	UserID      int    `json:"user_id" db:"userid"`
	NameTH      string `json:"name_th" db:"nameth"`
	DataRequest string `json:"data_request" db:"datarequest"`
	CreatedDate string `json:"created_date" db:"createddate"`
	UpdatedDate string `json:"updated_date" db:"updateddate"`
}

var SQL_LOG_GET = `SELECT 
ID,
COALESCE (Action,'') AS Action,
COALESCE (MethodName,'') AS MethodName,
COALESCE (UserID,'') AS UserID,
COALESCE (NameTH,'') AS NameTH,
COALESCE (DataRequest,'') AS DataRequest,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
FROM SureSureLog`
var SQL_LOG_GET_BYID = `SELECT 
ID,
COALESCE (Action,'') AS Action,
COALESCE (MethodName,'') AS MethodName,
COALESCE (UserID,'') AS UserID,
COALESCE (NameTH,'') AS NameTH,
COALESCE (DataRequest,'') AS DataRequest,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate FROM SureSureLog WHERE UserID = $1`
var SQL_LOG_DELETE = "DELETE FROM SureSureLog WHERE ID = $1"
