package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PatrikOlin/neo-burton/db"
	"github.com/go-chi/chi/v5"
)

type ToggleResponse struct {
	Status string `json:"status"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	id := chi.URLParam(r, "id")
	toMap := r.URL.Query().Has("map")

	var err error
	status, err := signIn(id, toMap)

	if err != nil {
		HandleHTTPError(w, err)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	if toMap {
		json.NewEncoder(w).Encode(status)
	} else {
		json.NewEncoder(w).Encode(ToggleResponse{Status: "SIGNED_IN"})
	}

}

func signIn(id string, toMap bool) (ToggleResponse, error) {
	// stmt := "INSERT INTO checkins (employee_id, checkin_time) VALUES ($1, $2)"
	// _, err := db.DBClient.Exec(stmt, id, time.Now())

	// if err != nil {
	// 	return err
	// }

	var status ToggleResponse
	var err error
	if toMap {
		status, err = toggleMapSignedIn(id)

		if err != nil {
			return status, err
		}

		fmt.Println("användare loggade in! " + id)
	}

	return status, nil
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

func toggleMapSignedIn(id string) (ToggleResponse, error) {
	url := "https://blmap.blinfo.se/api/map/signInOrOut/" + id
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, nil)
	var resp ToggleResponse

	if err != nil {
		return resp, err
	}

	res, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
