package dbus

// ScreenSaver See: https://github.com/freedesktop/xdg-utils/blob/master/scripts/xdg-screensaver.in#L253
type ScreenSaver struct {
	adapter *DBusAdapter

	dbusDestinationObject string `default:"org.freedesktop.ScreenSaver"`
}

func (self *ScreenSaver) Init(adapter *DBusAdapter) {
	self.adapter = adapter
	self.dbusDestinationObject = "org.freedesktop.ScreenSaver"
}

func (self ScreenSaver) Read(listener chan DBusMessage) {
	// TODO: Implement
}

func (self *ScreenSaver) IsApplicable() bool {
	return self.adapter.isObjectPresent(self.dbusDestinationObject)
}
