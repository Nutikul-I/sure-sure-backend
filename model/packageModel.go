package model

type SureSurePackage struct {
	ID           int     `json:"id"`
	PackageName  string  `json:"package_name"`
	PackagePrice float64 `json:"package_price"`
	QuotaLimit   int     `json:"quota_limit"`
	Amount       float64 `json:"amount"`
	Ordered      int     `json:"ordered"`
	Duration     int     `json:"duration"`
	IsActive     bool    `json:"is_active"`
	CreatedDate  string  `json:"created_date"`
	UpdatedDate  string  `json:"updated_date"`
}

var SQL_PACKAGE_GET = `SELECT 
ID,
COALESCE (PackageName,'') AS PackageName,
COALESCE (PackagePrice,0.00) AS PackagePrice,
COALESCE (QuotaLimit,0) AS QuotaLimit,
COALESCE (Amount,0.00) AS Amount,
COALESCE (Ordered,0) AS Ordered,
COALESCE (Duration,0) AS Duration,
COALESCE (IsActive,0) AS IsActive,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
 FROM SureSurePackage`
var SQL_PACKAGE_GET_BYID = `SELECT 
ID,
COALESCE (PackageName,'') AS PackageName,
COALESCE (PackagePrice,0.00) AS PackagePrice,
COALESCE (QuotaLimit,0) AS QuotaLimit,
COALESCE (Amount,0.00) AS Amount,
COALESCE (Ordered,0) AS Ordered,
COALESCE (Duration,0) AS Duration,
COALESCE (IsActive,0) AS IsActive,
COALESCE (TO_CHAR(CreatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS CreatedDate,
COALESCE (TO_CHAR(UpdatedDate, 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),'') AS UpdatedDate
 FROM SureSurePackage WHERE ID = $1`
var SQL_PACKAGE_DELETE = "DELETE FROM SureSurePackage WHERE ID = $1"
var SQL_PACKAGE_FROM_ORDER_UPDATE = `
UPDATE SureSurePackage
SET
    Amount = Amount - 1,
    Ordered = Ordered + 1
WHERE ID = (
    SELECT TOP 1 PackageID
    FROM SureSureOrderPackage
    WHERE RefNo = @RefNo
);`

var SQL_USER_FROM_ORDER_UPDATE = `

UPDATE SureSureUser
SET
    PackageID = (
        SELECT TOP 1 PackageID
        FROM SureSureOrderPackage
        WHERE RefNo = @RefNo
    ),
    QuotaUsage = 0,
    QuotaLeft = (
            SELECT COALESCE(QuotaLimit, 0)
            FROM SureSurePackage
            WHERE ID = (
                SELECT TOP 1 PackageID
                FROM SureSureOrderPackage
                WHERE RefNo = @RefNo
            )
        ),
    QuotaAll = (
            SELECT COALESCE(QuotaLimit, 0)
            FROM SureSurePackage
            WHERE ID = (
                SELECT TOP 1 PackageID
                FROM SureSureOrderPackage
                WHERE RefNo = @RefNo
            )
        ),
    PackageChangeDate = DATEADD(DAY, (
        SELECT COALESCE(Duration, 0)
        FROM SureSurePackage
        WHERE ID = (
            SELECT TOP 1 PackageID
            FROM SureSureOrderPackage
            WHERE RefNo = @RefNo
        )
    ), GETDATE())
WHERE ID = (
    SELECT TOP 1 UserID
    FROM SureSureOrderPackage
    WHERE RefNo = @RefNo
);
`
