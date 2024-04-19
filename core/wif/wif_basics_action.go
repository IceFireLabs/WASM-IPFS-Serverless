package wif

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func (WIS *WIS) LoadBasics() (err error) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.Infoln(WIS.Confer.Opts.ApptypeConf, "Basics load success. ")
	return nil
}
