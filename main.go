package main

import "fmt"

func main() {
	var config = configure()

	var dbus = new(DBusAdapter)
	dbus.Init(&config)

	var dbusChannel = make(chan DBusMessage)
	dbus.Listen(dbusChannel)

	for message := range dbusChannel {
		fmt.Println(message)
	}
}
