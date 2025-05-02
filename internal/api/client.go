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

	//"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/util"
)

type CortexCloudAPIClientConfig struct {
	ApiURL         *string `tfsdk:"api_url" json:"api_url"`
	ApiKey         *string `tfsdk:"api_key" json:"api_key"`
	ApiKeyId       *int    `tfsdk:"api_key_id" json:"api_key_id"`
	Insecure       *bool   `tfsdk:"insecure" json:"insecure"`
	RequestTimeout *int    `tfsdk:"request_timeout" json:"request_timeout"`
	ConfigFile     *string `tfsdk:"config_file" json:"config_file"`
}

type CortexCloudAPIClient struct {
	Config     CortexCloudAPIClientConfig
	HTTPClient *http.Client
}

type ErrResponse struct {
	Err string
}

//type AuthRequest struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//type AuthResponse struct {
//	Token string `json:"token"`
//}

func NewCortexCloudAPIClient(ctx context.Context, config CortexCloudAPIClientConfig) (*CortexCloudAPIClient, error) {
	// Parse request timeout value
	if config.RequestTimeout == nil {
		defaultTimeout := 60
		config.RequestTimeout = &defaultTimeout
	} else if *config.RequestTimeout > math.MaxInt {
		return nil, fmt.Errorf("error occured while creating API client: Invalid value supplied for request_timeout. Value must be an integer between 1 and %d.", math.MaxInt)
	}

	requestTimeout, err := time.ParseDuration(fmt.Sprintf("%ds", *config.RequestTimeout))
	if err != nil {
		return nil, fmt.Errorf("error occured while creating API client: Failed to parse request timeout value\n%s", err.Error())
	}

	//            //    fmt.Sprintf("Error configuring provider: Invalid value specified for \"request_timeout\" in configuration file. Value must be an integer between 1 and %d", math.MaxInt),

	// Instantiate HTTP client
	httpClient := &http.Client{
		Timeout: requestTimeout,
	}

	// If the insecure flag is set to true, add TLS configuration with InsecureSkipVerify enabled
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

	// Authenticate to API
	if err := apiClient.Authenticate(ctx); err != nil {
		return nil, err
	}

	return apiClient, nil
}

// TODO: change this to use the equivalent of prisma's _ping endpoint just to make sure the api key works
func (c *CortexCloudAPIClient) Authenticate(ctx context.Context) (err error) {
	//util.LogDebug("Authenticating to Prisma Cloud Compute API")

	//res := AuthResponse{}

	//if c == nil {
	//	return fmt.Errorf("error occured while authenticating to Prisma Cloud Compute API: client uninitialized")
	//}

	////if c.Config.ConsoleURL == nil {
	//if c.Config.ApiURL == nil {
	//	return fmt.Errorf("error occured while authenticating to Prisma Cloud Compute API: nil console URL")
	//}

	//if c.Config.ApiKey == nil {
	//	return fmt.Errorf("error occured while authenticating to Prisma Cloud Compute API: nil username")
	//}

	//if c.Config.ApiKeyId == nil {
	//	return fmt.Errorf("error occured while authenticating to Prisma Cloud Compute API: nil password")
	//}

	//if err := c.Request(ctx, http.MethodPost, "api/v1/authenticate", nil, AuthRequest{*c.Config.Username, *c.Config.Password}, &res); err != nil {
	//	return fmt.Errorf("error occured while authenticating to Prisma Cloud Compute API: %v", err)
	//}
	//c.JWT = res.Token

	return nil
}

func (c *CortexCloudAPIClient) Request(ctx context.Context, method, endpoint string, query, data, response interface{}) (error) {
    var (
        payloadBuffer bytes.Buffer
        errorResponse ErrResponse
        err error
    )

	// Parse API URL from config
	apiUrl, err := url.Parse(*c.Config.ApiURL)
	if err != nil {
		return err
	}

	// Append endpoint to URL
	apiUrl.Path = path.Join(apiUrl.Path, endpoint)

	// Marshal request payload into buffer, if not nil
	if data != nil {
		data_json, err := json.Marshal(data)
		if err != nil {
			return err
		}

		payloadBuffer = *bytes.NewBuffer(data_json)
	}

	// Create new HTTP request object
	req, err := http.NewRequestWithContext(ctx, method, apiUrl.String(), &payloadBuffer)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("x-xdr-auth-id", string(*c.Config.ApiKeyId))
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

	// Execute request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
        //if c.Context.Err() != nil {
        if ctx.Err() != nil {
            return fmt.Errorf("context cancelled or timeout exceeded: %s", ctx.Err())
        }

		return err
	}
	defer res.Body.Close()

	// If API responds with HTTP 429 (Too Many Requests), sleep 3 seconds and try again
	if res.StatusCode == 429 {
		time.Sleep(3 * time.Second)
		return c.Request(ctx, method, endpoint, query, data, &response)
	}

	// If API responds with a non-OK status, return error
	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body from non-OK response: %s", err)
		}

		if err = json.Unmarshal(body, &errorResponse); err != nil {
			return err
		}

		return fmt.Errorf("API returned non-OK status code %d: %s", res.StatusCode, errorResponse.Err)
	}

	// Parse response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// If response body is non-empty, unmarshal into response object
	if len(body) > 0 && response != nil {
		if err = json.Unmarshal(body, response); err != nil {
			return err
		}
	}

	return nil
}
