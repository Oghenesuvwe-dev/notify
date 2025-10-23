package main

import (
	"context"
	"log"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/streamlabs"
)

func main() {
	// Create Streamlabs service with your access token
	streamlabsService := streamlabs.New("your-access-token-here")

	// Use the service with notify
	notify.UseServices(streamlabsService)

	// Send a basic notification (follow alert)
	err := notify.Send(
		context.Background(),
		"New Follower Alert",
		"Thanks for following the stream!",
	)
	if err != nil {
		log.Fatalf("Failed to send notification: %v", err)
	}

	// Send specific alert types
	ctx := context.Background()

	// Donation alert
	err = streamlabsService.SendDonation(ctx, "generous_viewer", "10.00", "Thanks for the amazing donation!")
	if err != nil {
		log.Fatalf("Failed to send donation alert: %v", err)
	}

	// Follow alert
	err = streamlabsService.SendFollow(ctx, "new_follower", "Welcome to the community!")
	if err != nil {
		log.Fatalf("Failed to send follow alert: %v", err)
	}

	// Subscription alert
	err = streamlabsService.SendSubscription(ctx, "loyal_viewer", "Thanks for subscribing!")
	if err != nil {
		log.Fatalf("Failed to send subscription alert: %v", err)
	}

	log.Println("All Streamlabs alerts sent successfully!")
	log.Println("Check your Streamlabs overlay to see the alerts.")
}