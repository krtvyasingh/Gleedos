package main

import "github.com/krtvysingh/gleedos/pkg/notify"

func notifyCompletion(title, file string) {
	_ = notify.SendNotification(title, "Downloaded: "+file)
}
