package x11

import (
	"fmt"
	"github.com/jezek/xgb/dpms"
	"os"
)

type DPMS struct {
	adapter *X11Connection

	IsDpmsAvailable bool
}

func (self *DPMS) Init(adapter *X11Connection) {
	self.adapter = adapter

	err := dpms.Init(adapter.conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "No DPMS", err)
		self.IsDpmsAvailable = false
	} else {
		self.IsDpmsAvailable = true
	}
}

func (self *DPMS) IsApplicable() bool {
	return self.IsDpmsAvailable
}

func (self DPMS) Read() {
	self.readTimeouts()
}

func (self DPMS) readTimeouts() {
	timeouts, err := dpms.GetTimeouts(self.adapter.conn).Reply()

	if err != nil {
		fmt.Fprintln(os.Stderr, "DPMS timeouts: ", err)
	} else {
		// TODO: Assemble facts
		fmt.Fprintln(os.Stdout, "DPMS: ", timeouts.OffTimeout, timeouts.StandbyTimeout, timeouts.StandbyTimeout)
	}
}
