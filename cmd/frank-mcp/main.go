package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	opts, err := optionsFromArgs(os.Args[1:], os.Getenv, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if err := NewMCPServer(opts).Serve(); err != nil {
		fmt.Fprintf(os.Stderr, "frank-mcp protocol error: %v\n", err)
		os.Exit(1)
	}
}

func optionsFromArgs(args []string, getenv func(string) string, stdin io.Reader, stdout, stderr io.Writer) (Options, error) {
	fs := flag.NewFlagSet("frank-mcp", flag.ContinueOnError)
	fs.SetOutput(stderr)
	socket := fs.String("socket", getenv("FRANK_SOCKET"), "frank conductor socket path")
	credentialFile := fs.String("credential-file", "", "0600 file containing the frank credential")
	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}
	credential := getenv("FRANK_CREDENTIAL")
	if credential == "" && *credentialFile != "" {
		data, err := os.ReadFile(*credentialFile)
		if err != nil {
			return Options{}, fmt.Errorf("read credential file: %w", err)
		}
		info, err := os.Stat(*credentialFile)
		if err != nil {
			return Options{}, fmt.Errorf("stat credential file: %w", err)
		}
		if info.Mode().Perm()&0o077 != 0 {
			return Options{}, fmt.Errorf("credential file must be 0600")
		}
		credential = strings.TrimSpace(string(data))
	}
	return Options{
		SocketPath: *socket,
		Credential: credential,
		Context:    context.Background(),
		Stdin:      stdin,
		Stdout:     stdout,
		Stderr:     stderr,
	}, nil
}
