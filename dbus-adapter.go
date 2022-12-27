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

	dbusInterface string `default:"org.freedesktop.login1.Manager"`

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
	//defer conn.Close() // TODO at some point close the connection

	self.conn = conn
	return conn
}

// See: https://github.com/godbus/dbus/tree/master/_examples

func (self DBusAdapter) Listen(listener chan DBusMessage) {
	self.readCurrentProperties(listener)

	// DBUs:
	// method call time=1672161608.325197 sender=:1.65 -> destination=org.freedesktop.login1 serial=2 path=/org/freedesktop/login1; interface=org.freedesktop.login1.Manager; member=Inhibit
	//   string "idle:sleep:shutdown"

	//for _, v := range []string{"Inhibit"} {
	//	call := self.conn.BusObject().Call("org.freedesktop.login1.Manager", 0,
	//		"eavesdrop='true',type='"+v+"'")
	//
	//	if call.Err != nil {
	//		fmt.Fprintln(os.Stderr, "Failed to add match:", call.Err)
	//		os.Exit(1)
	//	} else {
	//		listener <- DBusMessage{
	//			messageType: call.Method,
	//		}
	//	}
	//}

	//c := make(chan *dbus.Message, 10)
	//self.conn.Eavesdrop(c)
	//fmt.Println("Listening for everything")
	//for v := range c {
	//	fmt.Println(v)
	//}
}

func (self DBusAdapter) readCurrentProperties(listener chan DBusMessage) {
	//    string "org.freedesktop.login1.Manager"
	//   array [
	//      dict entry(
	//         string "BlockInhibited"
	//         variant             string "shutdown:sleep:idle"
	//      )
	//   ]
	//   array [
	//   ]

	var variant, err = self.conn.BusObject().GetProperty("org.freedesktop.login1.Manager.BlockInhibited")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	} else {
		listener <- DBusMessage{
			messageType: variant.String(),
		}
	}
}
