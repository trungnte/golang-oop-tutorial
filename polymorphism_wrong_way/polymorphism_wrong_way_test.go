package vin_test

import (
	"testing"
	"vin"
)

const euSmallVIN = "W09000051T2123456"

// ! this works!
func TestVIN_EU_SmallManufacturer(t *testing.T) {
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

// ! this fails with an error
func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {
	var testVINs []vin.VIN
	testVIN, _ := vin.NewEUVIN(euSmallVIN)
	// having to cast testVIN already hints something is odd
	testVINs = append(testVINs, vin.VIN(testVIN))

	for _, vin := range testVINs {
		// this call vin.Manufacturer() will call super -> only return "W09"
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}
