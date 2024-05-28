package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/jovinkendrico/futurefarmerapi/controllers/authcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/dashboardcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/datacontroller"
	"github.com/zivenlouis/futurefarmerapi/models"
)

func loopingFunction() {
    for {
        fmt.Println("Looping...")
        time.Sleep(time.Second) // Sleep for one second before looping again
    }
}


func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	go loopingFunction()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/insertdata", datacontroller.InsertData).Methods("POST")
	r.HandleFunc("/api/v1/dashboard", dashboardcontroller.Index).Methods("GET")
	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
