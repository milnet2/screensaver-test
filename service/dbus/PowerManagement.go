package dbus

type PowerManagement struct {
	adapter *DBusAdapter
}

func (self *PowerManagement) Init(adapter *DBusAdapter) {
	self.adapter = adapter
}

func (self PowerManagement) Read(listener chan DBusMessage) {
	self.readHasInhibit(listener)
}

func (self *PowerManagement) IsApplicable() bool {
	return self.adapter.isObjectPresent("org.freedesktop.PowerManagement")
}

func (self PowerManagement) readHasInhibit(listener chan DBusMessage) {
	// Query legacy PM:
	// like `dbus-send --system --print-reply --dest=org.freedesktop.PowerManagement /org/freedesktop/PowerManagement/Inhibit org.freedesktop.PowerManagement.Inhibit.HasInhibit`
	self.adapter.readCurrentPropertiesFrom(
		listener,
		"org.freedesktop.PowerManagement",
		"org.freedesktop.PowerManagement.Inhibit.HasInhibit", // returns boolean
		"/org/freedesktop/PowerManagement/Inhibit")
}
