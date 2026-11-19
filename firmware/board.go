//go:build esp32c3_generic

package main

import "machine"

const (
	qwiicSDA = machine.GPIO5
	qwiicSCL = machine.GPIO6
)
