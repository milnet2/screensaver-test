package dbus

// SystemD interface to gdbus introspect --system --dest org.freedesktop.systemd1 --object-path /org/freedesktop/systemd1
type SystemD struct {
	adapter *DBusAdapter

	dbusDestinationObject string `default:"org.freedesktop.systemd1"`
	dbusDestinationPath   string `default:"/org/freedesktop/systemd1"`
}

func (self *SystemD) Init(adapter *DBusAdapter) {
	self.adapter = adapter
	self.dbusDestinationObject = "org.freedesktop.SystemD"
	self.dbusDestinationPath = "/org/freedesktop/systemd1"
}

func (self SystemD) Read(listener chan DBusMessage) {
	// TODO
}

func (self *SystemD) IsApplicable() bool {
	return self.adapter.isObjectPresent(self.dbusDestinationObject)
}
