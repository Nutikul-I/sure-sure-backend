package model

import "time"

type SureSureTransaction struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	QrCode           string    `json:"qr_code"`
	RefNo            string    `json:"ref_no"`
	LineUserID       string    `json:"line_user_id"`
	LineGroupID      string    `json:"line_group_id"`
	Amount           float64   `json:"amount"`
	CSTID            string    `json:"cstid"`
	RQUID            string    `json:"rquid"`
	TXID             string    `json:"txid"`
	SenderBankCode   string    `json:"sender_bank_code"`
	SenderAccountNo  string    `json:"sender_account_no"`
	SenderName       string    `json:"sender_name"`
	SenderName2      string    `json:"sender_name2"`
	ReceiveBankCode  string    `json:"receive_bank_code"`
	ReceiveAccountNo string    `json:"receive_account_no"`
	ProxyAccountNo   string    `json:"proxy_account_no"`
	Ref1             string    `json:"ref1"`
	Ref2             string    `json:"ref2"`
	ReceiveName      string    `json:"receive_name"`
	ReceiveName2     string    `json:"receive_name2"`
	Message          string    `json:"message"`
	StatusCode       string    `json:"status_code"`
	Status           string    `json:"status"`
	TransDate        string    `json:"trans_date"`
	TransTime        string    `json:"trans_time"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedDate      time.Time `json:"updated_date"`
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
COALESCE (CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS CreatedDate,
COALESCE (UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS UpdatedDate
 FROM SureSureTransaction`
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
COALESCE (CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS CreatedDate,
COALESCE (UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS UpdatedDate
 FROM SureSureTransaction WHERE UserID = @ID`
var SQL_DELETE_TRANSACTION = "DELETE FROM SureSureTransaction WHERE ID = $1"
