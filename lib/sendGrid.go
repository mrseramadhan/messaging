package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PayloadSendGrid represents a request to run a command.
type PayloadSendGrid struct {
	Personalizations []PersonalizationsSendGrid `json:"personalizations"`
	Content          []ContentSendGrid          `json:"content"`
	From             FromSendGrid               `json:"from"`
	ReplyTo          ReplyToSendGrid            `json:"reply_to"`
}

// ToSendGrid represents a request to run a command.
type ToSendGrid struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// PersonalizationsSendGrid represents a request to run a command.
type PersonalizationsSendGrid struct {
	To      []ToSendGrid `json:"to"`
	Subject string       `json:"subject"`
}

// ContentSendGrid represents a request to run a command.
type ContentSendGrid struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// FromSendGrid represents a request to run a command.
type FromSendGrid struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ReplyToSendGrid represents a request to run a command.
type ReplyToSendGrid struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// SendEmail represents a request to run a command.
func SendEmail() {
	data := PayloadSendGrid{
		// fill struct
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
