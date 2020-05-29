package exithandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Init accepts a callback function which will be
// invoked when program exits unexpectedly or is terminated by user
func Init(cb func()) {
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println("exit reasons: ", sig)
		terminate <- true
	}()

	<-terminate
	cb()
}
