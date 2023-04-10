/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	errorsutil "k8s.io/apimachinery/pkg/util/errors"
)

const (
	// DefaultErrorExitCode defines exit the code for failed action generally
	DefaultErrorExitCode = 1
	// PreFlightExitCode defines exit the code for preflight checks
	PreFlightExitCode = 2
	// ValidationExitCode defines the exit code validation checks
	ValidationExitCode = 3
)

type preflightError interface {
	Preflight() bool
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

func CheckErr(err error) {
	checkErr(err, fatal)
}

func checkErr(err error, handleErr func(string, int)) {

	var msg string
	if err != nil {
		msg = fmt.Sprintf("%s\nTo see the stack trace of this error execute with --v=5 or higher", err.Error())
		// check if the verbosity level in klog is high enough and print a stack trace.
		f := flag.CommandLine.Lookup("v")
		if f != nil {
			// assume that the "v" flag contains a parseable Int32 as per klog's "Level" type alias,
			// thus no error from ParseInt is handled here.
			if v, e := strconv.ParseInt(f.Value.String(), 10, 32); e == nil {
				// https://git.k8s.io/community/contributors/devel/sig-instrumentation/logging.md
				// klog.V(5) - Trace level verbosity
				if v > 4 {
					msg = fmt.Sprintf("%+v", err)
				}
			}
		}
	}

	switch err.(type) {
	case nil:
		return
	case preflightError:
		handleErr(msg, PreFlightExitCode)
	case errorsutil.Aggregate:
		handleErr(msg, ValidationExitCode)

	default:
		handleErr(msg, DefaultErrorExitCode)
	}
}
