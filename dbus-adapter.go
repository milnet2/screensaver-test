package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

type DBus interface {
	Listen()
	Stop()
}

type DBusMessage struct {
	messageType string
}

// DBusAdapter - Classes are structs in go...
type DBusAdapter struct {
	config *config

	conn *dbus.Conn
}

func (self *DBusAdapter) Init(config *config) {
	self.config = config             // TODO: Null-check
	self.conn = self.connectOrExit() // TODO: Logic in constructor is bad
}

func (self DBusAdapter) connectOrExit() *dbus.Conn {
	var conn *dbus.Conn
	var err error

	if self.config.isDbusSystemConnection {
		conn, err = dbus.ConnectSystemBus()

	} else {
		conn, err = dbus.ConnectSessionBus()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	self.conn = conn
	return conn
}

func (self DBusAdapter) Listen() {

	for _, v := range []string{"method_call", "method_return", "error", "signal"} {
		call := self.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
			"eavesdrop='true',type='"+v+"'")
		if call.Err != nil {
			fmt.Fprintln(os.Stderr, "Failed to add match:", call.Err)
			os.Exit(1)
		}
	}
	c := make(chan *dbus.Message, 10)
	self.conn.Eavesdrop(c)
	fmt.Println("Listening for everything")
	for v := range c {
		fmt.Println(v)
	}
}
