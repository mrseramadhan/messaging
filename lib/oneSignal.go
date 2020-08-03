package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// PayloadOneSignal represents a request to run a command.
type PayloadOneSignal struct {
	AppID            string            `json:"app_id"`
	Contents         ContentsOneSignal `json:"contents"`
	Headings         HeadingsOneSignal `json:"headings"`
	IncludePlayerIds []string          `json:"include_player_ids"`
}

// ContentsOneSignal represents a request to run a command.
type ContentsOneSignal struct {
	En string `json:"en"`
}

// HeadingsOneSignal represents a request to run a command.
type HeadingsOneSignal struct {
	En string `json:"en"`
}

func oneSignalSingle() {

	serverURI := os.Getenv("one_signal_server")
	serverKey := os.Getenv("one_signal_lkey")

	data := PayloadOneSignal{
		// fill struct
	}

	data.AppID = serverKey
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", serverURI, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
}
