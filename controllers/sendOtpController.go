// package controllers

// import (
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/ramvasanth/wavecell"

// 	"../models"
// 	u "../utils"
// 	"github.com/gorilla/mux"
// )

// func SendOtpMessage(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		u.Respond(w, u.Message(false, "There was an error in your request"))
// 	}

// 	message := models.GetMessaging(id)

// 	if message == nil {
// 		u.Respond(w, u.Message(false, "Message not found"))
// 		return
// 	}

// 	if message.Type == "SMS" {

// 		mBody := wavecell.Message{
// 			From: "WL INFO",
// 			To:   message.DestinationID,
// 			Text: message.MessageBody,
// 		}

// 		err := mBody.Validate()

// 		if err != nil {
// 			log.Fatalf("wavecell message error: %v", err)
// 		}
// 		client := wavecell.ClientWithAuthKey("Vfg3DLjMNLZYWUEULxnVHy8X7ybohii24mMiMvWK3E", message.MessageType)

// 		r, err := client.SingleMessage(mBody)
// 		if (err) != nil {
// 			log.Fatalf("wavecell error: %v", err)
// 			return
// 		}

// 		resp := u.Message(true, "success")

// 		resp["data"] = r
// 		u.Respond(w, resp)
// 	} else if (message.Type == "EMAIL") {
// 		mBody :=
// 	}

// 	// resp := u.Message(true, "success")
// 	// resp["data"] = message.Type
// 	// u.Respond(w, resp)
// 	// return

// }
