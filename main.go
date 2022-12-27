package main

import (
	"fmt"
	dbus2 "screensaver-test/service/dbus"
	"screensaver-test/service/x11"
)

func main() {
	var config = configure()

	// DBUS -----------------------------------

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

	// X11 ------------------------------------

	var x11Adapter = new(x11.X11Connection)
	x11Adapter.Init()

	x11Sources := []x11.X11Source{
		new(x11.DPMS),
		new(x11.ProtoScreensaver),
		new(x11.MitScreensaver),
	}

	for _, source := range x11Sources {
		source.Init(x11Adapter)
		if source.IsApplicable() {
			source.Read()
		}
	}

	// OUT ------------------------------------

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
