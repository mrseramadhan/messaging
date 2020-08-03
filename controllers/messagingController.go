package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
)

// CreateMessage this func will create a message and save it to scheduler
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	messaging := &models.MessagingModel{}

	err := json.NewDecoder(r.Body).Decode(messaging)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	defer r.Body.Close()

	resp := messaging.CreateMessaging()
	u.Respond(w, resp)
}

// SendMessage this func will retrieve data from scheduler and send to other message provider.
func SendMessage(w http.ResponseWriter, r *http.Request) {
	messaging := &models.MessagingModel{}

	err := json.NewDecoder(r.Body).Decode(messaging)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	defer r.Body.Close()

	resp := messaging.CreateMessaging()
	u.Respond(w, resp)
}
