// Package endpoint validates the restricted, unique-spelling provider endpoint
// grammar owned by the m-3 egress contract.
package endpoint

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Validate accepts an endpoint only when it already has the unique spelling
// required by the provider-request policy. It never normalizes input.
func Validate(value string) error {
	const scheme = "https://"
	if !strings.HasPrefix(value, scheme) {
		return fmt.Errorf("endpoint: scheme must be https")
	}
	remainder := value[len(scheme):]
	slash := strings.IndexByte(remainder, '/')
	if slash <= 0 {
		return fmt.Errorf("endpoint: path is required")
	}
	authority, path := remainder[:slash], remainder[slash:]
	if strings.ContainsAny(authority, "@[]") || strings.Count(authority, ":") > 1 {
		return fmt.Errorf("endpoint: invalid authority")
	}
	host := authority
	if colon := strings.LastIndexByte(authority, ':'); colon >= 0 {
		host = authority[:colon]
		port := authority[colon+1:]
		if port == "" || (len(port) > 1 && port[0] == '0') {
			return fmt.Errorf("endpoint: invalid port")
		}
		for i := range port {
			if port[i] < '0' || port[i] > '9' {
				return fmt.Errorf("endpoint: invalid port")
			}
		}
		parsed, err := strconv.ParseUint(port, 10, 16)
		if err != nil || parsed == 0 || parsed == 443 {
			return fmt.Errorf("endpoint: invalid or default port")
		}
	}
	if err := validateHost(host); err != nil {
		return err
	}
	if err := validatePath(path); err != nil {
		return err
	}
	return nil
}

func validateHost(host string) error {
	if host == "" || net.ParseIP(host) != nil || strings.HasSuffix(host, ".") {
		return fmt.Errorf("endpoint: invalid host")
	}
	for _, label := range strings.Split(host, ".") {
		if label == "" || !isLowerAlnum(label[0]) || !isLowerAlnum(label[len(label)-1]) {
			return fmt.Errorf("endpoint: invalid host label")
		}
		for i := 1; i+1 < len(label); i++ {
			if !isLowerAlnum(label[i]) && label[i] != '-' {
				return fmt.Errorf("endpoint: invalid host label")
			}
		}
	}
	return nil
}

func validatePath(path string) error {
	if path == "/" {
		return nil
	}
	if len(path) < 2 || path[0] != '/' || path[len(path)-1] == '/' || strings.ContainsAny(path, "%?#") {
		return fmt.Errorf("endpoint: invalid path")
	}
	for _, segment := range strings.Split(path[1:], "/") {
		if segment == "" || segment == "." || segment == ".." {
			return fmt.Errorf("endpoint: invalid path segment")
		}
		for i := range segment {
			if !isPathByte(segment[i]) {
				return fmt.Errorf("endpoint: invalid path byte")
			}
		}
	}
	return nil
}

func isLowerAlnum(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= '0' && value <= '9'
}

func isPathByte(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9' || strings.ContainsRune("._~-", rune(value))
}
