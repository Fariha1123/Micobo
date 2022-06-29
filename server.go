package main
import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main(){
	// setup MUX router
	router:=mux.NewRouter()
	fmt.Println("router setup ")

	// setup db
	db:=setupDB()
	// route handles and endpoints

	// Register a new employee
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		AddEmployee(w, r, db)
	}).Methods("POST")

	// Get all employees
	//router.HandleFunc("/employees", repo.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		GetEmployees(w, r, db)
	}).Methods("GET")

	// Update employee information
	//router.HandleFunc("/employees/{employee_id}", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employees/{employee_id}", func(w http.ResponseWriter, r *http.Request) {
		UpdateEmployee(w, r, db)
	}).Methods("PUT")

	// // delete an employee
	// router.HandleFunc("/employees/{employee_id}", DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/employees/{employee_id}", func(w http.ResponseWriter, r *http.Request) {
		DeleteEmployee(w, r, db)
	}).Methods("DELETE")

	// Get all events
	router.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		GetEvents(w, r, db)
	}).Methods("GET")

	// // Get specific event
	// router.HandleFunc("/events/{event_id}", GetEvent).Methods("GET")
    router.HandleFunc("/events/{event_id}", func(w http.ResponseWriter, r *http.Request) {
		GetEvent(w, r, db)
	}).Methods("GET")

	// // Get employees assisting to an event
	// router.HandleFunc("/events/{event_id}/employees", EmployeeInEvent).Methods("GET")
    router.HandleFunc("/events/{event_id}/employees", func(w http.ResponseWriter, r *http.Request) {
		EmployeeInEvent(w, r, db)
	}).Methods("GET")
	
	// serve the app
	fmt.Println("server is up at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

