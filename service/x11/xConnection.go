package x11

import (
	"github.com/jezek/xgb"
	"log"
)

type X11Source interface {
	Init(adapter *X11Connection)
	Read()
	IsApplicable() bool
}

type X11Connection struct {
	conn *xgb.Conn
}

func (self *X11Connection) Init() {
	X, err := xgb.NewConn()

	if err != nil {
		log.Fatal(err)
	}

	self.conn = X
}
