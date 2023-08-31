package main

import "buffe-map.com/buffe-map-api/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
