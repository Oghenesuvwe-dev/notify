/*
Package streamlabs provides message notification integration for Streamlabs (https://streamlabs.com).

Streamlabs is a live streaming software that allows content creators to manage
donations, alerts, and viewer engagement. This service enables sending custom
alerts to your Streamlabs overlay.

Usage:

	package main

	import (
		"context"
		"log"

		"github.com/nikoksr/notify"
		"github.com/nikoksr/notify/service/streamlabs"
	)

	func main() {
		streamlabsService := streamlabs.New("your-access-token")

		notify.UseServices(streamlabsService)

		err := notify.Send(context.Background(), "New Follower", "Thanks for following!")
		if err != nil {
			log.Fatal(err)
		}
	}
*/
package streamlabs
