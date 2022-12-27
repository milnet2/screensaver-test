package main

import (
	"fmt"
	dbus2 "screensaver-test/service/dbus"
)

func main() {
	var config = configure()

	var dbus = new(dbus2.DBusAdapter)
	dbus.Init(&dbus2.DbusConfig{config.isDbusSystemConnection})

	var dbusChannel = make(chan dbus2.DBusMessage)
	dbus.Listen(dbusChannel)

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
