package parallaxsdk

import "time"

type clientConfig struct {
	timeout            time.Duration
	proxy              *string
	insecureSkipVerify bool
}

type Option any
type CustomTimeoutOption time.Duration
type ClientProxyOption string
type InsecureSkipVerifyOption bool

func WithCustomTimeout(d time.Duration) CustomTimeoutOption {
	return CustomTimeoutOption(d)
}

func WithClientProxy(proxy string) ClientProxyOption {
	return ClientProxyOption(proxy)
}

func WithInsecureSkipVerify() InsecureSkipVerifyOption {
	return InsecureSkipVerifyOption(true)
}

func parseOptions(options []Option) *clientConfig {
	config := &clientConfig{
		timeout:            time.Second * 30,
		proxy:              nil,
		insecureSkipVerify: false,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case CustomTimeoutOption:
			config.timeout = time.Duration(v)
		case ClientProxyOption:
			proxyStr := string(v)
			config.proxy = &proxyStr
		case InsecureSkipVerifyOption:
			config.insecureSkipVerify = bool(v)
		}
	}

	return config
}
