package main

import (
	"fmt"
	"os"
)

// The worker runtime is transport-injected by the app control plane. The
// standalone command intentionally does not invent a provider or authority
// endpoint; live app-side wiring lands with the owning m-10/m-8 slices.
func main() {
	fmt.Fprintln(os.Stderr, "frank-worker: app control-plane and provider peers are required")
	os.Exit(2)
}
