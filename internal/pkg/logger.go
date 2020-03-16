package logger

import (
	log "github.com/sirupsen/logrus"
)

// Init will initial logger for application
func Init() {
	log.SetFormatter(&log.JSONFormatter{})
}
