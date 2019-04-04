package main

import "fmt"

func main() {
	fmt.Println("Maps")
	colors := map[string]string{
		"Red":   "#12345",
		"Green": "#345",
	}

	fmt.Printf("%+v", colors)
}
