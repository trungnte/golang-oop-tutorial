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

// ! this works. Compare code with ../polymorphism_wrong_way/polymorphism_wrong_way_test.go
func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {
	var testVINs []vin.VIN // this type is struct that has interface
	testVIN, _ := vin.NewEUVIN(euSmallVIN)

	//! now there is no need to cast!
	testVINs = append(testVINs, testVIN)

	for _, vin := range testVINs {
		// this call vin.Manufacturer() will call super -> only return "W09"
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}
