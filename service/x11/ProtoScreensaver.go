package x11

import (
	"fmt"
	"github.com/jezek/xgb/xproto"
	"os"
)

type ProtoScreensaver struct {
	adapter *X11Connection
}

func (self *ProtoScreensaver) Init(adapter *X11Connection) {
	self.adapter = adapter
}

func (self *ProtoScreensaver) IsApplicable() bool {
	return true // TODO: Not always the case!
}

func (self ProtoScreensaver) Read() {
	self.readTimeouts()
}

func (self ProtoScreensaver) readTimeouts() {
	settings, err := xproto.GetScreenSaver(self.adapter.conn).Reply()

	if err != nil {
		fmt.Fprintln(os.Stderr, "X11 ScreenSaver: ", err)
	} else {
		// TODO: Assemble facts
		fmt.Fprintln(os.Stdout, "X11 ScreenSaver: ", settings.Interval, settings.Timeout, settings.PreferBlanking)
	}
}
