// Copyright (c) Palo Alto Networks, Inc.
// SPDX-License-Identifier: MPL-2.0

package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/PaloAltoNetworks/terraform-provider-cortexcloud/internal/util"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	defaultInsecure             bool = false
	defaultRequestTimeout       int  = 60
	defaultRequestRetryInterval int  = 3
	//obscuredApiKey string
)

type CortexCloudAPIClientConfig struct {
	ApiURL               *string `tfsdk:"api_url" json:"api_url"`
	ApiKey               *string `tfsdk:"api_key" json:"api_key"`
	ApiKeyId             *int    `tfsdk:"api_key_id" json:"api_key_id"`
	Insecure             *bool   `tfsdk:"insecure" json:"insecure"`
	RequestTimeout       *int    `tfsdk:"request_timeout" json:"request_timeout"`
	RequestRetryInterval *int    `tfsdk:"request_retry_interval" json:"request_retry_interval"`
	CrashStackDir        *string `tfsdk:"crash_stack_dir" json:"crack_stack_dir"`
}

// CortexCloudAPIClient implements the HTTP client that will be used to execute
// requests to the Cortex Cloud API.
type CortexCloudAPIClient struct {
	Config     CortexCloudAPIClientConfig
	HTTPClient *http.Client
}

// NewCortexCloudAPIClient configures and returns a new CortexCloudAPIClient
// using the values defined in the provider.
//
// If a given provider attribute is not configured, the value will be retrieved
// from the associated environment variable. If the environment variable is empty
// or not set, the provider will return an error diagnostic if the attribute is
// optional, else it will use the specified default value.
func NewCortexCloudAPIClient(ctx context.Context, config CortexCloudAPIClientConfig) (*CortexCloudAPIClient, error) {
	// Parse request timeout config value
	if config.RequestTimeout == nil {
		config.RequestTimeout = &defaultRequestTimeout
	} else if *config.RequestTimeout > math.MaxInt { // TODO: cap this at something sane
		return nil, fmt.Errorf("error occured while creating API client: Invalid value supplied for request_timeout. Value must be an integer between 1 and %d.", math.MaxInt)
	}

	requestTimeout, err := time.ParseDuration(fmt.Sprintf("%ds", *config.RequestTimeout))
	if err != nil {
		return nil, fmt.Errorf("error occured while creating API client: Failed to parse request timeout value\n%s", err.Error())
	}
	//            //    fmt.Sprintf("Error configuring provider: Invalid value specified for \"request_timeout\" in configuration file. Value must be an integer between 1 and %d", math.MaxInt),

	// Parse request retry interval config value
	if config.RequestRetryInterval == nil {
		config.RequestRetryInterval = &defaultRequestRetryInterval
	} else if *config.RequestRetryInterval > math.MaxInt { // TODO: cap this at something sane
		return nil, fmt.Errorf("error occured while creating API client: Invalid value supplied for request_retry_interval. Value must be an integer between 1 and %d.", math.MaxInt)
	}

	//requestRetryInterval, err := time.ParseDuration(fmt.Sprintf("%ds", *config.RequestRetryInterval))
	//if err != nil {
	//	return nil, fmt.Errorf("error occured while creating API client: Failed to parse request retry interval value\n%s", err.Error())
	//}

	// Instantiate HTTP client
	httpClient := &http.Client{
		Timeout: requestTimeout,
	}

	// If the insecure flag is set to true, add TLS client configuration with InsecureSkipVerify enabled
	if config.Insecure != nil && *config.Insecure {
		transport := http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		(*httpClient).Transport = &transport
	}

	apiClient := &CortexCloudAPIClient{
		Config:     config,
		HTTPClient: httpClient,
	}

	// Send request to health check endpoint to ensure credentials are valid
	// and the Cortex Cloud tenant is reachable
	tflog.Debug(ctx, "Running health check")
	healthCheckResponse, err := apiClient.HealthCheck(ctx)
	if err != nil {
		return nil, err
	} else if healthCheckResponse.Status != "available" {
		return nil, fmt.Errorf("Health check request failed (returned status \"%s\")", healthCheckResponse.Status)
	}

	return apiClient, nil
}

// HealthCheckResponse represents the response body of the health check endpoint.
type HealthCheckResponse struct {
	Status string `tfsdk:"status" json:"status"`
}

// HealthCheck sends a request to the health check endpoint and returns the response.
//
// If the request fails, an error message is returned suggesting that the user
// check the provider configuration and verify that they can reach their Cortex
// Cloud tenant's API URL.
func (c *CortexCloudAPIClient) HealthCheck(ctx context.Context) (HealthCheckResponse, error) {
	var response HealthCheckResponse

	if err := c.Request(ctx, "GET", HealthCheckEndpoint, nil, nil, &response); err != nil {
		return response, fmt.Errorf("Health check request failed: %s \nVerify that your provider configuration is correct and the Cortex Cloud API is reachable, then try again.", err.Error())
	}

	return response, nil
}

