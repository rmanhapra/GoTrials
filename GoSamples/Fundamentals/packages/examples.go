package packages

import "fmt"

var mapData map[string]string

func init() {
	mapData = make(map[string]string)
	fmt.Println("Map initialised")
}

func AddToMap(k, v string) {
	if _, ok := mapData[k]; ok {
		fmt.Println("Cannot Add, Key - Value already exists")
	} else {
		mapData[k] = v
		fmt.Println("Key - Value Added")
	}

}

func DeleteFromMap(k string) {
	delete(mapData, k)
	fmt.Println("Deleted from Map")
}

func GetFromMap(k string) string {
	return mapData[k]
}

func GetAllFromMap() map[string]string {
	return mapData
}
