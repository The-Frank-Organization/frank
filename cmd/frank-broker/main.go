package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/The-Frank-Organization/frank/internal/broker"
)

func main() {
	os.Exit(run())
}

func run() int {
	flags := flag.NewFlagSet("frank-broker", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	configHome := flags.String("config-home", "", "private broker configuration home")
	if err := flags.Parse(os.Args[1:]); err != nil || *configHome == "" || flags.NArg() != 0 {
		return 2
	}
	tokenPipe := os.NewFile(3, "broker-control-token")
	if tokenPipe == nil {
		return 1
	}
	server, err := broker.NewServer(*configHome, tokenPipe)
	_ = tokenPipe.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "frank-broker: startup unavailable")
		return 1
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := server.Serve(ctx, ".", os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "frank-broker: service unavailable")
		return 1
	}
	return 0
}
