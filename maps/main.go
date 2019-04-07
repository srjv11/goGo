package main

import "fmt"

func main() {
	fmt.Println("Maps")
	colors := map[string]string{
		"Red":   "#12345",
		"Green": "#345",
	}

	colors["yellow"] = "#DE4G5C"

	delete(colors, "yellow")

	//fmt.Printf("%+v", colors)
	print(colors)
}

func print(m map[string]string) {
	for ckey, cval := range m {
		fmt.Println("key = ", ckey, "& values = ", cval)
	}
}
