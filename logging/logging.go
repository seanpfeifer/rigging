// Package logging contains very general utility functions for logging.
// These are meant to be useful only in very basic situations, and they use the standard "log" lib.
//
// Examples of use cases for this:
//   - Small scripts without the need for complex logging (eg, Advent of Code)
//   - Prior to setting up your actual log system (eg, failure to set up remote logging)
package logging

import "log"

// FatalIfError will log the error and exit if it is non-nil.
// This is useful in particular for non-recoverable errors when starting an application.
// Note that calls to `defer` will not be triggered by this - no cleanup is done!
func FatalIfError(err error, extraInfo ...any) {
	if err != nil {
		log.Fatal(err, extraInfo)
	}
}

// LogIfError will log the [error + extra info] if the error is non-nil.
// Returns true if err is non-nil.
func LogIfError(err error, extraInfo ...any) bool {
	if err != nil {
		log.Println(err, extraInfo)
		return true
	}
	return false
}
