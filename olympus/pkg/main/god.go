package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type God struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OtherNames  string `json:"other_names"`
}

/*
* Functions for REST API endpoints
 */
func getGods(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description FROM gods")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var gods []God
	for rows.Next() {
		var god God
		err := rows.Scan(&god.ID, &god.Name, &god.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		gods = append(gods, god)
	}

	json.NewEncoder(w).Encode(gods)
}
func addGod(w http.ResponseWriter, r *http.Request) {
	var god God
	err := json.NewDecoder(r.Body).Decode(&god)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO gods (name, description) VALUES ($1, $2)", god.Name, god.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "God added successfully")
}

func updateGod(w http.ResponseWriter, r *http.Request) {
	var god God
	err := json.NewDecoder(r.Body).Decode(&god)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE gods SET name=$1, description=$2 WHERE id=$3", god.Name, god.Description, god.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "God updated successfully")
}

func deleteGod(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM gods WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "God deleted successfully")
}
