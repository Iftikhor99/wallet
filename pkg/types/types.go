package types

//Money for
type Money int64
//PaymentCategory for
type PaymentCategory string
//PaymentStatus for
type PaymentStatus string

const (
	//PaymentStatusOk for 
	PaymentStatusOk PaymentStatus = "OK"
	//PaymentStatusFail for
	PaymentStatusFail PaymentStatus = "FAIL" 
	//PaymentStatusInProgress for
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment for
type Payment struct {
	ID string 
	AccountID int64
	Amount Money
	Category PaymentCategory
	Status PaymentStatus
}
//Favorite for
type Favorite struct {
	ID string 
	AccountID int64
	Name string
	Amount Money
	Category PaymentCategory
}

//Phone for
type Phone string

//Account for
type Account struct {
    ID int64
    Phone Phone
    Balance Money 
}