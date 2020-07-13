package appgate

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/appgate/terraform-provider-appgate/client/v13/openapi"

	"github.com/google/uuid"
)

const (
	// Version is the Appgate controller version.
	Version = 12

	// DefaultDescription is the default string for terraform resources.
	DefaultDescription = "Managed by terraform"
)

// Config for appgate provider.
type Config struct {
	URL      string
	Username string
	Password string
	Provider string
	Insecure bool
	Timeout  int
}

// Client is the appgate API client.
type Client struct {
	// The AuthToken required for subsequent API calls.
	Token string
	// The controller version
	ControllerVersion string
	API               *openapi.APIClient
}

// Client creates
func (c *Config) Client() (*Client, error) {
	timeoutDuration := time.Duration(c.Timeout)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.Insecure,
		},
		Dial: (&net.Dialer{
			Timeout: timeoutDuration * time.Second,
		}).Dial,
		TLSHandshakeTimeout: timeoutDuration * time.Second,
	}

	httpclient := &http.Client{
		Transport: tr,
		Timeout:   ((timeoutDuration * 2) * time.Second),
	}
	clientCfg := &openapi.Configuration{
		// Host:   c.URL,
		// Scheme: "https",
		DefaultHeader: map[string]string{
			"Accept": fmt.Sprintf("application/vnd.appgate.peer-v%d+json", Version),
		},
		UserAgent: "Appgate-TerraformProvider/1.0.0/go",
		Debug:     true,
		Servers: []openapi.ServerConfiguration{
			{
				URL:         c.URL,
				Description: "Controller one",
			},
		},
		HTTPClient: httpclient,
	}
	apiClient := openapi.NewAPIClient(clientCfg)

	loginResponse, err := loginResponse(apiClient, c)
	if err != nil {
		return nil, err
	}
	client := &Client{
		API:               apiClient,
		Token:             fmt.Sprintf("Bearer %s", *openapi.PtrString(*loginResponse.Token)),
		ControllerVersion: *openapi.PtrString(*loginResponse.Version),
	}
	return client, nil
}

func loginResponse(apiClient *openapi.APIClient, cfg *Config) (openapi.LoginResponse, error) {

	ctx := context.Background()
	// Login first, save token
	loginOpts := openapi.LoginRequest{
		ProviderName: cfg.Provider,
		Username:     openapi.PtrString(cfg.Username),
		Password:     openapi.PtrString(cfg.Password),
		DeviceId:     uuid.New().String(),
	}

	loginResponse, _, err := apiClient.LoginApi.LoginPost(ctx).LoginRequest(loginOpts).Execute()
	if err != nil {
		return loginResponse, err
	}
	return loginResponse, nil
}
