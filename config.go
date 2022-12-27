package main

import "os"

type config struct {
	isRoot bool

	isDbusSystemConnection bool
}

func configure() config {
	var isRoot = os.Geteuid() == 0

	return config{
		isRoot:                 isRoot,
		isDbusSystemConnection: isRoot,
	}
}
