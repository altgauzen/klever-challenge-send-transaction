package routes

import (
	"klever/api/controllers"
	"klever/api/middleware"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/details/{address}", controllers.Details).Methods("Get")
	r.HandleFunc("/balance/{address}", controllers.Balances).Methods("Get")
	r.HandleFunc("/send", controllers.CreateTransaction).Methods("Post")
	r.HandleFunc("/tx/{tx}", controllers.ReceiveTransaction).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
