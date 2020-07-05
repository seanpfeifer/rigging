// Package logging contains very general utility functions for logging.
// These are meant to useful only in very basic situations, and use the standard "log" lib.
package logging

import "log"

// FatalIfError will log the error and exit if it is non-nil.
// This is useful in particular for non-recoverable errors when starting an application.
// Note that calls to `defer` will not be triggered by this - no cleanup is done!
func FatalIfError(err error, extraInfo ...interface{}) {
	if err != nil {
		log.Fatal(err, extraInfo)
	}
}

// LogIfError will log the [error + extra info] if the error is non-nil.
// Returns true if err is non-nil.
func LogIfError(err error, extraInfo ...interface{}) bool {
	if err != nil {
		log.Println(err, extraInfo)
		return true
	}
	return false
}
