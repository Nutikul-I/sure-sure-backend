package model

type SureSureRoom struct {
	ID             int     `json:"id" db:"id"`
	UserID         int     `json:"user_id" db:"userid"`
	LineGroupID    string  `json:"line_group_id" db:"linegroupid"`
	RoomName       string  `json:"room_name" db:"roomname"`
	QRToken        string  `json:"qr_token" db:"qrtoken"`
	QuotaUsed      int     `json:"quota_used" db:"quotaused"`
	MinRecieve     float64 `json:"min_receive" db:"minrecieve"`
	ShowTransferor bool    `json:"show_transferor" db:"showtransferor"`
	ShowRecipient  bool    `json:"show_recipient" db:"showrecipient"`
	ListBank       string  `json:"list_bank" db:"listbank"`
	CreatedDate    string  `json:"created_date" db:"createddate"`
	UpdatedDate    string  `json:"updated_date" db:"updateddate"`
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
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
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
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate 
FROM SureSureRoom WHERE UserID = $1`
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
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate 
FROM SureSureRoom WHERE ID = $1`
var SQL_ROOM_DELETE = "DELETE FROM SureSureRoom WHERE ID = $1"
