package route

import (
	"management/internal/http/handler"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleAPI(port, rootPath string) {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	
	s.Handle("/tickets", &handler.TicketHandler{}).Methods("POST")

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":"+port, handler)
}





