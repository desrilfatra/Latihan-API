package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Biodata struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

var mapEmployees = map[int]Biodata{
	1: {Id: 1, Username: "Desril", Email: "desrilfatra@gmail.com", Password: "123456", Age: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	2: {Id: 2, Username: "Arief", Email: "arief@gmail.com", Password: "123456", Age: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	3: {Id: 3, Username: "Usama", Email: "usama@gmail.com", Password: "123456", Age: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/user/", getEmployees)
	// http.HandleFunc("/user/", getEmployeesId)

	fmt.Println("Server is running on port " + PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	paths := strings.Split(r.URL.Path[1:], "/")
	fmt.Printf("%+v", paths)
	fmt.Println(len(paths))

	switch r.Method {
	case "GET":
		if paths[1] != "" {
			fmt.Println(paths[1])
			fmt.Println("validasi 1")
			if Id, err := strconv.Atoi(paths[1]); err == nil {
				fmt.Println("validasi 2")
				jsonData, _ := json.Marshal(mapEmployees[Id])
				fmt.Println(mapEmployees[Id])
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			} else {
				fmt.Println(err)
				fmt.Println("Error")
			}
		} else {
			fmt.Println("no parama")
			jsonData, _ := json.Marshal(mapEmployees)
			w.Header().Add("Content-Type", "application/json")
			w.Write(jsonData)
		}
	case "POST":
		var employee Biodata
		json.NewDecoder(r.Body).Decode(&employee)
		employee.Id = len(mapEmployees) + 1
		employee.CreatedAt = time.Now()
		employee.UpdatedAt = time.Now()
		mapEmployees[employee.Id] = employee
		jsonData, _ := json.Marshal(mapEmployees)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	case "PUT":
		var employee Biodata
		json.NewDecoder(r.Body).Decode(&employee)
		employee.CreatedAt = time.Now()
		employee.UpdatedAt = time.Now()
		mapEmployees[employee.Id] = employee
		jsonData, _ := json.Marshal(mapEmployees)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	case "DELETE":
		if paths[1] != "" {
			fmt.Println(paths[1])
			if Id, err := strconv.Atoi(paths[1]); err == nil {
				delete(mapEmployees, Id)
				jsonData, _ := json.Marshal(mapEmployees)
				
				w.Header().Add("Content-Type", "application/json")
				w.Write(jsonData)
			}
		}
	}
}
