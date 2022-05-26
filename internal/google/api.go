package google

import (
	"iot/pkg/google"
	"iot/pkg/middleware"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewGoogle(router *mux.Router,
	middleware *middleware.Middleware,
	calendarGoogle *google.CalendarGoogle,
) {
	handler := NewGoogleHandler(calendarGoogle)
	SetRoutes(handler, router, middleware)
}

func SetRoutes(
	handler *GoogleHandler,
	router *mux.Router,
	mw *middleware.Middleware,
) {
	r := router.PathPrefix("/v1/google").Subrouter()

	rCalendar := r.PathPrefix("/calendar").Subrouter()
	rCalendar.Handle(
		"/events",
		handlers.CompressHandler(
			mw.HandlerError(handler.GetCalendarEvents),
		),
	).Methods(http.MethodGet)
}
