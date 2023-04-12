package util

import (
	"fmt"
	"os"
	"strings"
)

type MyError struct {
	error
	message string
	code    int
}

func Trap(err error) {
	trap(err, fatal)
}

func fatal(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

func trap(err error, errorHandler func(string, int)) {
	var msg string

	if err == nil {
		return
	}

	switch v := err.(type) {
	case MyError:
		errorHandler(v.message, v.code)
	default:
		errorHandler(msg, 1)
	}
}
