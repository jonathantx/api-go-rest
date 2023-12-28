package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathantx/go-rest-api/controllers"
	"github.com/jonathantx/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.ReturnPersonality).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPesonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
