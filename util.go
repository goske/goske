package main

import (
	"os"
	"strings"
)

func stripPrefix(name string) string {
	if strings.HasPrefix(name, "goske-") {
		name = name[6:]
	}
	return name
}

func goskeRepo() string {
	s := os.Getenv("GITHUB_GOSKE")
	if s != "" {
		return s
	}
	return "goske"
}
