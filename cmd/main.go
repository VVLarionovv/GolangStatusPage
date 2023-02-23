package main

import (
	h "itog/internal/hendler"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		h.Start()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}
