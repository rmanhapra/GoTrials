package mapstore

import (
	"Interfaces/models"
	"errors"
)

type MapStore struct {
	Store map[string]models.Customer
}

func New() *MapStore {
	return &MapStore{Store: make(map[string]models.Customer)}
}

func (st MapStore) Create(cust models.Customer) error {
	if _, ok := st.Store[cust.Id]; ok {
		return errors.New("Cannot make customer")
	} else {
		st.Store[cust.Id] = cust
	}
	return nil
}

func (st MapStore) Update(cust models.Customer, key string) error {
	if _, ok := st.Store[key]; ok {
		st.Store[key] = cust

	} else {
		return errors.New("Cannot Update customer")
	}
	return nil
}

func (st MapStore) Delete(key string) error {
	if _, ok := st.Store[key]; ok {
		delete(st.Store, key)
	} else {
		return errors.New("Cannot delete customer")
	}
	return nil
}

func (st MapStore) GetByKey(key string) (models.Customer, error) {
	if val, ok := st.Store[key]; ok {
		return val, nil
	} else {
		return val, errors.New("No customer exists")
	}

}

func (st MapStore) GetAll() ([]models.Customer, error) {
	custSlice := make([]models.Customer, 3)
	i := 0
	for _, v := range st.Store {
		if i > 2 {
			custSlice = append(custSlice, v)
		} else {
			custSlice[i] = v
		}

		i++
	}
	return custSlice, nil

}
