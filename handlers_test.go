package main

import (
    //"io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
	"fmt"
	"regexp"
	"strings"
    
    "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var u = Employee{
	EmployeeID: 1,
	FullName:  "Momo",
	Birthday: "1990-3-2",
	Gender: "M",
}

var e = Event{
	EventID: 1,
	Name: "Party",
	EventDate: "2020-12-2",
}

func TestGetEmployees(t *testing.T) {

    // setup Mock
    db, mock, err := sqlmock.New()
	//repo := &repository{db}
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

    query := "SELECT * FROM employees"

    rows := sqlmock.NewRows([]string{"id", "fullname", "birthday", "gender", "eventId", "accomodation"}).AddRow(u.EmployeeID, u.FullName, u.Birthday, u.Gender, 1, 'Y')

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)


    req := httptest.NewRequest(http.MethodGet, "/employees", nil)
    w := httptest.NewRecorder()
    GetEmployees(w, req, db)
    res := w.Result()
    defer res.Body.Close()
    //data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    // ensure all expectations have been met
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}
}

func TestGetEvents(t *testing.T) {

    // setup Mock
    db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

    query := "SELECT * FROM events"

    rows := sqlmock.NewRows([]string{"id", "name", "eventDate"}).AddRow(e.EventID, e.Name, e.EventDate)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)


    req := httptest.NewRequest(http.MethodGet, "/events", nil)
    w := httptest.NewRecorder()
    GetEvents(w, req, db)
    res := w.Result()
    defer res.Body.Close()
    //data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Error: %v", err)
    }

    // ensure all expectations have been met
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}
}

func TestAddEmployeeSuccess(t *testing.T){
	// setup Mock
    db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query:="INSERT INTO employees"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(u.FullName, u.Birthday, u.Gender).WillReturnResult(sqlmock.NewResult(1,1))
	_, err = db.Exec("INSERT INTO employees(fullname, birthday, gender) VALUES (?, ?, ?)", u.FullName, u.Birthday, u.Gender)

    w := httptest.NewRecorder()
	reader := strings.NewReader("fullname="+u.FullName+",birthday="+u.Birthday+",gender="+u.Gender)
    req := httptest.NewRequest(http.MethodPost, "/employees", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	AddEmployee(w, req, db)
    res := w.Result()
    defer res.Body.Close()
    //data, err := ioutil.ReadAll(res.Body)

    // ensure all expectations have been met
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Printf("unmet expectation error: %s", err)
	}
}
