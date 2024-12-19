package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1ef7yy/devops-course/internal/models"
)

func (v *view) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		v.Logger.Info("Name is empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Name is empty"))
		return
	}

	data, err := v.domain.GetUserByName(name)
	if err != nil {
		v.Logger.Error("Error getting user: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: %s", err.Error())
		return
	}

	resp, err := json.Marshal(&data)
	if err != nil {
		v.Logger.Error("Error marshaling data: " + err.Error())
		v.Logger.Debug(fmt.Sprintf("Data: %v", data))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (v *view) InsertUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		v.Logger.Error("Error decoding user: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request: %s", err.Error())
		return
	}

	err = v.domain.InsertUser(user)
	if err != nil {
		v.Logger.Error("Error inserting user: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
