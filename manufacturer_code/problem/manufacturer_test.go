package vin_test

import (
	"testing"
	"vin"
)

const testVIN = "W09000051T2123456"

func TestVIN_Manufacturer(t *testing.T) {
	manufacturer := vin.Manufacturer(testVIN)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

// panic
// func TestVIN_ManufacturerShort(t *testing.T) {
// 	testVINShort := "01"
// 	manufacturer := vin.Manufacturer(testVINShort)
// 	if manufacturer != "01" {
// 		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
// 	}
// }
