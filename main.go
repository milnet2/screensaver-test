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

	dbusSources := [2]dbus2.DBusSource{
		new(dbus2.LoginManager),
		new(dbus2.PowerManagement),
	}

	for _, source := range dbusSources {
		source.Init(dbusAdapter)
		source.Read(dbusChannel)
	}

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
