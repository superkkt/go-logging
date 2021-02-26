/*
 * Copyright 2021 Samjung Data Service, Inc. All Rights Reserved.
 *
 * Authors:
 * 	Kitae Kim <superkkt@sds.co.kr>
 */

package logging

var defaultNotifier Notifier
var defaultNotifyLevel Level

type Notifier interface {
	Send(title, body string) error
}

func SetNotifier(notifier Notifier, level Level) {
	defaultNotifier = notifier
	defaultNotifyLevel = level
}
