package dbus

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

type DbusConfig struct {
	IsDbusSystemConnection bool
}

type DBus interface {
	Listen()
	Stop()
}

type DBusMessage struct {
	messageType string
}

// DBusAdapter - Classes are structs in go...
type DBusAdapter struct {
	config *DbusConfig

	dbusInterfaceName string `default:"org.freedesktop.login1.Manager"`

	conn *dbus.Conn
}

func (self *DBusAdapter) Init(config *DbusConfig) {
	self.config = config             // TODO: Null-check
	self.conn = self.connectOrExit() // TODO: Logic in constructor is bad
}

func (self DBusAdapter) connectOrExit() *dbus.Conn {
	var conn *dbus.Conn
	var err error

	if self.config.IsDbusSystemConnection {
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
