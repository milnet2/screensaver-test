package main

import (
	"fmt"
	dbus2 "screensaver-test/service/dbus"
)

func main() {
	var config = configure()

	var dbusAdapter = new(dbus2.DBusAdapter)
	dbusAdapter.Init(&dbus2.DbusConfig{config.isDbusSystemConnection})

	var dbusChannel = make(chan dbus2.DBusMessage)

	var loginManagerDbus = new(dbus2.LoginManager)
	loginManagerDbus.Init(dbusAdapter)

	dbusAdapter.Listen(dbusChannel)
	loginManagerDbus.Read(dbusChannel)

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
