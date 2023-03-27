package main

import (
	"asidikfauzi/reservation-of-sport-fields-golang/config"
	"asidikfauzi/reservation-of-sport-fields-golang/routes"
)

func main() {
	config.InitDB()
	routes.Start()
}
