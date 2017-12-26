package main

import "strings"

func stripPrefix(name string) string {
	if strings.HasPrefix(name, "goske-") {
		name = name[6:]
	}
	return name
}
