package parallaxsdk

import "time"

type clientConfig struct {
	timeout time.Duration
	proxy   *string
}

type Option any
type CustomTimeoutOption time.Duration
type ClientProxyOption string

func WithCustomTimeout(d time.Duration) CustomTimeoutOption {
	return CustomTimeoutOption(d)
}

func WithClientProxy(proxy string) ClientProxyOption {
	return ClientProxyOption(proxy)
}

func parseOptions(options []Option) *clientConfig {
	config := &clientConfig{
		timeout: time.Second * 30,
		proxy:   nil,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case CustomTimeoutOption:
			config.timeout = time.Duration(v)
		case ClientProxyOption:
			proxyStr := string(v)
			config.proxy = &proxyStr
		}
	}

	return config
}
