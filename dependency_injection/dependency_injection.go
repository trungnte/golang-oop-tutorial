package vin

type VINAPIClient struct {
	apiURL string
	apiKey string
	// ... internals go here...
}

func NewVINAPIClient(apiURL, apiKey string) *VINAPIClient {
	return &VINAPIClient{apiURL, apiKey}
}

func (client *VINAPIClient) IsEuropean(code string) bool {
	// calls external API and returns correct value
	//! API could be offline or just not reachable. Also
	//! calling an external API takes time and may cost money
	return true
}

/*
The only issue here is that the test requires a live connection to the external API.
This is unfortunate, as the API could be offline or just not reachable.
Also, calling an external API takes time and may cost money.

As the result of the API call is known, it should be possible to replace it with a mock.
Unfortunately, in the code above, the VINService itself creates the API client,
so there is no easy way to replace it. To make this possible, the API client
dependency should be injected into the VINService. That is, it should be
created before calling the VINService constructor.

The only issue here is that the test requires a live connection to the external API.
This is unfortunate, as the API could be offline or just not reachable.
Also, calling an external API takes time and may cost money.

As the result of the API call is known, it should be possible to replace
it with a mock. Unfortunately, in the code above, the VINService itself
creates the API client, so there is no easy way to replace it.
To make this possible, the API client dependency should be injected into
the VINService. That is, it should be created before calling the VINService constructor.
*/
type VINService struct {
	client *VINAPIClient
}

type VINServiceConfig struct {
	APIURL string
	APIKEY string
	// more configuration values
}

func NewVINService(config *VINServiceConfig) *VINService {
	// use config to create the API client
	apiClient := NewVINAPIClient(config.APIURL, config.APIKEY)
	return &VINService{apiClient}
}

func (s *VINService) CreateFromCode(code string) (VIN, error) {
	if s.client.IsEuropean(code) {
		return NewEUVIN(code)
	}

	return NewVIN(code)
}
