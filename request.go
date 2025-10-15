package parallaxsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func NewPerimeterxSDK(key, host string) *PerimeterxSDK {
	return &PerimeterxSDK{
		SDK: CreateClient(
			key,
			host,
		),
	}
}

func NewDatadomeSDK(key, host string) *DatadomeSDK {
	return &DatadomeSDK{
		SDK: CreateClient(
			key,
			host,
		),
	}
}

func CreateClient(authKey, apiHost string) *SDK {
	if apiHost == "" {
		if strings.HasPrefix(authKey, "PX") {
			apiHost = DefaultPXHost
		} else if strings.HasPrefix(authKey, "DD") {
			apiHost = DefaultDDHost
		}
	}
	return &SDK{
		AuthKey: authKey,
		APIHost: apiHost,
		client:  &http.Client{},
	}
}

func (s *SDK) request(endpoint string, payload any, out any) error {
	b, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	uri := fmt.Sprintf("%s%s", s.APIHost, endpoint)

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
		return fmt.Errorf("unmarshal envelope: %w, body: %s", err, string(body))
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
