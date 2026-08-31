package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/The-Frank-Organization/frank/internal/connector/control"
	"github.com/The-Frank-Organization/frank/internal/connector/service"
)

var errRuntimeDir = errors.New("connector: invalid private runtime directory")

type config struct {
	credentialPath string
	catalogPath    string
	policyPath     string
	controlFD      int
	dataFD         int
	deathFD        int
	runtimeDir     string
	buildInfo      string
}

func main() {
	config, err := parseConfig(os.Args[1:])
	if err != nil || run(context.Background(), config) != nil {
		os.Exit(1)
	}
}

func parseConfig(arguments []string) (config, error) {
	set := flag.NewFlagSet("frank-connector", flag.ContinueOnError)
	set.SetOutput(ioDiscard{})
	var result config
	set.StringVar(&result.credentialPath, "credential", "", "credential artifact path")
	set.StringVar(&result.catalogPath, "catalog", "", "lane catalog path")
	set.StringVar(&result.policyPath, "policy", "", "egress policy path")
	set.IntVar(&result.controlFD, "control-fd", -1, "inherited CTRL-C descriptor")
	set.IntVar(&result.dataFD, "data-fd", -1, "inherited DATA-P descriptor")
	set.IntVar(&result.deathFD, "death-fd", -1, "inherited parent-death descriptor")
	set.StringVar(&result.runtimeDir, "runtime-dir", "", "private runtime directory")
	set.StringVar(&result.buildInfo, "build-info", "", "build identity")
	if err := set.Parse(arguments); err != nil || set.NArg() != 0 {
		return config{}, errors.New("connector: invalid argv")
	}
	paths := []string{result.credentialPath, result.catalogPath, result.policyPath, result.runtimeDir}
	for _, path := range paths {
		if path == "" || !filepath.IsAbs(path) || filepath.Clean(path) != path {
			return config{}, errors.New("connector: artifact paths must be clean and absolute")
		}
	}
	for _, artifactPath := range paths[:3] {
		relative, err := filepath.Rel(result.runtimeDir, artifactPath)
		if err != nil || relative == "." || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			return config{}, errors.New("connector: artifact path escapes private runtime directory")
		}
	}
	if result.controlFD < 3 || result.dataFD < 3 || result.deathFD < 3 || result.controlFD == result.dataFD || result.controlFD == result.deathFD || result.dataFD == result.deathFD || result.buildInfo == "" {
		return config{}, errors.New("connector: invalid descriptor or build identity")
	}
	return result, nil
}

func run(ctx context.Context, config config) error {
	if err := validateRuntimeDir(config.runtimeDir); err != nil {
		return err
	}
	if err := disableCoreDumps(); err != nil {
		return err
	}
	if err := markCloseOnExec(config.controlFD); err != nil {
		return err
	}
	if err := markCloseOnExec(config.dataFD); err != nil {
		return err
	}
	if err := markCloseOnExec(config.deathFD); err != nil {
		return err
	}
	controlFile := os.NewFile(uintptr(config.controlFD), "frank-ctrl-c")
	dataFile := os.NewFile(uintptr(config.dataFD), "frank-data-p")
	deathFile := os.NewFile(uintptr(config.deathFD), "frank-parent-death")
	if controlFile == nil || dataFile == nil || deathFile == nil {
		return errors.New("connector: inherited descriptor unavailable")
	}
	defer controlFile.Close()
	defer dataFile.Close()
	defer deathFile.Close()
	ctx, cancel := cancelOnParentDeath(ctx, deathFile)
	defer cancel()
	os.Clearenv()
	session, err := control.Bootstrap(ctx, controlFile, control.Config{
		Artifacts: control.Artifacts{
			CredentialPath: config.credentialPath,
			CatalogPath:    config.catalogPath,
			PolicyPath:     config.policyPath,
		},
		BuildInfo: config.buildInfo,
		PID:       os.Getpid(),
	})
	if err != nil {
		return err
	}
	defer session.Close()
	server, err := service.New(service.Config{Control: session, Data: dataFile, BuildInfo: config.buildInfo})
	if err != nil {
		return err
	}
	err = server.Serve(ctx)
	if errors.Is(err, control.ErrShutdown) {
		return nil
	}
	return err
}

func cancelOnParentDeath(parent context.Context, death io.Reader) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	go func() {
		var signal [1]byte
		_, _ = death.Read(signal[:])
		cancel()
	}()
	return ctx, cancel
}

func disableCoreDumps() error {
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_CORE, &limit); err != nil {
		return err
	}
	limit.Cur = 0
	return syscall.Setrlimit(syscall.RLIMIT_CORE, &limit)
}

func validateRuntimeDir(path string) error {
	info, err := os.Lstat(path)
	if err != nil || !info.IsDir() || info.Mode().Perm() != 0o700 || info.Mode()&os.ModeSymlink != 0 {
		return errRuntimeDir
	}
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok || int(stat.Uid) != os.Geteuid() {
		return errRuntimeDir
	}
	return nil
}

func markCloseOnExec(fd int) error {
	_, _, errno := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(syscall.F_SETFD), uintptr(syscall.FD_CLOEXEC))
	if errno != 0 {
		return errno
	}
	return nil
}

type ioDiscard struct{}

func (ioDiscard) Write(value []byte) (int, error) { return len(value), nil }
