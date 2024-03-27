package vin_test

import (
	"testing"
	"vin"
)

const euSmallVIN = "W09000051T2123456"

type mockAPIClient struct {
	apiCalls int
}

func NewMockAPIClient() *mockAPIClient {

	return &mockAPIClient{}
}

func (client *mockAPIClient) IsEuropean(code string) bool {

	client.apiCalls++
	return true
}

func TestVIN_EU_SmallManufacturer(t *testing.T) {
	apiClient := NewMockAPIClient()
	service := vin.NewVINService(&vin.VINServiceConfig{}, apiClient)
	testVIN, _ := service.CreateFromCode(euSmallVIN)

	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}

	if apiClient.apiCalls != 1 {
		t.Errorf("unexpected number of API calls: %d", apiClient.apiCalls)
	}
}
