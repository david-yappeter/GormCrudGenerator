package gormgenerator

//CashFlowLines s
type CashFlowLines struct {
	ID             int     `gorm:"type:int AUTO_INCREMENT;not null"`
	Date           string  `gorm:"type:datetime;not null"`
	Information    string  `gorm:"type:varchar(255);not null"`
	EntryType      int     `gorm:"type:int;not null"`
	Amount         int     `gorm:"type:int;not null"`
	EndingBalance  int     `gorm:"type:int;not null"`
	CreatedAt      *string `gorm:"type:timestamp;not null"`
	UpdatedAt      *string `gorm:"type:timestamp;not null"`
	ReceiptsID     *int    `gorm:"type:int;null;default:NULL"`
	CashAccountID  int     `gorm:"type:int;not null"`
	TransactionsID int     `gorm:"type:int;not null"`
}

//CashAccounts Cash Accounts
type CashAccounts struct {
	ID         int     `gorm:"type:int AUTO_INCREMENT;not null"`
	Name       string  `gorm:"type:varchar(255);not null"`
	Balance    int     `gorm:"type:int;not null"`
	CreatedAt  *string `gorm:"type:timestamp;null"`
	UpdatedAt  *string `gorm:"type:timestamp;null"`
	ProjectsID int     `gorm:"type:int;not null"`
	Status     int     `gorm:"type:int;not null"`
}

//CashCategories Cash Categories
type CashCategories struct {
	ID         int     `gorm:"type:int AUTO_INCREMENT;not null"`
	Name       string  `gorm:"type:varchar(255);not null"`
	Status     int     `gorm:"type:int;not null"`
	Type       int     `gorm:"type:int;not null"`
	CreatedAt  *string `gorm:"type:timestamp;null"`
	UpdatedAt  *string `gorm:"type:timestamp;null"`
	ProjectsID int     `gorm:"type:int;not null"`
	ParentID   int     `gorm:"type:int;not null"`
}

//CashFlowTransactionDetails Cash Flow Details
type CashFlowTransactionDetails struct {
	ID             int    `gorm:"type:int AUTO_INCREMENT;not null"`
	Name           string `gorm:"type:varchar(255);not null"`
	Price          int    `gorm:"type:int;not null"`
	TransactionsID int    `gorm:"type:int;not null"`
}

//Hero Hero
type Hero struct {
	ID   int
	Name string
	Job  string
}
