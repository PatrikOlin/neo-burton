package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PatrikOlin/neo-burton/db"
	"github.com/go-chi/chi/v5"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	id := chi.URLParam(r, "id")
	isHex := r.URL.Query().Get("hex")

	fmt.Println(isHex)
	var err error
	if isHex == "true" {
		err = signInWithHex(id)
	} else {
		err = signIn(id)
	}

	if err != nil {
		HandleHTTPError(w, err)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Signed in")
}

func signIn(id string) error {
	stmt := "INSERT INTO checkins (employee_id, checkin_time) VALUES ($1, $2)"
	_, err := db.DBClient.Exec(stmt, id, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func signInWithHex(hex string) error {
	id, err := getUserID(hex)
	if err != nil {
		return err
	}

	stmt := "INSERT INTO checkins (employee_id, checkin_time) VALUES ($1, $2)"
	_, err = db.DBClient.Exec(stmt, id, time.Now())

	if err != nil {
		return err
	}

	err = signInToMap(id)

	if err != nil {
		return err
	}

	return nil
}

func getUserID(hex string) (string, error) {
	var id string
	stmt := "SELECT employee_id FROM users WHERE nfc_hex = $1"
	err := db.DBClient.Get(&id, stmt, hex)
	if err != nil {
		return id, err
	}

	return id, nil
}

func signInToMap(id string) error {
	url := "https://blmap.blinfo.se/api/map/signIn/" + id
	_, err := http.Post(url, "application/json", nil)

	if err != nil {
		return err
	}

	return nil
}
