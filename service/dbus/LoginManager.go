package dbus

type LoginManager struct {
	adapter *DBusAdapter

	dbusDestinationObject string `default:"org.freedesktop.login1"`
}

func (self *LoginManager) Init(adapter *DBusAdapter) {
	self.adapter = adapter
	self.dbusDestinationObject = "org.freedesktop.login1"
}

func (self LoginManager) Read(listener chan DBusMessage) {
	self.readIdleHint(listener)
	self.readBlockInhibited(listener)
	self.readCurrentInhibitors(listener)
}

func (self LoginManager) IsApplicable() bool {
	print(self.dbusDestinationObject)
	return self.adapter.isObjectPresent(self.dbusDestinationObject)
}

func (self LoginManager) readIdleHint(listener chan DBusMessage) {
	// Query login-manager:
	// like `dbus-send --print-reply --dest=org.freedesktop.login1.Manager /org/freedesktop/PowerManagement/Inhibit org.freedesktop.PowerManagement.Inhibit.HasInhibit`
	//dbus-send --system --print-reply \
	//--dest=org.freedesktop.login1 /org/freedesktop/login1/session/self \
	// "org.freedesktop.login1.Session.SetIdleHint" boolean:true
	// See: `gdbus introspect -y -d org.freedesktop.login1 -o /org/freedesktop/login1/session/auto`
	self.adapter.readCurrentPropertiesFrom(
		listener,
		self.dbusDestinationObject,
		"org.freedesktop.login1.Session.IdleHint",
		"/org/freedesktop/login1/session/self")
}

func (self LoginManager) readBlockInhibited(listener chan DBusMessage) {
	self.adapter.readCurrentPropertiesFrom(
		listener,
		self.dbusDestinationObject,
		"org.freedesktop.login1.Manager.BlockInhibited",
		"/org/freedesktop/login1")
}

func (self LoginManager) readCurrentInhibitors(listener chan DBusMessage) {
	// See:
	// dbus-send --system --print-reply --type="method_call"  --dest=org.freedesktop.login1 /org/freedesktop/login1  org.freedesktop.login1.Manager.ListInhibitors
	var response = new(InhibitorResponse)
	self.adapter.readCurrentInhibitorsFrom(
		listener,
		"org.freedesktop.login1",
		"org.freedesktop.login1.Manager.ListInhibitors",
		"/org/freedesktop/login1",
		response)
}

type InhibitorResponse struct {
	inhibitions string
	issuer      string
	reason      string
	action      string
	a           int32
	b           int32
}
