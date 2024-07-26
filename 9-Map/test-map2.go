package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key=", key, "value=", value)
	}

}

func changeMap(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {

	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// for key, value := range cityMap {
	// 	fmt.Println("key=", key, "value=", value)
	// }
	printMap(cityMap)

	delete(cityMap, "China")

	cityMap["USA"] = "DC"
	changeMap(cityMap)

	fmt.Println("After delete")
	// for key, value := range cityMap {
	// 	fmt.Println("key=", key, "value=", value)
	// }
	printMap(cityMap)
}
