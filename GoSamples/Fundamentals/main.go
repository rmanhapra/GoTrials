package main

import (
	"Fundamentals/packages"
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	packages.AddToMap("A", "First")
	packages.AddToMap("B", "Second")
	packages.AddToMap("C", "Third")
	//Add a dupplicate
	packages.AddToMap("A", "Third")

	packages.DeleteFromMap("A")
	//Delete a key not exist
	packages.DeleteFromMap("E")

	str1 := packages.GetFromMap("B")
	fmt.Println(str1)
	//Try get a non existent key
	str2 := packages.GetFromMap("E")
	if len(str2) == 0 {
		fmt.Println("key dont exist in map")
	}

	//Get all data from map
	for k, v := range packages.GetAllFromMap() {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}

}
