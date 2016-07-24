package main

import (
"fmt"
"encoding/json"
"io/ioutil"
"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func saveUser(u *User) error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("user.txt", b, 0600)
}

func loadData() ([]byte, error) {
	data, err := ioutil.ReadFile("user.txt")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	if firstName == "" {
		http.Error(w, "first_name parameter missing", http.StatusBadRequest)
		return
	}
	if lastName == "" {
		http.Error(w, "last_name parameter missing", http.StatusBadRequest)
		return
	}	
	u := &User{FirstName : firstName, LastName : lastName}
	fmt.Println("Found user " + " firstName:" + firstName + " lastName:" + lastName)
	err := saveUser(u)
	if err != nil {
		http.Error(w, "Could not save data", http.StatusInternalServerError)
		return
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := loadData()
	if err != nil {
		http.Error(w, "No data found", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(":8080", nil)
}
