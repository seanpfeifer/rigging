package logging

// Contains more general utility functions for logging

import "log"

// FatalIfError will log the error and exit if it is non-nil.
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
