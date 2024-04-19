package openlog

import (
	"log"
)

const (
	i = "INFO: "
	d = "DEBUG: "
	e = "ERROR: "
	w = "WARN: "
	f = "FATAL: "
)

type golog struct {
}

func (l golog) Debug(message string, opts ...Option) {
	log.Println(d + message)
}

func (l golog) Info(message string, opts ...Option) {
	log.Println(i + message)
}
func (l golog) Warn(message string, opts ...Option) {
	log.Println(w + message)
}
func (l golog) Error(message string, opts ...Option) {
	log.Println(e + message)
}
func (l golog) Fatal(message string, opts ...Option) {
	log.Panic(f + message)
}
