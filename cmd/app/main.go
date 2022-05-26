package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	internalAdmin "iot/internal/admin"
	internalGoogle "iot/internal/google"
	"iot/internal/google/calendar"
	"iot/pkg/google"
	"iot/pkg/middleware"
	"iot/rest"
)

func main() {
	mw := middleware.New()

	clientGoogle := google.NewClientGoogle()
	calendarGoogle := google.NewCalendarGoogle(clientGoogle)

	// Init the mux router
	router := mux.NewRouter()

	rest.SetupRoutes(router)

	fmt.Println("Passou aqui")

	internalAdmin.NewAdminApi(internalAdmin.Config{
		Router:         router,
		Middleware:     mw,
		CalendarGoogle: calendarGoogle,
	})

	service := calendar.NewCalendarService(calendarGoogle)

	go service.GetEvents()

	internalGoogle.NewGoogle(router, mw, calendarGoogle)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
	fmt.Println("Subiu")

	// router.PathPrefix("/health").
	// 	Methods(http.MethodGet).
	// 	Handler(handlers.CompressHandler)
}
