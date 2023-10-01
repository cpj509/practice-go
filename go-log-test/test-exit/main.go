package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "sdi:nsdi@1234@tcp(192.168.5.204:3307)/sdi_ems")
	if err != nil {
		// handle the error
	}

	// Create a context that will be cancelled when the process is killed
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive signals from the operating system
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Start a goroutine to wait for signals and cancel the context
	go func() {
		<-signalCh
		cancel()

	}()

	// Use the context to ensure that the database connection is closed when the process is killed
	err = db.PingContext(ctx)
	if err != nil {
		// handle the error
	}

	err = db.Close()
	if err != nil {
		// handle the error
	}
}
