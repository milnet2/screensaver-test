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

	dbusInterfaceName string `default:"org.freedesktop.login1.Manager"`

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
	self.readCurrentInhibitors(listener)
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

func (self DBusAdapter) readCurrentInhibitors(listener chan DBusMessage) {
	self.readCurrentInhibitorsFrom(
		listener,
		"org.freedesktop.login1",
		"org.freedesktop.login1.Manager.ListInhibitors",
		"/org/freedesktop/login1")
}

func (self DBusAdapter) readCurrentInhibitorsFrom(listener chan DBusMessage, dbusDestObject string, dbusMethod string, dbusPath dbus.ObjectPath) {
	var dbusDestination = self.conn.Object(dbusDestObject, dbusPath)

	var response string                                         // TODO: Fix type!
	err := dbusDestination.Call(dbusMethod, 0).Store(&response) // "eavesdrop='true',type='"+v+"'")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to add match:", err)
		os.Exit(1)
	}

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

	//    bus.add_match_string_non_blocking("eavesdrop=true, path='/org/freedesktop/PowerManagement/Inhibit', interface='org.freedesktop.PowerManagement.Inhibit'")
	//
	//dbus-send --print-reply \
	//      --type=method_call \
	//      --system \
	//      /org/freedesktop/login1  \
	//      org.freedesktop.DBus.Properties.Get \
	//      string:org.freedesktop.NetworkManager.Device \
	//      string:Interface
	//
	//const dbusInterface = "org.freedesktop.DBus.Properties"
	//const dbusPath = "/org/freedesktop/login1"
	//const dbusObject = "org.freedesktop.login1.Manager"
	//const dbusPropertyKey = "BlockInhibited"

	// Query login-manager:
	// like `dbus-send --print-reply --dest=org.freedesktop.login1.Manager /org/freedesktop/PowerManagement/Inhibit org.freedesktop.PowerManagement.Inhibit.HasInhibit`
	//dbus-send --system --print-reply \
	//--dest=org.freedesktop.login1 /org/freedesktop/login1/session/self \
	// "org.freedesktop.login1.Session.SetIdleHint" boolean:true
	// See: `gdbus introspect -y -d org.freedesktop.login1 -o /org/freedesktop/login1/session/auto`
	self.readCurrentPropertiesFrom(
		listener,
		"org.freedesktop.login1",
		"org.freedesktop.login1.Session.IdleHint",
		"/org/freedesktop/login1/session/self")

	// Query legacy PM:
	// like `dbus-send --system --print-reply --dest=org.freedesktop.PowerManagement /org/freedesktop/PowerManagement/Inhibit org.freedesktop.PowerManagement.Inhibit.HasInhibit`
	self.readCurrentPropertiesFrom(
		listener,
		"org.freedesktop.PowerManagement",
		"org.freedesktop.PowerManagement.Inhibit.HasInhibit", // returns boolean
		"/org/freedesktop/PowerManagement/Inhibit")
}

func (self DBusAdapter) readCurrentPropertiesFrom(listener chan DBusMessage, dbusDestObject string, dbusInterface string, dbusPath dbus.ObjectPath) {
	var dbusPropertiesObject = self.conn.Object(dbusDestObject, dbusPath)

	// GetProperty calls org.freedesktop.DBus.Properties.Get on the given object
	var variant, err = dbusPropertiesObject.GetProperty(dbusInterface)
	//var variant, err = self.conn.BusObject().GetProperty(dbusPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "Read", dbusInterface, " = ", variant)
		//listener <- DBusMessage{
		//	messageType: variant.String(),
		//}
	}
}
