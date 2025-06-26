package model

type SureSurePackage struct {
	ID           int     `json:"id" db:"ID"`
	PackageName  string  `json:"package_name" db:"PackageName"`
	PackagePrice float64 `json:"package_price" db:"PackagePrice"`
	QuotaLimit   int     `json:"quota_limit" db:"QuotaLimit"`
	Amount       float64 `json:"amount" db:"Amount"`
	Ordered      int     `json:"ordered" db:"Ordered"`
	Duration     int     `json:"duration" db:"Duration"`
	IsActive     int     `json:"is_active" db:"IsActive"`
	CreatedDate  string  `json:"created_date" db:"CreatedDate"`
	UpdatedDate  string  `json:"updated_date" db:"UpdatedDate"`
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
    SELECT PackageID
    FROM SureSureOrderPackage
    WHERE RefNo = $1
    LIMIT 1
);`

var SQL_USER_FROM_ORDER_UPDATE = `
UPDATE SureSureUser
SET
    PackageID = (
        SELECT PackageID
        FROM SureSureOrderPackage
        WHERE RefNo = $1
        LIMIT 1
    ),
    QuotaUsage = 0,
    QuotaLeft = (
            SELECT COALESCE(QuotaLimit, 0)
            FROM SureSurePackage
            WHERE ID = (
                SELECT PackageID
                FROM SureSureOrderPackage
                WHERE RefNo = $1
                LIMIT 1
            )
        ),
    QuotaAll = (
            SELECT COALESCE(QuotaLimit, 0)
            FROM SureSurePackage
            WHERE ID = (
                SELECT PackageID
                FROM SureSureOrderPackage
                WHERE RefNo = $1
                LIMIT 1
            )
        ),
    PackageChangeDate = CURRENT_DATE + INTERVAL '1 day' * (
        SELECT COALESCE(Duration, 0)
        FROM SureSurePackage
        WHERE ID = (
            SELECT PackageID
            FROM SureSureOrderPackage
            WHERE RefNo = $1
            LIMIT 1
        )
    )
WHERE ID = (
    SELECT UserID
    FROM SureSureOrderPackage
    WHERE RefNo = $1
    LIMIT 1
);`
