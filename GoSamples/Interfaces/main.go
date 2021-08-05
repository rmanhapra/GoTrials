package main

import (
	"Interfaces/mapstore"
	"Interfaces/models"
	"fmt"
)

type CustomerController struct {
	DataStore models.CustomerStore
}

func (c CustomerController) Add(cust models.Customer) {
	x := c.DataStore.Create(cust)
	if x != nil {
		fmt.Println("Error creating customer")
	} else {
		fmt.Println("Customer creating successfull")
	}
	return
}

func (c CustomerController) GetAll() {
	custColl, _ := c.DataStore.GetAll()
	for _, v := range custColl {
		fmt.Println(v)
	}
}

func (c CustomerController) GetByKey(key string) {
	custColl, _ := c.DataStore.GetByKey(key)
	fmt.Println("Customer for key ", key, custColl)

}

func (c CustomerController) UpdateCustomer(key string, cust models.Customer) {
	error := c.DataStore.Update(cust, key)
	if error != nil {
		fmt.Println("Updation Unsuccessfull")
	} else {
		fmt.Println("Updation Successfull")
	}
}

func (c CustomerController) DeleteCustomer(key string) {
	error := c.DataStore.Delete(key)
	if error != nil {
		fmt.Println("Deletion Unsuccessfull")
	} else {
		fmt.Println("Deletion Successfull")
	}
}

func main() {

	fmt.Println("starting main")
	cust := models.Customer{
		"1", "@yahoo", "One",
	}

	storeInst := mapstore.New()
	custController := CustomerController{DataStore: storeInst}

	//Add some customers
	custController.Add(cust)
	custController.Add(models.Customer{"2", "@gmail", "two"})
	custController.Add(models.Customer{"3", "@live", "three"})

	//Get values
	fmt.Println("\n")
	custController.GetAll()
	fmt.Println("\n")
	custController.GetByKey("2")

	// Update and Get Values
	fmt.Println("\n")
	custController.UpdateCustomer("2", models.Customer{"2", "updEmail", "updName"})
	custController.GetByKey("2")

	//delete
	fmt.Println("\n")
	custController.DeleteCustomer("2")
	custController.GetAll()

}
