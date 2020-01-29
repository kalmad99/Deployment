package bank

import "github.com/kalmad99/Deployment/entity"

//PayService specifies user payment related services
type PayService interface {
	MakePayment(accno string, amount float64) (*entity.Bank, []error)
	BankExists (accno string) bool
}
