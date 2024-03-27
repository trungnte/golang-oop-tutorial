package vin

type VINAPIClient interface {
	IsEuropean(code string) bool
}

type vinAPIClient struct {
	apiURL string
	apiKey string
}

func (client *vinAPIClient) IsEuropean(code string) bool {

	// calls external API and returns something more useful
	return true
}

func NewVINAPIClient(apiURL, apiKey string) *vinAPIClient {
	return &vinAPIClient{apiURL, apiKey}
}

type VINService struct {
	client VINAPIClient
}

type VINServiceConfig struct {
	// more configuration values
}

func NewVINService(config *VINServiceConfig, apiClient VINAPIClient) *VINService {
	// apiClient is created elsewhere and injected here
	return &VINService{apiClient}
}

func (s *VINService) CreateFromCode(code string) (VIN, error) {
	if s.client.IsEuropean(code) {
		return NewEUVIN(code)
	}

	return NewVIN(code)
}
