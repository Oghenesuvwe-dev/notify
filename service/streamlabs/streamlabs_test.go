package streamlabs

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	service := New("test-token")
	assert.NotNil(t, service)
	assert.Equal(t, "test-token", service.accessToken)
	assert.NotNil(t, service.client)
}

func TestSendNoToken(t *testing.T) {
	service := New("")

	err := service.Send(context.Background(), "Test", "Message")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "access token is required")
}

func TestSendSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))

		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "test-token", r.FormValue("access_token"))
		assert.Equal(t, "follow", r.FormValue("type"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := New("test-token")
	service.apiURL = server.URL

	err := service.SendAlert(context.Background(), "follow", "Test message", "testuser", "")
	assert.NoError(t, err)
}

func TestSendDonation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "donation", r.FormValue("type"))
		assert.Equal(t, "testuser", r.FormValue("user_name"))
		assert.Equal(t, "10.00", r.FormValue("amount"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := New("test-token")
	service.apiURL = server.URL

	err := service.SendDonation(context.Background(), "testuser", "10.00", "Thanks for the donation!")
	assert.NoError(t, err)
}

func TestSendFollow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "follow", r.FormValue("type"))
		assert.Equal(t, "newfollower", r.FormValue("user_name"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := New("test-token")
	service.apiURL = server.URL

	err := service.SendFollow(context.Background(), "newfollower", "Thanks for following!")
	assert.NoError(t, err)
}

func TestSendSubscription(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "subscription", r.FormValue("type"))

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	service := New("test-token")
	service.apiURL = server.URL

	err := service.SendSubscription(context.Background(), "subscriber", "Thanks for subscribing!")
	assert.NoError(t, err)
}

func TestSendFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	service := New("test-token")
	service.apiURL = server.URL

	err := service.SendAlert(context.Background(), "follow", "Test", "user", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "streamlabs API returned status 401")
}
