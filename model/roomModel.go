package model

type SureSureRoom struct {
	ID             int     `json:"id"`
	UserID         int     `json:"user_id"`
	LineGroupID    string  `json:"line_group_id"`
	RoomName       string  `json:"room_name"`
	QRToken        string  `json:"qr_token"`
	QuotaUsed      int     `json:"quota_used"`
	MinRecieve     float64 `json:"min_receive"`
	ShowTransferor bool    `json:"show_transferor"`
	ShowRecipient  bool    `json:"show_recipient"`
	ListBank       string  `json:"list_bank"`
	CreatedDate    string  `json:"created_date"`
	UpdatedDate    string  `json:"updated_date"`
}

var SQL_ROOM_GET = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (LineGroupID,'') AS LineGroupID,
COALESCE (RoomName,'') AS RoomName,
COALESCE (QRToken,'') AS QRToken,
COALESCE (QuotaUsed,0) AS QuotaUsed,
COALESCE (MinRecieve,0.00) AS MinRecieve,
COALESCE (ShowTransferor,0) AS ShowTransferor,
COALESCE (ShowRecipient,0) AS ShowRecipient,
COALESCE (ListBank,'') AS ListBank,
COALESCE (CreatedDate,'') AS CreatedDate,
COALESCE (UpdatedDate,'') AS UpdatedDate
 FROM SureSureRoom`
var SQL_ROOM_GET_BYUSERID = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (LineGroupID,'') AS LineGroupID,
COALESCE (RoomName,'') AS RoomName,
COALESCE (QRToken,'') AS QRToken,
COALESCE (QuotaUsed,0) AS QuotaUsed,
COALESCE (MinRecieve,0.00) AS MinRecieve,
COALESCE (ShowTransferor,0) AS ShowTransferor,
COALESCE (ShowRecipient,0) AS ShowRecipient,
COALESCE (ListBank,'') AS ListBank,
COALESCE (CreatedDate,'') AS CreatedDate,
COALESCE (UpdatedDate,'') AS UpdatedDate 
FROM SureSureRoom WHERE UserID = @ID`
var SQL_ROOM_GET_BYID = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (LineGroupID,'') AS LineGroupID,
COALESCE (RoomName,'') AS RoomName,
COALESCE (QRToken,'') AS QRToken,
COALESCE (QuotaUsed,0) AS QuotaUsed,
COALESCE (MinRecieve,0.00) AS MinRecieve,
COALESCE (ShowTransferor,0) AS ShowTransferor,
COALESCE (ShowRecipient,0) AS ShowRecipient,
COALESCE (ListBank,'') AS ListBank,
COALESCE (CreatedDate,'') AS CreatedDate,
COALESCE (UpdatedDate,'') AS UpdatedDate 
FROM SureSureRoom WHERE ID = @ID`
var SQL_ROOM_DELETE = "DELETE FROM SureSureRoom WHERE ID = @ID"
