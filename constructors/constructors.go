package vin

import "fmt"

type VIN string

// validate len of string
func NewVIN(code string) (VIN, error) {
	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	// ... check for disallowed characters ...

	//! VIN(code): cast code (string) as VIN type and return
	return VIN(code), nil
}

// parse and get 3 first characters and 3 (11, 12, 13) characters of VIN string
func (v VIN) Manufacturer() string {
	manufacturer := v[:3]
	if manufacturer[2] == '9' {
		manufacturer += v[11:14]
	}
	return string(manufacturer)
}
