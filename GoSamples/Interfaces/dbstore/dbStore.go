package dbstore

import (
	"Interfaces/models"
)

//Sampple
type DBStore struct {
	//Should actuall use any DB object
	Store map[string]models.Customer
}

//Sample function
func (st DBStore) Create(cust models.Customer) error {

	return nil
}
