package handlers

import (
	"net/http"
)

func HandleHTTPError(w http.ResponseWriter, err error) {
	// if err == gorm.ErrRecordNotFound {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
