package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/zivenlouis/futurefarmerapi/controllers/authcontroller"
	"github.com/zivenlouis/futurefarmerapi/controllers/dashboardcontroller"
	"github.com/zivenlouis/futurefarmerapi/controllers/datacontroller"
	"github.com/zivenlouis/futurefarmerapi/models"
)

func loopingFunction() {
	i := 0
    for {
		i++
        fmt.Println("Looping - " + strconv.Itoa(i))

		var relayStatusData []models.RelayStatus
		models.DB.Find(&relayStatusData)
	
		for _, status := range relayStatusData {
			if status.Status == 1 && time.Since(status.UpdatedAt).Seconds() >= float64(status.Duration) {
                status.Status = 0
				fmt.Println("Turned off relay status id = " + strconv.Itoa(int(status.Id)))
                models.DB.Save(&status)
            }
		}

		// Ini untuk testing aja random on
		for _, status := range relayStatusData {
			if status.Status == 0 &&  rand.Intn(20) == 1  {
                status.Status = 1
				fmt.Println("Turning on relay status id = " + strconv.Itoa(int(status.Id)) + " for " + 
					strconv.Itoa(int(status.Duration)) + " seconds")
                models.DB.Save(&status)
            }
		}


        time.Sleep(time.Second)
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
