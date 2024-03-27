package vin_test

import (
	"testing"
	"vin"
)

const (
	validVIN   = "W0L000051T2123456"
	invalidVIN = "W0"
)

// func TestVIN_New(t *testing.T) {
// 	_, err := vin.NewVIN(validVIN)
// 	if err != nil {
// 		t.Errorf("creating valid VIN returned an error: %s", err.Error())
// 	}

// 	_, err = vin.NewVIN(invalidVIN)
// 	if err == nil {
// 		t.Error("creating invalid VIN did not return an error")
// 	}
// }

// func TestVIN_Manufacturer(t *testing.T) {
// 	testVIN, _ := vin.NewVIN(validVIN)
// 	manufacturer := testVIN.Manufacturer()
// 	if manufacturer != "W0L" {
// 		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVIN)
// 	}

// 	// panic
// 	// invalidVIN.Manufacturer()
// }

// panic
// func TestVIN_ManufacturerShort(t *testing.T) {
// 	testVINShort := "01"
// 	manufacturer := vin.Manufacturer(testVINShort)
// 	if manufacturer != "01" {
// 		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
// 	}
// }

const euSmallVIN = "W09000051T2123456"

// this works!
func TestVIN_EU_SmallManufacturer(t *testing.T) {

	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

// this fails with an error
func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {

	var testVINs []vin.VIN
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	// having to cast testVIN already hints something is odd
	testVINs = append(testVINs, vin.VIN(testVIN))

	for _, vin := range testVINs {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}
