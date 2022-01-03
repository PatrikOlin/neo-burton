package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/PatrikOlin/neo-burton/db"
	"github.com/PatrikOlin/neo-burton/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		HandleHTTPError(w, err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	err = addUser(user)

	if err != nil {
		HandleHTTPError(w, err)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func addUser(user models.User) error {
	stmt := "INSERT INTO users (employee_id, name, current_report, nfc_hex) VALUES ($1, $2, $3, $4)"
	_, err := db.DBClient.Exec(stmt, user.EmployeeID, user.Name, user.CurrentReport, user.NFCHex)

	if err != nil {
		return err
	}

	return nil
}
