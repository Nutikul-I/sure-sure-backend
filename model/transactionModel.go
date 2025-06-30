package model

type SureSureTransaction struct {
	ID               int64   `json:"id" db:"id"`
	UserID           int64   `json:"user_id" db:"userid"`
	QrCode           string  `json:"qr_code" db:"qrcode"`
	RefNo            string  `json:"ref_no" db:"refno"`
	LineUserID       string  `json:"line_user_id" db:"lineuserid"`
	LineGroupID      string  `json:"line_group_id" db:"linegroupid"`
	Amount           float64 `json:"amount" db:"amount"`
	CSTID            string  `json:"cstid" db:"cstid"`
	RQUID            string  `json:"rquid" db:"rquid"`
	TXID             string  `json:"txid" db:"txid"`
	SenderBankCode   string  `json:"sender_bank_code" db:"senderbankcode"`
	SenderAccountNo  string  `json:"sender_account_no" db:"senderaccountno"`
	SenderName       string  `json:"sender_name" db:"sendername"`
	SenderName2      string  `json:"sender_name2" db:"sendername2"`
	ReceiveBankCode  string  `json:"receive_bank_code" db:"receivebankcode"`
	ReceiveAccountNo string  `json:"receive_account_no" db:"receiveaccountno"`
	ProxyAccountNo   string  `json:"proxy_account_no" db:"proxyaccountno"`
	Ref1             string  `json:"ref1" db:"ref1"`
	Ref2             string  `json:"ref2" db:"ref2"`
	ReceiveName      string  `json:"receive_name" db:"receivename"`
	ReceiveName2     string  `json:"receive_name2" db:"receivename2"`
	Message          string  `json:"message" db:"message"`
	StatusCode       string  `json:"status_code" db:"statuscode"`
	Status           string  `json:"status" db:"status"`
	TransDate        string  `json:"trans_date" db:"transdate"`
	TransTime        string  `json:"trans_time" db:"transtime"`
	CreatedDate      string  `json:"created_date" db:"createddate"`
	UpdatedDate      string  `json:"updated_date" db:"updateddate"`
}

var SQL_GET_ALL_TRANSACTIONS = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (QrCode,'') AS QrCode,
COALESCE (RefNo,'') AS RefNo,
COALESCE (LineUserID,'') AS LineUserID,
COALESCE (LineGroupID,'') AS LineGroupID,
COALESCE (Amount,0.00) AS Amount,
COALESCE (CSTID,'') AS CSTID,
COALESCE (RQUID,'') AS RQUID,
COALESCE (TXID,'') AS TXID,
COALESCE (SenderBankCode,'') AS SenderBankCode,
COALESCE (SenderAccountNo,'') AS SenderAccountNo,
COALESCE (SenderName,'') AS SenderName,
COALESCE (SenderName2,'') AS SenderName2,
COALESCE (ReceiveBankCode,'') AS ReceiveBankCode,
COALESCE (ReceiveAccountNo,'') AS ReceiveAccountNo,
COALESCE (ProxyAccountNo, '') AS ProxyAccountNo,
COALESCE (Ref1, '') AS Ref1,
COALESCE (Ref2, '') AS Ref2,
COALESCE (ReceiveName,'') AS ReceiveName,
COALESCE (ReceiveName2,'') AS ReceiveName2,
COALESCE (Message,'') AS Message,
COALESCE (StatusCode,'') AS StatusCode,
COALESCE (Status,'') AS Status,
COALESCE (TransDate,'') AS TransDate,
COALESCE (TransTime,'') AS TransTime,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate FROM SureSureTransaction`
var SQL_GET_TRANSACTION_BY_ID = `SELECT 
ID,
COALESCE (UserID,0) AS UserID,
COALESCE (QrCode,'') AS QrCode,
COALESCE (RefNo,'') AS RefNo,
COALESCE (LineUserID,'') AS LineUserID,
COALESCE (LineGroupID,'') AS LineGroupID,
COALESCE (Amount,0.00) AS Amount,
COALESCE (CSTID,'') AS CSTID,
COALESCE (RQUID,'') AS RQUID,
COALESCE (TXID,'') AS TXID,
COALESCE (SenderBankCode,'') AS SenderBankCode,
COALESCE (SenderAccountNo,'') AS SenderAccountNo,
COALESCE (SenderName,'') AS SenderName,
COALESCE (SenderName2,'') AS SenderName2,
COALESCE (ReceiveBankCode,'') AS ReceiveBankCode,
COALESCE (ReceiveAccountNo,'') AS ReceiveAccountNo,
COALESCE (ProxyAccountNo, '') AS ProxyAccountNo,
COALESCE (Ref1, '') AS Ref1,
COALESCE (Ref2, '') AS Ref2,
COALESCE (ReceiveName,'') AS ReceiveName,
COALESCE (ReceiveName2,'') AS ReceiveName2,
COALESCE (Message,'') AS Message,
COALESCE (StatusCode,'') AS StatusCode,
COALESCE (Status,'') AS Status,
COALESCE (TransDate,'') AS TransDate,
COALESCE (TransTime,'') AS TransTime,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
 FROM SureSureTransaction WHERE UserID = $1`
var SQL_DELETE_TRANSACTION = "DELETE FROM SureSureTransaction WHERE ID = $1"
