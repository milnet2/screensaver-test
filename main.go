package main

import (
	"fmt"
	dbus2 "screensaver-test/service/dbus"
)

func main() {
	var config = configure()

	var dbusAdapter = new(dbus2.DBusAdapter)
	dbusAdapter.Init(&dbus2.DbusConfig{IsDbusSystemConnection: config.isDbusSystemConnection})
	dbusAdapter.ListNames()

	var dbusChannel = make(chan dbus2.DBusMessage)

	dbusSources := []dbus2.DBusSource{
		new(dbus2.LoginManager),
		new(dbus2.PowerManagement),
		new(dbus2.SystemD),
	}

	for _, source := range dbusSources {
		source.Init(dbusAdapter)
		if source.IsApplicable() {
			source.Read(dbusChannel)
		}
	}

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
