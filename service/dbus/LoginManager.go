package dbus

type LoginManager struct {
	adapter *DBusAdapter
}

func (self *LoginManager) Init(adapter *DBusAdapter) {
	self.adapter = adapter
}

func (self LoginManager) Read(listener chan DBusMessage) {
	self.readIdleHint(listener)
	self.readCurrentInhibitors(listener)
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
		"org.freedesktop.login1",
		"org.freedesktop.login1.Session.IdleHint",
		"/org/freedesktop/login1/session/self")
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
