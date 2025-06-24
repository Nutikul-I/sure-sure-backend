package model

type SureSureTransaction struct {
	ID               int64   `json:"id"`
	UserID           int64   `json:"user_id"`
	QrCode           string  `json:"qr_code"`
	RefNo            string  `json:"ref_no"`
	LineUserID       string  `json:"line_user_id"`
	LineGroupID      string  `json:"line_group_id"`
	Amount           float64 `json:"amount"`
	CSTID            string  `json:"cstid"`
	RQUID            string  `json:"rquid"`
	TXID             string  `json:"txid"`
	SenderBankCode   string  `json:"sender_bank_code"`
	SenderAccountNo  string  `json:"sender_account_no"`
	SenderName       string  `json:"sender_name"`
	SenderName2      string  `json:"sender_name2"`
	ReceiveBankCode  string  `json:"receive_bank_code"`
	ReceiveAccountNo string  `json:"receive_account_no"`
	ProxyAccountNo   string  `json:"proxy_account_no"`
	Ref1             string  `json:"ref1"`
	Ref2             string  `json:"ref2"`
	ReceiveName      string  `json:"receive_name"`
	ReceiveName2     string  `json:"receive_name2"`
	Message          string  `json:"message"`
	StatusCode       string  `json:"status_code"`
	Status           string  `json:"status"`
	TransDate        string  `json:"trans_date"`
	TransTime        string  `json:"trans_time"`
	CreatedDate      string  `json:"created_date"`
	UpdatedDate      string  `json:"updated_date"`
}

var SQL_GET_ALL_TRANSACTIONS = `SELECT 
ID,
ISNULL (UserID,0) AS UserID,
ISNULL (QrCode,'') AS QrCode,
ISNULL (RefNo,'') AS RefNo,
ISNULL (LineUserID,'') AS LineUserID,
ISNULL (LineGroupID,'') AS LineGroupID,
ISNULL (Amount,0.00) AS Amount,
ISNULL (CSTID,'') AS CSTID,
ISNULL (RQUID,'') AS RQUID,
ISNULL (TXID,'') AS TXID,
ISNULL (SenderBankCode,'') AS SenderBankCode,
ISNULL (SenderAccountNo,'') AS SenderAccountNo,
ISNULL (SenderName,'') AS SenderName,
ISNULL (SenderName2,'') AS SenderName2,
ISNULL (ReceiveBankCode,'') AS ReceiveBankCode,
ISNULL (ReceiveAccountNo,'') AS ReceiveAccountNo,
ISNULL (ProxyAccountNo, '') AS ProxyAccountNo,
ISNULL (Ref1, '') AS Ref1,
ISNULL (Ref2, '') AS Ref2,
ISNULL (ReceiveName,'') AS ReceiveName,
ISNULL (ReceiveName2,'') AS ReceiveName2,
ISNULL (Message,'') AS Message,
ISNULL (StatusCode,'') AS StatusCode,
ISNULL (Status,'') AS Status,
ISNULL (TransDate,'') AS TransDate,
ISNULL (TransTime,'') AS TransTime,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
 FROM SureSureTransaction`
var SQL_GET_TRANSACTION_BY_ID = `SELECT 
ID,
ISNULL (UserID,0) AS UserID,
ISNULL (QrCode,'') AS QrCode,
ISNULL (RefNo,'') AS RefNo,
ISNULL (LineUserID,'') AS LineUserID,
ISNULL (LineGroupID,'') AS LineGroupID,
ISNULL (Amount,0.00) AS Amount,
ISNULL (CSTID,'') AS CSTID,
ISNULL (RQUID,'') AS RQUID,
ISNULL (TXID,'') AS TXID,
ISNULL (SenderBankCode,'') AS SenderBankCode,
ISNULL (SenderAccountNo,'') AS SenderAccountNo,
ISNULL (SenderName,'') AS SenderName,
ISNULL (SenderName2,'') AS SenderName2,
ISNULL (ReceiveBankCode,'') AS ReceiveBankCode,
ISNULL (ReceiveAccountNo,'') AS ReceiveAccountNo,
ISNULL (ProxyAccountNo, '') AS ProxyAccountNo,
ISNULL (Ref1, '') AS Ref1,
ISNULL (Ref2, '') AS Ref2,
ISNULL (ReceiveName,'') AS ReceiveName,
ISNULL (ReceiveName2,'') AS ReceiveName2,
ISNULL (Message,'') AS Message,
ISNULL (StatusCode,'') AS StatusCode,
ISNULL (Status,'') AS Status,
ISNULL (TransDate,'') AS TransDate,
ISNULL (TransTime,'') AS TransTime,
ISNULL (CreatedDate,'') AS CreatedDate,
ISNULL (UpdatedDate,'') AS UpdatedDate
 FROM SureSureTransaction WHERE UserID = @ID`
var SQL_DELETE_TRANSACTION = "DELETE FROM SureSureTransaction WHERE ID = @ID"
