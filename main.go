package main

import (
	"fmt"
	"log"
	"ungraded-challenge-4/config"
	"ungraded-challenge-4/handler"
)

func main() {
	router, server := config.SetupServer()
	db := &handler.NewRecordHandler{DB: config.GetDatabase()}

	router.GET("/records", db.GetCriminalRecord)
	router.GET("/records/:id", db.GetCriminalRecordById)
	router.POST("/records", db.AddNewCriminalRecord)
	router.PUT("/records/:id", db.UpdateCriminalRecord)
	router.DELETE("/records/:id", db.DeleteCriminalRecord)

	fmt.Println("Server running on port :8080")
	log.Fatal(server.ListenAndServe())
}
