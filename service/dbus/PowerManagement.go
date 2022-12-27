package dbus

type PowerManagement struct {
	adapter *DBusAdapter

	dbusDestinationObject string `default:"org.freedesktop.PowerManagement"`
}

func (self *PowerManagement) Init(adapter *DBusAdapter) {
	self.adapter = adapter
	self.dbusDestinationObject = "org.freedesktop.PowerManagement"
}

func (self PowerManagement) Read(listener chan DBusMessage) {
	self.readHasInhibit(listener)
}

func (self *PowerManagement) IsApplicable() bool {
	return self.adapter.isObjectPresent(self.dbusDestinationObject)
}

func (self PowerManagement) readHasInhibit(listener chan DBusMessage) {
	// Query legacy PM:
	// like `dbus-send --system --print-reply --dest=org.freedesktop.PowerManagement /org/freedesktop/PowerManagement/Inhibit org.freedesktop.PowerManagement.Inhibit.HasInhibit`
	self.adapter.readCurrentPropertiesFrom(
		listener,
		self.dbusDestinationObject,
		"org.freedesktop.PowerManagement.Inhibit.HasInhibit", // returns boolean
		"/org/freedesktop/PowerManagement/Inhibit")
}
