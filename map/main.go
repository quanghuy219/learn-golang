package main

import "fmt"

func main()  {
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#1111",
		"black": "#000000",
	}
	colors["white"] = "#ffffff"
	delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string)  {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