// Request sends an HTTP request with the http.Client in the CortexCloudAPIClient object.
func (c *CortexCloudAPIClient) Request(ctx context.Context, method, endpoint string, query, data, responseBody interface{}) error {
	// TODO: create getters/setters for config values so we dont need nil checks in here
	if c.Config.RequestRetryInterval == nil {
		requestRetryInterval := 3
		c.Config.RequestRetryInterval = &requestRetryInterval
	}

	requestRetryInterval := *c.Config.RequestRetryInterval

	// Parse API URL from config
	apiUrl, err := url.Parse(*c.Config.ApiURL)
	if err != nil {
		return err
	}

	// Append endpoint to URL
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	// Marshal request payload into buffer, if not nil
	var (
		payloadBuffer bytes.Buffer
		jsonData      []byte
	)
	if data != nil {
		jsonData, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		payloadBuffer = *bytes.NewBuffer(jsonData)
	}

	// Create new HTTP request object
	req, err := http.NewRequestWithContext(ctx, method, apiUrl.String(), &payloadBuffer)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("x-xdr-auth-id", fmt.Sprintf("%d", *c.Config.ApiKeyId)) // Hacky int-to-string conversion
	req.Header.Set("Authorization", *c.Config.ApiKey)
	//req.Header.Set("Content-Type", "application/json")

	// If query parameters are defined, add them to request URL
	if query != nil {
		queryParams := req.URL.Query()

		if queryMap, ok := query.(map[string]string); ok {
			for key, val := range queryMap {
				queryParams.Add(key, val)
			}
		} else {
			return fmt.Errorf("failed to parse query parameters: %v", query)
		}

		req.URL.RawQuery = queryParams.Encode()
	}

	// Print a debug log of the HTTP request details and create a UUID for the
	// request to track its response in the debug logs
	requestId := util.LogHttpRequest(ctx, req.Method, req.URL.String(), req.Header.Get("x-xdr-auth-id"), req.Header.Get("Authorization"), string(jsonData))

	// Execute HTTP request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			return fmt.Errorf("context cancelled or timeout exceeded: %s", ctx.Err())
		}

		return err
	}
	defer res.Body.Close()

	// Check HTTP response status and populate response body
	err, retryRequested := getHttpResponseBody(res, responseBody)
	if err != nil {
		return err
	}

	// Attempt to print a debug log of the response, or an error message if
	// the response body cannot be unmarshalled
	err = util.LogHttpResponse(ctx, requestId, res.StatusCode, responseBody)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("Failed to unmarshal body of HTTP response for request_uid=%s: %s", requestId, err.Error()))
	}

	// If getHttpResponseBody returns true for the boolean return value, the API
	// returned a 429 status code so we need to sleep for the number of seconds
	// dictated by the requestRetryInterval var and then retry the request
	if retryRequested {
		tflog.Debug(ctx, fmt.Sprintf("API returned 429 (too many requests), sleeping for %d seconds and retrying...", *c.Config.RequestRetryInterval))
		time.Sleep(time.Duration(requestRetryInterval) * time.Second)
		return c.Request(ctx, method, endpoint, query, data, responseBody)
	}

	return nil
}

func getHttpResponseBody(response *http.Response, responseValuePtr interface{}) (error, bool) {
	isOkResponse := (response.StatusCode == http.StatusOK)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		if !isOkResponse {
			return fmt.Errorf("error reading body from non-OK response: %s", err.Error()), false
		} else {
			return fmt.Errorf("error reading body from OK response: %s", err.Error()), false
		}
	}

	switch response.StatusCode {
	case 200:
		// Attempt to unmarshal the response body into responseValuePtr
		// and return
		if len(body) > 0 && responseValuePtr != nil {
			if err = json.Unmarshal(body, responseValuePtr); err != nil {
				return fmt.Errorf("error unmarshalling body from OK response: %s", err.Error()), false
			}
		}
		return nil, false
	case 429:
		// Return true for the boolean return value to signal to the Request
		// method that we need to sleep and then retry the request
		return nil, true
	case 401:
		// Return an error indicating that the provider configuration values
		// should be verified
		return fmt.Errorf("API returned 401 (Unauthorized). Verify your provider configuration values and try again."), false
	default:
		// If the response contains a non-zero length body, attempt to
		// unmarshal the response body and return an error with its contents
		if len(body) > 0 {
			var errorResponse map[string]any
			if err = json.Unmarshal(body, &errorResponse); err != nil {
				return fmt.Errorf("error unmarshalling body from non-OK response (HTTP status code %d): %s", response.StatusCode, err.Error()), false
			}

			var errorResponseStr string
			for k, v := range errorResponse["reply"].(map[string]any) {
				errorResponseStr += fmt.Sprintf("\t%s: %v\n", k, v)
			}

			return fmt.Errorf("API returned non-OK response: \n%v", errorResponseStr), false
			// Otherwise, return an error with just the status code
		} else {
			return fmt.Errorf("API returned non-OK response (HTTP status code %d)", response.StatusCode), false
		}
	}
}
