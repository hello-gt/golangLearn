package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3014159
	mapCreated["two"] = 3

	fmt.Printf("Map literal at \"one\" is :%d\n", mapLit["one"])
	fmt.Printf("Map literal at \"key2\" is :%d\n", mapLit["key2"])
	fmt.Printf("Map literal at \"two\" is :%d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is :%d\n", mapLit["ten"])
}
