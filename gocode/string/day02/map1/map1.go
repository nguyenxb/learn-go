package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// insert several(几个) pieces（块） of data（数据块） into the map
	personDB["1234"] = PersonInfo{"1234", "Tom", "Room 101,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 203,..."}
	personDB["222"] = PersonInfo{"789", "hasib", "Room 901,..."}

	// Look for information of a key "1234" in this map
	person, ok := personDB["1234"]
	// ok is a returned bool.Returns true to indicate(表示) that the corresponding（相应的） data was found
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
	//--------------------------------------------------------------

	var myMap map[string]PersonInfo

	myMap = make(map[string]PersonInfo)
	fmt.Println(myMap)

	myMap = make(map[string]PersonInfo, 100)
	fmt.Println(myMap)

	// var myMap3 map[string]PersonInfo
	myMap3 := map[string]PersonInfo{
		"1234": PersonInfo{"1", "Jack", "Room 101,..."},
	}
	fmt.Println(myMap3)
	a, b := myMap["1"]
	fmt.Println(a, ":", b)
	delete(myMap3, "1")
}
