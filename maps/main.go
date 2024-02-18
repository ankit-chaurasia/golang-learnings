package main

import "fmt"

func main() {
	var colors1 map[string]string // It will create an empty map and with nil value
	// colors1["white"] = "#ffffff" // Empty map is not initialised so assigning to empty map will create an error
	fmt.Println(colors1) // map[]

	colors2 := make(map[string]string) // Use make method to create empty map, it will initializes a hash map and returns a map value
	colors2["white"] = "#ffffff"
	fmt.Println(colors2) // map[]

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors3)

	delete(colors3, "red")
	fmt.Println(colors3)

	colors4 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	for key, value := range colors4 {
		fmt.Println("Key(color):", key, "Value(HexCode): ", value)
	}
	printMap(colors4)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}
