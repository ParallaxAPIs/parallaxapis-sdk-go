package parallaxsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	"net/http"
)

type SDK struct {
	AuthKey string
	APIHost string
	client  *http.Client
}

type DatadomeSDK struct {
	*SDK
}

type PerimeterxSDK struct {
	*SDK
}

func NewPerimeterxSDK(key, host string, options ...Option) *PerimeterxSDK {
	return &PerimeterxSDK{
		SDK: CreateClient(
			key,
			host,
			options...,
		),
	}
}

func NewDatadomeSDK(key, host string, options ...Option) *DatadomeSDK {
	return &DatadomeSDK{
		SDK: CreateClient(
			key,
			host,
			options...,
		),
	}
}

func CreateClient(authKey, apiHost string, options ...Option) *SDK {
	clientConfig := parseOptions(options)

	if apiHost == "" {
		if strings.HasPrefix(authKey, "PX") {
			apiHost = DefaultPXHost
		} else if strings.HasPrefix(authKey, "DD") {
			apiHost = DefaultDDHost
		}
	}

	if !strings.HasPrefix(apiHost, "http://") && !strings.HasPrefix(apiHost, "https://") {
		apiHost = "https://" + apiHost
	}

	sdk := &SDK{
		AuthKey: authKey,
		APIHost: apiHost,
		client:  &http.Client{Timeout: clientConfig.timeout},
	}

	needTransport := clientConfig.proxy != nil || clientConfig.insecureSkipVerify
	if needTransport {
		tr := &http.Transport{}
		if clientConfig.proxy != nil {
			tr.Proxy = func(r *http.Request) (*url.URL, error) {
				return url.Parse(*clientConfig.proxy)
			}
		}
		if clientConfig.insecureSkipVerify {
			tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}
		sdk.client.Transport = tr
	}

	return sdk
}

func (s *SDK) request(endpoint string, payload any, out any) error {
	b, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	uri, err := url.JoinPath(s.APIHost, endpoint)

	if err != nil {
		return fmt.Errorf("creating uri error: %w", err)
	}

	resp, err := s.client.Post(uri, "application/json", bytes.NewReader(b))

	if err != nil {
		return fmt.Errorf("POST %s failed: %w", endpoint, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	var env ErrorEnv

	if err := json.Unmarshal(body, &env); err != nil {
		return fmt.Errorf("unmarshal envelope: %w, body: \n%s", err, string(body))
	}

	if env.Error {
		if len(env.Message) > 0 {
			return &APIError{Message: env.Message}
		}

		return &APIError{Message: env.Cookie}
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("unmarshal body to %T: %w", out, err)
	}

	return nil
}

// checkUsage
func (s *SDK) checkUsage(site string) (UsageResponse, error) {
	response := UsageResponse{}

	uri, err := url.JoinPath(s.APIHost, `/usage`)
	if err != nil {
		return response, fmt.Errorf("creating uri error: %w", err)
	}

	prms := url.Values{}
	prms.Add("site", site)
	prms.Add("authToken", s.AuthKey)

	uri = fmt.Sprintf(`%s?%s`, uri, prms.Encode())
	resp, err := s.client.Get(uri)
	if err != nil {
		return response, fmt.Errorf("GET %s failed: %w", uri, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("read response: %w", err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf("unmarshal body to %T: %w", response, err)
	}
	return response, nil
}
