package server

import (
	"../database"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


var db *sql.DB

func callHandler(w http.ResponseWriter, r *http.Request)  {

	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		phoneNumber := vars["phoneNumber"]
		fmt.Fprintf(w, "your phone number is: "+ phoneNumber)
		database.InsertNewCall(phoneNumber)


	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
		
	}
	
}
func userHandler(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "POST":
		database.UpdateUserState(db,r.FormValue("userName"),"free")

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")

	}

}

func newUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		database.InsertNewUser(r.FormValue("userName"),r.FormValue("role"))

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")

	}
}



func Serve() {

	db = database.ConnectDB()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/calls/{phoneNumber}", callHandler).Methods("GET")
	myRouter.HandleFunc("/users/free", userHandler).Methods("POST")
	myRouter.HandleFunc("/users/NewUser", newUserHandler).Methods("POST")
	go AssignToEmployee()
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

