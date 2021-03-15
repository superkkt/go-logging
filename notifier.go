/*
 * Copyright 2021 Samjung Data Service, Inc. All Rights Reserved.
 *
 * Authors:
 * 	Kitae Kim <superkkt@sds.co.kr>
 */

package logging

var defaultNotifier Notifier

type Notifier interface {
	Notify(level Level, msg string)
}

func SetNotifier(notifier Notifier) {
	defaultNotifier = notifier
}
