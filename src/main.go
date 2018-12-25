package main

import "route"

// "encoding/json"
func main() {
	router := route.Router()
	router.Run(":9090")
}
