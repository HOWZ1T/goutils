package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

/*
	log levels based on a value system.
	if a given level is <= a level (e.g: debug) it will print out that stream of logs
	e.g: a level of ERROR will print out both ERROR & CRITICAL log streams
*/
type Level int

const (
	FATAL   Level = 40
	ERROR   Level = 30
	WARNING Level = 20
	INFO    Level = 10
	DEBUG   Level = 0
)

// simple object to represent a log for a given system / subsystem
type logger struct {
	lvl       Level
	fmt       string
	ioStreams []io.Writer
}

var log logger

func init() {
	log = logger{INFO, "%-30s | %-20s | %-10s | ", []io.Writer{os.Stdout}}
}

func getCaller() string {
	caller := "UNKNOWN"
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get the caller
	n := runtime.Callers(3, fpcs)
	if n != 0 {
		c := runtime.FuncForPC(fpcs[0] - 1)
		if c == nil {
			caller = "NIL"
		} else {
			caller = c.Name()
		}
	}
	return caller
}

/// ------ GETTERS ------ ///

func GetLevel() Level {
	return log.lvl
}

func GetFMT() string {
	return log.fmt
}

func GetIoStreams() []io.Writer {
	return log.ioStreams
}

/// ------ SETTERS ------ ///

func SetLevel(level Level) {
	log.lvl = level
}

// WARNING: this function is destructive and will reset the io streams for the logger !
func SetIOStreams(streams []io.Writer) {
	log.ioStreams = streams
}

func SetFMT(fmt string) {
	log.fmt = fmt
}

func AddStream(stream io.Writer) {
	// check if stream already exists
	for _, lStream := range log.ioStreams {
		// stream exists, return here as it is 'already added'
		if lStream == stream {
			return
		}
	}

	// append a new stream to the logger
	log.ioStreams = append(log.ioStreams, stream)
}

// removes the stream by reconstructing the stream slice without the stream to be removed
func RemoveStream(stream io.Writer) error {
	var newStreams []io.Writer
	for _, lStream := range log.ioStreams {
		if lStream != stream {
			newStreams = append(newStreams, lStream)
		}
	}

	if len(newStreams) == 0 {
		return errors.New("could not remove stream as it'd leave the logger empty")
	}

	log.ioStreams = newStreams
	return nil
}

// removes the streams (and the name & level) of the logger thereby disabling it, does not delete the actual logger!
func Delete() {
	log.lvl = 0
	log.fmt = ""
	log.ioStreams = nil
}

/*
	private log function.
	forms basis for all other level of logs.
*/
func writeLog(time, caller, level, format string, a ...interface{}) (n int, err error) {
	n = 0
	for _, stream := range log.ioStreams {
		a = append([]interface{}{time, caller, level}, a...) // pre-pending log data
		bytes, e := fmt.Fprintf(stream, log.fmt+format+"\n", a...)
		n += bytes
		if e != nil {
			return n, e
		}
	}

	return n, nil
}

// handles logging critical level messages
func Fatal(format string, a ...interface{}) (n int, err error) {
	if log.lvl <= FATAL {
		caller := getCaller()
		t := time.Now()
		return writeLog(t.Format(time.RFC850), caller, "FATAL", format, a...)
	}

	return 0, nil
}

// handles logging error level messages
func Error(format string, a ...interface{}) (n int, err error) {
	if log.lvl <= ERROR {
		caller := getCaller()
		t := time.Now()
		return writeLog(t.Format(time.RFC850), caller, "ERROR", format, a...)
	}

	return 0, nil
}

// handles logging warning level messages
func Warning(format string, a ...interface{}) (n int, err error) {
	if log.lvl <= WARNING {
		caller := getCaller()
		t := time.Now()
		return writeLog(t.Format(time.RFC850), caller, "WARNING", format, a...)
	}

	return 0, nil
}

// handles logging info level messages
func Info(format string, a ...interface{}) (n int, err error) {
	if log.lvl <= INFO {
		caller := getCaller()
		t := time.Now()
		return writeLog(t.Format(time.RFC850), caller, "INFO", format, a...)
	}

	return 0, nil
}

// handles logging debug level messages
func Debug(format string, a ...interface{}) (n int, err error) {
	if log.lvl <= DEBUG {
		caller := getCaller()
		t := time.Now()
		return writeLog(t.Format(time.RFC850), caller, "DEBUG", format, a...)
	}

	return 0, nil
}
