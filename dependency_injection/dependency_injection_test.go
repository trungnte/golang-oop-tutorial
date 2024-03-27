package vin_test

import (
	"testing"
	"vin"
)

const euSmallVIN = "W09000051T2123456"

func TestVIN_EU_SmallManufacturer(t *testing.T) {

	service := vin.NewVINService(&vin.VINServiceConfig{})
	testVIN, _ := service.CreateFromCode(euSmallVIN)

	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}
