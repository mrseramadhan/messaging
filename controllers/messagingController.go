package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ramvasanth/wavecell"

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
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
	}

	message := models.GetMessaging(id)

	if message == nil {
		u.Respond(w, u.Message(false, "Message Not found"))
		return
	}

	if message.Type == "SMS" {

		mBody := wavecell.Message{
			From: "WL INFO",
			To:   message.DestinationID,
			Text: message.MessageBody,
		}

		err := mBody.Validate()

		if err != nil {
			u.Respond(w, u.Message(false, "Error Validate Message to wavecell"))
		}
		client := wavecell.ClientWithAuthKey(os.Getenv("wave_cell_key"), message.MessageType)

		r, err := client.SingleMessage(mBody)
		if (err) != nil {
			u.Respond(w, u.Message(false, "Error Sending Message to Wavecell"))
			return
		}

		resp := u.Message(true, "success")

		resp["data"] = r
		u.Respond(w, resp)
	} else if message.Type == "EMAIL" {
		mBody := 
	}

	resp := u.Message(true, "success")
	resp["data"] = message.Type
	u.Respond(w, resp)
	return
}
