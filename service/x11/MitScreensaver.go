package x11

import (
	"fmt"
	"github.com/jezek/xgb/screensaver"
	"os"
)

type MitScreensaver struct {
	adapter *X11Connection
}

func (self *MitScreensaver) Init(adapter *X11Connection) {
	self.adapter = adapter

	screensaver.Init(adapter.conn)
}

func (self *MitScreensaver) IsApplicable() bool {
	return true // TODO: Not always the case!
}

func (self MitScreensaver) Read() {
	self.readInfo()
}

func (self MitScreensaver) readInfo() {
	// TODO: Fix Drawable
	info, err := screensaver.QueryInfo(self.adapter.conn, 0).Reply()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Mit ScreenSaver: ", err)
	} else {
		// TODO: Assemble facts
		fmt.Fprintln(os.Stdout, "Mit ScreenSaver: ", info.State, info.MsSinceUserInput, info.MsUntilServer)
	}
}
