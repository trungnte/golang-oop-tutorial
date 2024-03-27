package vin

import (
	"fmt"
)

// TODO: make it elegant with struct and struct -> so I comment this code to implement new struct
// type VIN struct {
// 	code     string
// 	european bool
// }

// // validate len of string
// func NewVIN(code string, european bool) (*VIN, error) {
// 	if len(code) != 17 {
// 		return nil, fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
// 	}

// 	// ... check for disallowed characters ...

// 	//! VIN(code): cast code (string) as VIN type and return
// 	return &VIN{code, european}, nil
// }

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
	return string(v[:3])
}

type EUVIN VIN

func NewEUVIN(code string) (EUVIN, error) {
	// call super constructor
	v, err := NewVIN(code)

	// and cast to subtype
	return EUVIN(v), err
}

func (v EUVIN) Manufacturer() string {
	// call manufacturer on supertype
	manufacturer := VIN(v).Manufacturer()

	// add EU specific postfix if appropriate
	if manufacturer[2] == '9' {
		manufacturer += string(v[11:14])
	}

	return manufacturer
}
