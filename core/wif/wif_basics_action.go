package wif

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// LoadBasics initializes the logging configuration and logs a success message.
func (WIS *WIS) LoadBasics() (err error) {
	// Set the log output to standard output (stdout)
	log.SetOutput(os.Stdout)

	// Set the log level to InfoLevel
	log.SetLevel(log.InfoLevel)

	// Configure the log formatter to use plain text with full timestamps and no colors
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true, // Disable colored output
		FullTimestamp: true, // Include full timestamp in the log entries
	})

	// Log a success message indicating that the basics have been loaded
	log.Infoln(WIS.Confer.Opts.ApptypeConf, "Basics load success.")

	return nil
}
