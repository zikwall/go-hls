package main

import (
	"os"
	"os/signal"
	"syscall"
)

func waitSystemNotify() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
}
