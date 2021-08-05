package models

type Customer struct {
	Id, Email, Name string
}

type CustomerStore interface {
	Create(Customer) error
	Update(Customer, string) error
	Delete(string) error
	GetByKey(string) (Customer, error)
	GetAll() ([]Customer, error)
}
