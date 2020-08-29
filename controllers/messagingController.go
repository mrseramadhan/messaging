package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mailgun/mailgun-go"

	"github.com/gorilla/mux"
	"github.com/ramvasanth/wavecell"

	"../models"
	u "../utils"
)

// CreateMessage this func will create a message and save it to scheduler
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var messaging models.MessagingModel

	err := json.NewDecoder(r.Body).Decode(&messaging)
	if err != nil {
		u.RespondMethodNotAllowed(w, u.Message(false, "Error while decoding request body"))
		return
	}

	defer r.Body.Close()

	messaging.Status = 0

	resp := messaging.CreateMessaging()

	if resp["status"] != true {
		u.RespondMethodNotAllowed(w, u.Message(false, "validation error"))
		return
	}
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
			From: os.Getenv("waave_cell_name"),
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

		resp["id"] = id
		resp["data"] = r
		u.Respond(w, resp)
		return
	} else if message.Type == "EMAIL" {
		mg := mailgun.NewMailgun(os.Getenv("mailgun_domain"), os.Getenv("mailgun_private_api"))

		sender := os.Getenv("mailgun_sender_email")
		subject := message.MessageName
		body := message.MessageBody
		recipient := message.DestinationID

		mgMessage := mg.NewMessage(sender, subject, "", recipient)
		mgMessage.SetHtml(body)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		r, id, err := mg.Send(ctx, mgMessage)

		if err != nil {
			log.Fatal(err)
			u.Respond(w, u.Message(false, "Error Sending Message to MailGun"))
			return
		}

		resp := u.Message(true, "success")

		resp["id"] = id
		resp["data"] = r

		u.Respond(w, resp)
		return
	} else if message.Type == "ONEPUSH" {
		// client := onesignal.NewClient(nil)

		// client.UserKey := os.Getenv("ONE_SIGNAL_USER_KEY")
		// client.AppKey := os.Getenv("ONE_SIGNAL_APP_KEY")
	}
}
