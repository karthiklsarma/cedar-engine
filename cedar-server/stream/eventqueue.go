package stream

import (
	"context"
	"fmt"
	"os"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
	"github.com/karthiklsarma/cedar-server/m/logging"
)

var hub *eventhub.Hub
var stream_connection_string string

func getConnectionString() string {
	return os.Getenv(STREAM_CONN_ENV)
}

func EmitLocation(location string) (bool, error) {
	if len(location) == 0 {
		logging.Error("invalid input location")
		return false, nil
	}

	logging.Info(fmt.Sprintf("received location: %v", location))
	if len(stream_connection_string) == 0 {
		logging.Info("stream connection string empty. Fetching...")
		stream_connection_string = getConnectionString()
	}

	var err error
	if hub == nil {
		logging.Info("hub empty. Initializing...")
		if hub, err = eventhub.NewHubFromConnectionString(stream_connection_string); err != nil {
			logging.Error(fmt.Sprintf("error initiating eventhub. error: %v", err))
			return false, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err = hub.Send(ctx, eventhub.NewEventFromString(location)); err != nil {
		logging.Error(fmt.Sprintf("Something went wrong while sending msg to eventhub: %v", err))
		return false, err
	}
	logging.Info(fmt.Sprintf("successfully sent message %v to eventhub", location))
	return true, nil
}
