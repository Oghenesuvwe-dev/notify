package streamlabs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Streamlabs struct holds necessary data to communicate with Streamlabs API.
type Streamlabs struct {
	accessToken string
	client      *http.Client
	apiURL      string
}

// AlertData represents the Streamlabs alert payload.
type AlertData struct {
	Type     string `json:"type"`
	Message  string `json:"message,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Amount   string `json:"amount,omitempty"`
}

// New returns a new instance of Streamlabs notification service.
func New(accessToken string) *Streamlabs {
	return &Streamlabs{
		accessToken: accessToken,
		client:      &http.Client{},
		apiURL:      "https://streamlabs.com/api/v1.0/alerts",
	}
}

// Send takes a message subject and a message body and sends them as Streamlabs alert.
func (s *Streamlabs) Send(ctx context.Context, subject, message string) error {
	if s.accessToken == "" {
		return errors.New("access token is required")
	}

	return s.SendAlert(ctx, "follow", subject, message, "")
}

// SendAlert sends a custom alert to Streamlabs.
func (s *Streamlabs) SendAlert(ctx context.Context, alertType, message, userName, amount string) error {
	if s.accessToken == "" {
		return errors.New("access token is required")
	}

	data := url.Values{}
	data.Set("access_token", s.accessToken)
	data.Set("type", alertType)
	data.Set("message", message)

	if userName != "" {
		data.Set("user_name", userName)
	}
	if amount != "" {
		data.Set("amount", amount)
	}

	apiURL := s.apiURL
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("streamlabs API returned status %d", resp.StatusCode)
	}

	return nil
}

// SendDonation sends a donation alert.
func (s *Streamlabs) SendDonation(ctx context.Context, userName, amount, message string) error {
	return s.SendAlert(ctx, "donation", message, userName, amount)
}

// SendFollow sends a follow alert.
func (s *Streamlabs) SendFollow(ctx context.Context, userName, message string) error {
	return s.SendAlert(ctx, "follow", message, userName, "")
}

// SendSubscription sends a subscription alert.
func (s *Streamlabs) SendSubscription(ctx context.Context, userName, message string) error {
	return s.SendAlert(ctx, "subscription", message, userName, "")
}
