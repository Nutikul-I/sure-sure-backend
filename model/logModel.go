package model

type SureSureLog struct {
	ID          int    `json:"id"`
	Action      string `json:"action"`
	MethodName  string `json:"method_name"`
	UserID      int    `json:"user_id"`
	NameTH      string `json:"name_th"`
	DataRequest string `json:"data_request"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}

var SQL_LOG_GET = `SELECT 
ID,
ISNULL (Action,'') AS Action,
ISNULL (MethodName,'') AS MethodName,
ISNULL (UserID,'') AS UserID,
ISNULL (NameTH,'') AS NameTH,
ISNULL (DataRequest,'') AS DataRequest,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
FROM SureSureLog`
var SQL_LOG_GET_BYID = `SELECT 
ID,
ISNULL (Action,'') AS Action,
ISNULL (MethodName,'') AS MethodName,
ISNULL (UserID,'') AS UserID,
ISNULL (NameTH,'') AS NameTH,
ISNULL (DataRequest,'') AS DataRequest,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate FROM SureSureLog WHERE UserID = @ID`
var SQL_LOG_DELETE = "DELETE FROM SureSureLog WHERE ID = @ID"
