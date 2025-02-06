package system

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	quit chan bool
)

// Exit Exit
func Exit() {
	quit <- true
}

// WaitSignal WaitSignal
func WaitSignal() {

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)

	for {
		select {
		case sig := <-c:
			//s.Debugf("Trapped %q signal", sig)
			switch sig {
			case syscall.SIGINT:
				os.Exit(0)
			case syscall.SIGUSR1:
				// File log re-open for rotating file logs.
				//s.ReOpenLogFile()
			case syscall.SIGUSR2:
				//go s.lameDuckMode()
			case syscall.SIGHUP:
				// Config reload.
				// if err := s.Reload(); err != nil {
				// 	s.Errorf("Failed to reload server configuration: %s", err)
				// }
			}
			_ = sig
			return
		case <-quit:
			return
		}
	}
}
