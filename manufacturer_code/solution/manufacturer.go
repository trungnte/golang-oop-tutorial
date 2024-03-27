package vin

import (
	"fmt"
)

type VIN string

// using constructors to validate data

// it is debatable if this func should be named New or NewVIN
// but NewVIN is better for greping and leaves room for other
// NewXY funcs in the same package

func NewVIN(code string) (VIN, error) {
	if len(code) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", code)
	}

	return VIN(code), nil
}

func (v VIN) Manufacturer() string {
	// manufacturer := v[:3]
	// if manufacturer[2] == '9' {
	// 	manufacturer += v[11:14]
	// }

	// return string(manufacturer)
	return string(v[:3])
}

type EUVIN VIN

func NewEUVIN(code string) (EUVIN, error) {
	// call super constructor
	v, err := NewVIN(code)

	// cast and return
	return EUVIN(v), err
}

func (v EUVIN) Manufacturer() string {
	manufacturer := VIN(v).Manufacturer()

	if manufacturer[2] == '9' {
		manufacturer += string(v[11:14])
	}

	return manufacturer
}
