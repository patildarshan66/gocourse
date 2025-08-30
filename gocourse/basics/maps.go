package basics

import (
	"fmt"
	"maps"
)

func main() {

	mapData := map[string]int{
		"key1": 20,
		"key2": 30,
	}

	fmt.Println(mapData)

	var mapData2 map[string]string
	mapData2 = make(map[string]string)
	mapData2["key1"] = "value1"
	mapData2["key2"] = "value2"

	fmt.Println(mapData2)

	mapData2["key2"] = "Darshan"
	mapData2["key3"] = "Golang"

	fmt.Println(mapData2)

	v, isValueAvailable := mapData["key1"]

	if isValueAvailable {
		fmt.Println("isValueAvailable:", isValueAvailable)
		fmt.Println("Value:", v)
	}

	var myMap map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println(myMap)

	delete(myMap, "two")
	clear(mapData2)
	fmt.Println(mapData2)
	fmt.Println(myMap)

	var myMap1 map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	var myMap2 map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	if maps.Equal(myMap1, myMap2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	for k, v := range myMap1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	var myMap3 map[string]int

	if myMap3 == nil {
		fmt.Println("myMap3 is nil")
	} else {
		fmt.Println("myMap3 is not nil")
	}

	myMap3 = make(map[string]int)
	myMap3["four"] = 4
	myMap3["five"] = 5
	myMap3["six"] = 6

	fmt.Println("myMap3:", myMap3)

	fmt.Println("myMap3 length:", len(myMap3))

	myMap4 :=
		map[string]map[string]int{
			"key1": myMap1,
			"key2": myMap2,
			"key3": myMap3,
		}

	fmt.Println(myMap4)

	myMap5 := make(map[string]map[string]int)
	myMap5["key1"] = myMap1
	myMap5["key2"] = myMap2
	myMap5["key3"] = myMap3

	fmt.Println(myMap5)
}
