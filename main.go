package main

import "fmt"

func main() {
	var config = configure()

	var dbus = new(DBusAdapter)
	dbus.Init(&config)

	fmt.Println(config)
	fmt.Println(&dbus)
}
