package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// func SendJSON(w http.ResponseWriter, payload interface{}) error {
// 	return WriteJSON(context.Background(), w, payload)
// }

// func WriteJSON(ctx context.Context, w http.ResponseWriter, payload interface{}, httpStatusCode ...int) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	if len(httpStatusCode) > 0 {
// 		w.WriteHeader(httpStatusCode[0])
// 	}

// 	if err := json.NewDecoder(w).Encode(payload); err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return errors.New("errou")
// 	}
// 	return nil
// }

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)
}
