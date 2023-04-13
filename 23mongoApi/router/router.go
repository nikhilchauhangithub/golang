package router

import (
	"github.com/gorilla/mux"
	"github.com/nikhilchauhangithub/mongoApi/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter()

router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
router.HandleFunc("/api/Deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}

//In Go, a router is a key component of building web applications. It is responsible for matching incoming HTTP requests to their corresponding handler functions. The Mux router is a powerful and flexible router that allows for advanced routing configurations.By creating a new instance of the Mux router, you can add routes to it using methods such as HandleFunc or PathPrefix, and then start listening for incoming HTTP requests on a specified port using the ListenAndServe method. 