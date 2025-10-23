# Streamlabs

## Prerequisites

To use the Streamlabs service, you need:

1. A Streamlabs account
2. An access token from Streamlabs API

## Usage

### Basic Usage
```go
package main

import (
    "context"
    "log"
    
    "github.com/nikoksr/notify"
    "github.com/nikoksr/notify/service/streamlabs"
)

func main() {
    // Create Streamlabs service
    streamlabsService := streamlabs.New("your-access-token")
    
    // Add service to notify
    notify.UseServices(streamlabsService)
    
    // Send notification (defaults to follow alert)
    err := notify.Send(
        context.Background(),
        "New Follower",
        "Thanks for following the stream!",
    )
    if err != nil {
        log.Fatal(err)
    }
}
```

### Specific Alert Types

#### Donation Alert
```go
err := streamlabsService.SendDonation(ctx, "username", "5.00", "Thanks for the donation!")
```

#### Follow Alert
```go
err := streamlabsService.SendFollow(ctx, "newfollower", "Welcome to the stream!")
```

#### Subscription Alert
```go
err := streamlabsService.SendSubscription(ctx, "subscriber", "Thanks for subscribing!")
```

#### Custom Alert
```go
err := streamlabsService.SendAlert(ctx, "custom", "Custom message", "username", "amount")
```

## Getting Access Token

1. Login to your Streamlabs account
2. Go to Settings → API Settings
3. Generate or copy your access token
4. Use the token in your application

## Alert Types

The service supports various alert types:
- `donation` - For donation notifications
- `follow` - For new follower alerts
- `subscription` - For subscription alerts
- `host` - For host notifications
- `raid` - For raid alerts
- `custom` - For custom alerts

## Features

- ✅ Multiple alert types support
- ✅ Custom message content
- ✅ User name and amount parameters
- ✅ Access token authentication
- ✅ Context cancellation
- ✅ Comprehensive error handling
- ✅ Full test coverage

## Links

- [Streamlabs](https://streamlabs.com/)
- [Streamlabs API Documentation](https://dev.streamlabs.com/)
- [Alert Box Setup](https://streamlabs.com/content-hub/post/alert-box-setup)