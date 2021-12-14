package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/drepione/onesignal-go"
	"github.com/mailgun/mailgun-go"

	"github.com/gorilla/mux"
	"github.com/ramvasanth/wavecell"

	"../models"
	u "../utils"
)

type PushNotif struct {
	URL_PATH    string
	PREFIX_PATH string
}

// CreateMessage this func will create a message and save it to scheduler
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	var messaging models.MessagingModel

	var scheduling models.ScheduleModel

	err := json.NewDecoder(r.Body).Decode(&messaging)
	if err != nil {
		u.RespondMethodNotAllowed(w, u.Message(false, "Error while decoding request body"))
		return
	}

	defer r.Body.Close()

	messaging.ScheduleDate = time.Now()
	messaging.Status = 0

	resp := messaging.CreateMessaging()

	scheduling.Status = 0
	scheduling.Counterfailed = 0
	scheduling.Type = 0
	scheduling.Name = messaging.MessageName
	scheduling.URLPATH = os.Getenv("base_url") + "/messaging/" + strconv.FormatUint(uint64(messaging.ID), 10)
	scheduling.StartDate = messaging.ScheduleDate
	scheduling.CreateSchedule()

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
		u.RespondNotFound(w, u.Message(false, "Failed to find messaging"))
	}

	message := models.GetMessaging(id)

	if message == nil {
		u.RespondNotFound(w, u.Message(false, "Message Not found"))
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
			u.RespondBadRequest(w, u.Message(false, "Error Sending Message to MailGun"))
			return
		}

		resp := u.Message(true, "success")

		resp["id"] = id
		resp["data"] = r

		u.Respond(w, resp)
		return
	} else if message.Type == "PUSH_NOTIF" {
		var pushNotif PushNotif
		var urlPath = ""
		if message.MessageDesc != nil {
			json.Unmarshal([]byte(*message.MessageDesc), &pushNotif)

			if pushNotif.URL_PATH != "" {
				urlPath = pushNotif.URL_PATH
			} else if pushNotif.PREFIX_PATH != "" {
				urlPath = pushNotif.PREFIX_PATH
			} else {
				urlPath = ""
			}
		}

		client := onesignal.NewClient(nil)

		client.UserKey = os.Getenv("ONE_SIGNAL_USER_KEY")
		client.AppKey = os.Getenv("ONE_SIGNAL_APP_KEY")

		appID := os.Getenv("ONE_SIGNAL_APP_ID")

		notificationReq := &onesignal.NotificationRequest{
			AppID:                  appID,
			Contents:               map[string]string{"en": message.MessageBody},
			URL:                    urlPath,
			IncludeExternalUserIDs: []string{message.DestinationID},
		}

		createRes, res, err := client.Notifications.Create(notificationReq)

		if err != nil {
			u.RespondBadRequest(w, u.Message(false, "Error Sending Message to Onepush"))
			return
		}

		resp := u.Message(true, "Success")
		resp["data"] = createRes
		resp["res"] = res

		u.Respond(w, resp)
		return

	}
}
