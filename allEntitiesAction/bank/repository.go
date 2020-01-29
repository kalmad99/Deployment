package bank

import "github.com/kalmad99/Deployment/entity"

// PayRepository specifies payment related database operations
type PayRepository interface {
	MakePayment(accno string, amount float64) (*entity.Bank, []error)
	BankExists (accno string) bool
}
