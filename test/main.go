package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/router"
	//mid "github.com/SPSOAFM-IT18/dmp-plant-hub/test/middleware"
)

func kokot(cPumpState chan bool) {
	//requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})

	for {
		time.Sleep(1 * time.Second)

		// var kokotismus = true

		// if kokotismus == true {
		// 	kokotismus = false
		// 	time.AfterFunc(3*time.Second, func() {
		// 		hours := time.Now().format
		// 		fmt.Println(hours)
		// 		kokotismus = true
		// 	})
		// }
		//requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
		//requests.GetLiveControl()

		fmt.Println("", <-cPumpState)

		//go kokoti()
		//go requests.PostLiveNotify(model.LiveNotify{Title: "kokot jsi", State: "active", Action: "debil"})
	}
}

func main() {
	cPumpState := make(chan bool, 1)

	r := router.Router()

	t0 := time.Now()
	time.Sleep(5 * time.Second)

	fmt.Println(int(time.Since(t0).Seconds()))

	//go sequences.MeasurementSequence(temp)

	// adc.Adc()

	// port := fmt.Sprint(":" + env.Process("GO_API_PORT"))
	// fmt.Print(port)
	fmt.Println("Starting server on the port", env.Process("GO_API_PORT"))
	go kokot(cPumpState)
	//go executeCronJob()
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r))
}
