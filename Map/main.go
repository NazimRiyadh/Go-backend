package main

import (
	"fmt"
	"maps"
)

func main() {
	//map creation
	m := map[string]int{
		"apple":  18,
		"banana": 10,
		"lichi":  5,
	}

	//adding new element
	m["orange"] = 15

	//iteratiing over the map
	for idx, val := range m { //idx->key, val->value
		fmt.Println(idx, "value: ", val)
	}

	fmt.Println(m["lichi"])

	//map creation using make
	n := make(map[string]int)

	n["lichi"] = 20
	n["Bichi"] = 10
	n["Kechi"] = 5

	//if any value not found for a key, will return 0/false according to the value type
	fmt.Println(n["lacchi"]) //will return 0, as value type is 0

	clear(n)

	fmt.Println(n["lichi"]) //will return 0 as i cleared the map

	map1 := map[string]string{
		"name":    "Riyadh",
		"country": "Bangladesh",
		"city":    "Dhaka",
	}
	map2 := map[string]string{
		"name":    "Habib",
		"country": "India",
		"city":    "Mumbai",
	}

	// Comparing maps
	//fmt.Println("Are maps equal?", map1 == map2) -> This will cause a compilation error
	fmt.Println(maps.Equal(map1, map2)) // Using reflect.DeepEqual to compare maps
}
