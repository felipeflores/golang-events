package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	internalAdmin "iot/internal/admin"
	"iot/internal/config"
	internalGoogle "iot/internal/google"
	"iot/internal/google/calendar"
	"iot/pkg/google"
	"iot/pkg/middleware"
	"iot/pkg/mqtt"
	"iot/rest"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "c", "env.yaml", "env file path")
	flag.Parse()

	c, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error on load config file %v", err)
	}

	mw := middleware.New()

	clientGoogle := google.NewClientGoogle()
	calendarGoogle := google.NewCalendarGoogle(clientGoogle)

	mqttClient, err := mqtt.NewMqttService(c.Mqtt.ClientID, c.Mqtt.Broker, c.Mqtt.Port)
	if err != nil {
		panic(err)
	}

	// Init the mux router
	router := mux.NewRouter()

	rest.SetupRoutes(router)

	internalAdmin.NewAdminApi(internalAdmin.Config{
		Router:         router,
		Middleware:     mw,
		CalendarGoogle: calendarGoogle,
	})

	service := calendar.NewCalendarService(calendarGoogle, mqttClient, c.Mqtt.Event)

	go service.GetEvents()

	internalGoogle.NewGoogle(router, mw, calendarGoogle, service)

	err = http.ListenAndServe(c.Server.Address, router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Subiu")

	// router.PathPrefix("/health").
	// 	Methods(http.MethodGet).
	// 	Handler(handlers.CompressHandler)
}
