package logging

import (
	"log"
)

// FATAL - Syntatic sugar for logging
func FATAL(message string) {
	log.Fatalf(message)
}

// INFO - Syntatic sugar for logging
func INFO(message string) {
	log.Printf("INFO %s", message)
}

// DEBUG - Syntatic sugar for logging
func DEBUG(message string) {
	log.Printf("DEBUG %s", message)
}
