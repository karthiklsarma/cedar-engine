package server

import (
	"github.com/karthiklsarma/cedar-engine/m/logging"
)

func InitiateServerEntry() {
	logging.SetInfoLogLevel()
	logging.Info("Initializing server entry")
}
