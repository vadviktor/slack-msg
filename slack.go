// Package slack_msg provides a very basic and simple mean to send messages
// to a Slack webhook.
package slack_msg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Slack struct {
	client     *http.Client
	webhookURL string
}

// Create initiates the Slack messenger using all the vital information provided.
func (s *Slack) Create(webhookURL string) {
	s.webhookURL = webhookURL
	s.client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

// Send uses the Stringer interface to compose the messages sent over to Slack.
func (s *Slack) Send(text string, params ...interface{}) {
	t := map[string]string{"text": fmt.Sprintf(text, params...)}
	payload, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Failed to create json payload for Slack: %s\n",
			err.Error())
	}

	p := strings.NewReader(string(payload))
	resp, err := s.client.Post(s.webhookURL, "application/json", p)
	if err != nil {
		log.Fatalf("Failed to pass text to Slack: %s\n", err.Error())
	}
	defer resp.Body.Close()
}
