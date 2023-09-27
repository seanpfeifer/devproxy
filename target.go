package main

import (
	"fmt"
	"net/url"
	"strings"
)

// A single proxy target is a string in the format of "/url/on/local/->http://remote:port/url/on/remote/"
type proxyTarget string

// Parse parses a proxy target into a local path and a remote URL.
func (p proxyTarget) Parse() (local string, remote *url.URL, err error) {
	s := string(p)
	i := strings.Index(s, "->")
	if i == -1 {
		return "", nil, fmt.Errorf("invalid proxy target %q", s)
	}

	proxyURL, err := url.Parse(s[i+2:])
	return s[:i], proxyURL, err
}

// proxyTargets is our implementation of the flag.Value interface so it can be used with flag.Var()
type proxyTargets []proxyTarget

func (p *proxyTargets) String() string {
	return fmt.Sprintf("%v", *p)
}

func (p *proxyTargets) Set(value string) error {
	*p = append(*p, proxyTarget(value))
	return nil
}
