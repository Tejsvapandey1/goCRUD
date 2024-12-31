package routes

import (
	"github.com/Tejsvapandey1/goCRUD.git/pkg/controller"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func (r *mux.Router){
	
    r.HandleFunc("/book/", controller.CreateBook).Methods("POST")
    r.HandleFunc("/book/", controller.Getbooks).Methods("GET")
    r.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
	r.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("UPDATE")
}	