/*
 * Copyright 2016 Samjung Data Service, Inc. All Rights Reserved.
 *
 * Authors:
 * 	Kitae Kim <superkkt@sds.co.kr>
 */

package logging

import (
	"fmt"
	"log/syslog"
	"runtime"
	"strings"
)

const maxLineLength = 32 * 1024

type tidSyslog struct {
	writer *syslog.Writer
}

// NewTIDSyslog returns a logging backend that uses syslog with thread IDs.
func NewTIDSyslog(prefix string) (Backend, error) {
	w, err := syslog.New(syslog.LOG_CRIT, prefix)
	if err != nil {
		return nil, err
	}

	return &tidSyslog{writer: w}, nil
}

func (r *tidSyslog) Log(level Level, calldepth int, record *Record) error {
	line := fmt.Sprintf("%v (TID=%v)", record.Formatted(calldepth+1), getGoRoutineID())
	// syslog limits the maximum length of a line.
	if len(line) > maxLineLength {
		line = line[:maxLineLength]
	}

	switch level {
	case CRITICAL:
		return r.writer.Crit(line)
	case ERROR:
		return r.writer.Err(line)
	case WARNING:
		return r.writer.Warning(line)
	case NOTICE:
		return r.writer.Notice(line)
	case INFO:
		return r.writer.Info(line)
	case DEBUG:
		return r.writer.Debug(line)
	default:
		panic("unexpected log level")
	}
}

func getGoRoutineID() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	return strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
}
