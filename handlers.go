package main

import (
	"fmt"
	"net/http"
	"encoding/json"
    "database/sql"

    "github.com/gorilla/mux"
)

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// Function for handling messages
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// Function to Get all Employees

// response and request handlers
func GetEmployees(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    
    printMessage("Getting employees...")

    // Get all employees from employees table
    rows, err := db.Query("SELECT * FROM employees")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var employees []Employee
    counter := 0
    // Foreach movie
    for rows.Next() {
        var id int
        var fullname string
        var birthday string
        var gender string
        var eventId *int
        var accomodation *string

        err = rows.Scan(&id, &fullname, &birthday, &gender, &eventId, &accomodation)

        // check errors
        checkErr(err)

        employees = append(employees, Employee{EmployeeID: id, FullName: fullname, Birthday: birthday, Gender: gender})
        counter++
    }
    var response = JsonResponse{}
    if(counter == 0){
        response = JsonResponse{Type: "warning", Data: nil, Message: "No Records Found"}
    } else {
        printMessage("Employees Record fetched...")

        response = JsonResponse{Type: "success", Data: employees, Message: "All Employees Record"}
    }
    json.NewEncoder(w).Encode(response)
}

// Add New Employee
// response and request handlers
func AddEmployee(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    fullName := r.FormValue("fullname")
    birthday := r.FormValue("birthday")
    gender := r.FormValue("gender")

    var response = JsonResponse{}

    if fullName == "" || birthday == "" || gender == ""   {
        response = JsonResponse{Type: "error", Message: "You are missing one or more required parameter {id, fullname, birthday, gender}"}
    } else {

        printMessage("Opening DB")
        
        db := setupDB()

        printMessage("Adding employee into DB")

        fmt.Println("Adding new employee")

        var lastInsertID int
        err := db.QueryRow("INSERT INTO employees(fullname, birthday, gender) VALUES($1, $2, $3) returning id;", fullName, birthday, gender).Scan(&lastInsertID)

        if err != nil{
            response = JsonResponse{Type: "error", Message: err}
            json.NewEncoder(w).Encode(response)
            // check errors
            checkErr(err)
        } else {
            printMessage("Record inserted ")
            response = JsonResponse{Type: "success", Message: "The employee has been inserted successfully!"}
            json.NewEncoder(w).Encode(response)
        }
    }
    
}

// Update an Employee
// response and request handlers
func UpdateEmployee(w http.ResponseWriter, r *http.Request, db *sql.DB){

    vars := mux.Vars(r)
    employeeId := vars["employee_id"]

    fullName := r.FormValue("fullname")
    birthday := r.FormValue("birthday")
    gender := r.FormValue("gender")
    eventId := r.FormValue("eventId")
    accomodation := r.FormValue("accomodation")
    printMessage(eventId)
    var response = JsonResponse{}

    if employeeId != "" {
    
            printMessage("Updating employee with ID: " + employeeId)
    
            _, err := db.Exec(`UPDATE employees SET fullname = COALESCE($1, fullname), birthday = COALESCE($2, birthday), gender = COALESCE($3, gender), eventId = COALESCE($4, eventId), accomodation = COALESCE($5, accomodation) WHERE id = $6;`, NewNullString(fullName), NewNullString(birthday), NewNullString(gender), NewNullString(eventId), NewNullString(accomodation), employeeId)
            
            if(err == nil) {
                printMessage("Record updated")
                response = JsonResponse{Type: "success", Message: "The employee record has been updated successfully!"}
            } else if err == sql.ErrNoRows{
                response = JsonResponse{Type: "error", Message: "No rows found"}
            } else {
                checkErr(err)
            }
    } else {
        response = JsonResponse{Type: "error", Message: "You are missing employee id, please provide"}
    }

    json.NewEncoder(w).Encode(response)
        
}

// Delete an Employee
// response and request handlers
func DeleteEmployee(w http.ResponseWriter, r *http.Request,db *sql.DB){

    vars := mux.Vars(r)
    employeeId := vars["employee_id"]
    
    var response = JsonResponse{}

    if employeeId != "" {
    
        printMessage("Deleting employee with ID: " + employeeId)
    
        _, err := db.Exec("Delete FROM employees WHERE id = $1 returning id;", employeeId)
    
        // check errors
        checkErr(err)
            
        if(err == nil) {
            printMessage("Record deleted ")
            response = JsonResponse{Type: "success", Message: "The employee record has been deleted successfully!"}
        } else {
                response = JsonResponse{Type: "error", Message: "Error encountered"}
        }
    } else {
        response = JsonResponse{Type: "error", Message: "You are missing employee id, please provide"}
    }

    json.NewEncoder(w).Encode(response)
        
}



//Function to Get all Events

//response and request handlers
func GetEvents(w http.ResponseWriter, r *http.Request, db *sql.DB) {

    printMessage("Getting events...")

    //Get all events from events table 
    rows, err := db.Query("SELECT * FROM events")

    //check errors
    checkErr(err)

    var response = JsonResponse{}
    var events []Event

    //Foreach movie
    for rows.Next() {
        var id int
        var name string
        var eventDate string

        err = rows.Scan(&id, &name, &eventDate)

        //check errors
        checkErr(err)

        events = append(events, Event{EventID: id, Name: name, EventDate: eventDate})
    }

    printMessage("Events Record fetched...")

    response = JsonResponse{Type: "success", Data: events, Message: "All Employees Record"}

    json.NewEncoder(w).Encode(response)
}

// response and request handlers
func GetEvent(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    vars := mux.Vars(r)
    eventId := vars["event_id"]

    printMessage("Getting events...")

    // Get all events from events table 
    row := db.QueryRow("SELECT * FROM events WHERE id=$1", eventId)

    // var response []JsonResponse
    var events []Event

    var id int
    var name string
    var eventDate string

    err := row.Scan(&id, &name, &eventDate)

    // check errors
    checkErr(err)

    events = append(events, Event{EventID: id, Name: name, EventDate: eventDate})

    printMessage("Event Record fetched...")

    var response = JsonResponse{Type: "success", Data: events, Message: "Specific Event Record"}

    json.NewEncoder(w).Encode(response)
}

func EmployeeInEvent (w http.ResponseWriter, r *http.Request, db *sql.DB){
    vars := mux.Vars(r)
    eventId := vars["event_id"]

    accomodationFlag := r.URL.Query().Get("accomodation")

    printMessage("Getting employees...")
    sqlStatment := `SELECT * FROM employees WHERE eventId=$1`
    if accomodationFlag == "t"{
        sqlStatment = `SELECT * FROM employees WHERE eventId=$1 AND accomodation='Y'`
    }

    // Get all events from events table 
    rows, err := db.Query(sqlStatment, eventId)
    checkErr(err)

    // var response []JsonResponse
    var employees []Employee
    count := 0
    // Foreach movie
    for rows.Next() {
        
        var id int
        var fullname string
        var birthday string
        var gender string
        var eventID int
        var accomodation string

        err := rows.Scan(&id, &fullname, &birthday, &gender, &eventID, &accomodation)

        // check errors
        checkErr(err)

        employees = append(employees, Employee{EmployeeID: id, FullName: fullname, Birthday: birthday, Gender: gender})
        count++
    }
    if(count == 0){
        var response = JsonResponse{Type: "success", Data: employees, Message: "No records found"}

        json.NewEncoder(w).Encode(response) 
    } else {
        printMessage("Event Record fetched...")

        var response = JsonResponse{Type: "success", Data: employees, Message: "All employees going to event"}

        json.NewEncoder(w).Encode(response)
    }
}