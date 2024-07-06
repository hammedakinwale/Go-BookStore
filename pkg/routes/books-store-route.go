package routes

import(
	"github.com/gorilla/mux"
	"github.com/hammed/Go-BookStore/pkg/controllers"
)


// books routes
var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("Get")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}