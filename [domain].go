package main

import "strings"

// matchesEmailDomain checks a normalized email against an @domain suffix.
func matchesEmailDomain(email, suffix string) bool {
	return strings.Contains(email, suffix)
}
