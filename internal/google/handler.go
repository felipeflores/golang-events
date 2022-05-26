package google

import (
	"fmt"
	"iot/internal/google/calendar"
	"iot/pkg/google"
	"net/http"
)

type GoogleHandler struct {
	CalendarService *calendar.CalendarService
}

func NewGoogleHandler(cg *google.CalendarGoogle) *GoogleHandler {
	service := calendar.NewCalendarService(cg)

	return &GoogleHandler{
		CalendarService: service,
	}
}

func (h *GoogleHandler) GetCalendarEvents(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	h.CalendarService.GetEvent(ctx)
	// err =
	fmt.Println("Passou aqui222")

	// return ferrors.NewBadRequest(errors.New("Errou aqui"))
	return nil
}
